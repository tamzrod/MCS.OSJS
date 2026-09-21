#!/usr/bin/env python3
"""One-shot, fail-closed, synthetic Docker cross-process VERIFY for EM-003-V.
Run ONLY when named by the current Operation CWAL handoff. No production target.
"""
import concurrent.futures
import hashlib
import json
import os
import re
import socket
import subprocess
import sys
import time
from pathlib import Path

PROJECT = "mcs-em003v-20260920-one"
COMPOSE = "deploy/verify/compose.yaml"
SOURCE = "538324a472eb15ff8ef97ee66f826fa02ba9462f"
DOCKER = ["sudo", "-n", "docker"]
C = DOCKER + ["compose", "-f", COMPOSE, "-p", PROJECT]
FILES = ["config/mma2/config.yaml", "config/mma2/owners.yaml",
         "config/simulator/devices.yaml", "config/replicator/devices.yaml",
         "config/mma2/restart-request.yaml", "config/mma2/restart-ack"]
CLIENT = r'''
const net=require('net');const path=process.argv[1],message=Buffer.from(process.argv[2]);
let received=Buffer.alloc(0),done=false;const sock=net.createConnection(path,()=>{
 const head=Buffer.alloc(4);head.writeUInt32BE(message.length);sock.write(Buffer.concat([head,message]));
});
sock.setTimeout(45000,()=>{console.error('IPC_TIMEOUT');sock.destroy();process.exitCode=2;});
sock.on('data',b=>{received=Buffer.concat([received,b]);if(received.length>=4){
 const size=received.readUInt32BE(0);if(size<1||size>1048576){console.error('BAD_FRAME');sock.destroy();process.exitCode=2;return;}
 if(received.length>=4+size&&!done){done=true;console.log(received.subarray(4,4+size).toString());sock.end();}
}});
sock.on('error',e=>{console.error('IPC_ERROR '+e.message);process.exitCode=2;});
sock.on('end',()=>{if(!done){console.error('IPC_TRUNCATED');process.exitCode=2;}});
'''
SNAP = r'''
const fs=require('fs'),crypto=require('crypto');const paths=JSON.parse(process.argv[1]),out={};
for(const p of paths){const f='/data/'+p;try{const b=fs.readFileSync(f);out[p]={sha256:crypto.createHash('sha256').update(b).digest('hex'),text:b.toString('utf8')};}
catch(e){if(e.code!=='ENOENT')throw e;out[p]=null;}}console.log(JSON.stringify(out));
'''
MONITOR = r'''
const fs=require('fs'),crypto=require('crypto');
const paths=['/data/config/mma2/restart-request.yaml','/data/config/mma2/restart-ack'];
const output='/tmp/em003v-restarts.ndjson',ready='/tmp/em003v-monitor.ready',stop='/tmp/em003v-monitor.stop';
if(fs.existsSync(output)||fs.existsSync(ready)||fs.existsSync(stop))process.exit(7);
const seen={};fs.writeFileSync(ready,'ready');
const poll=()=>{for(const p of paths){let b=null;try{b=fs.readFileSync(p);}catch(e){if(e.code!=='ENOENT')throw e;}
const hash=b?crypto.createHash('sha256').update(b).digest('hex'):null;
if(seen[p]!==hash){seen[p]=hash;if(b)fs.appendFileSync(output,JSON.stringify({time:new Date().toISOString(),kind:p.endsWith('restart-ack')?'ack':'request',body:b.toString('utf8')})+'\n');}}
if(fs.existsSync(stop)){clearInterval(timer);process.exit(0);}};
const timer=setInterval(poll,5);poll();setTimeout(()=>process.exit(0),120000);
'''

class Gate(Exception):
    pass

class ProductFailure(Exception):
    pass


def stamp(label, **values):
    print(json.dumps({"at": time.strftime('%Y-%m-%dT%H:%M:%SZ', time.gmtime()), "stage": label, **values}, sort_keys=True), flush=True)


def cmd(argv, timeout=90, ok=(0,)):
    stamp("command", argv=argv)
    try:
        p=subprocess.run(argv, text=True, capture_output=True, timeout=timeout)
    except subprocess.TimeoutExpired as e:
        stamp("timeout", argv=argv, seconds=timeout)
        raise Gate("command timeout") from e
    stamp("exit", code=p.returncode, stdout=p.stdout[-16000:], stderr=p.stderr[-16000:])
    if p.returncode not in ok:
        raise Gate("command exit %s: %s" % (p.returncode, argv))
    return p.stdout.strip()


def d(*parts, timeout=90, ok=(0,)):
    return cmd(DOCKER+list(parts),timeout,ok)


def dc(*parts, timeout=90, ok=(0,)):
    return cmd(C+list(parts),timeout,ok)


def logs(container):
    argv=DOCKER+['logs','--timestamps',container]
    stamp('command',argv=argv)
    p=subprocess.run(argv,text=True,capture_output=True,timeout=45)
    merged=p.stdout+p.stderr
    stamp('exit',code=p.returncode,stdout=merged[-16000:])
    check(p.returncode==0,'Docker logs unavailable')
    return merged

def check(cond, message, product=False):
    if not cond:
        raise (ProductFailure if product else Gate)(message)


def ipc(container, name, req, timeout=60):
    start=time.monotonic(); stamp('request_start',runtime=name,request=req)
    argv=DOCKER+["exec",container,"node","-e",CLIENT,"/data/run/modbus-"+name+".sock",json.dumps(req,separators=(',',':'))]
    try:
        p=subprocess.run(argv,text=True,capture_output=True,timeout=timeout)
    except subprocess.TimeoutExpired as e:
        raise ProductFailure(name+' IPC timed out') from e
    end=time.monotonic()
    stamp('request_end',runtime=name,exit=p.returncode,start_monotonic=start,end_monotonic=end,stdout=p.stdout,stderr=p.stderr)
    check(p.returncode==0,name+' IPC process exit '+str(p.returncode),True)
    try:
        response=json.loads(p.stdout)
    except ValueError as e:
        raise ProductFailure(name+' invalid JSON response') from e
    check(response.get('request_id')==req['request_id'] and response.get('ok') is True,
          name+' apply was not acknowledged OK: '+p.stdout,True)
    return (start,end,response)


def snap(container, label):
    value=d('exec',container,'node','-e',SNAP,json.dumps(FILES))
    out=json.loads(value)
    stamp('snapshot',label=label,files={p: None if o is None else o['sha256'] for p,o in out.items()})
    return out


def owners(data):
    doc=data['config/mma2/owners.yaml']
    check(doc is not None,'ownership file absent',True)
    entries=re.findall(r'port:\s*(\d+)\s*\n\s*unit_id:\s*(\d+)\s*\n\s*owner:\s*(\w+)',doc['text'])
    return {(int(port),int(unit),owner) for port,unit,owner in entries}


def monitor(container):
    raw=d('exec',container,'node','-e',"const fs=require('fs');process.stdout.write(fs.existsSync('/tmp/em003v-restarts.ndjson')?fs.readFileSync('/tmp/em003v-restarts.ndjson','utf8'):'');")
    return [json.loads(line) for line in raw.splitlines() if line.strip()]


def safe_down(created):
    if not created:
        return
    # Fail closed: never remove anything whose project label is absent or unexpected.
    for typ, name in [('volume',PROJECT+'_verify-data'),('network',PROJECT+'_verify-runtime'),('network',PROJECT+'_verify-ui')]:
        names=d(typ,'ls','--format','{{.Name}}').splitlines()
        if name not in names:
            stamp('cleanup_resource_absent',kind=typ,name=name)
            continue
        label=d(typ,'inspect',name,'--format','{{index .Labels "com.docker.compose.project"}}')
        check(label==PROJECT,'refuse cleanup of unowned '+typ+' '+name)
    ids=dc('ps','-aq').split()
    for ident in ids:
        label=d('inspect',ident,'--format','{{index .Config.Labels "com.docker.compose.project"}}')
        check(label==PROJECT,'refuse cleanup of unowned container '+ident)
    stamp('cleanup_start',project=PROJECT,container_ids=ids)
    dc('down','--remove-orphans',timeout=120)
    volume=d('volume','inspect',PROJECT+'_verify-data','--format','{{index .Labels "com.docker.compose.project"}}')
    check(volume==PROJECT,'retained volume not correctly labelled')
    check(not dc('ps','-aq'),'project containers remain after down')
    stamp('cleanup_complete',retained_volume=PROJECT+'_verify-data')
    return PROJECT+'_verify-data'


def finalize(created, container, status, reason):
    retained_volume=None
    try:
        if container:
            stamp('final_snapshot',snapshot={p:None if o is None else o['sha256'] for p,o in snap(container,'final_before_down').items()})
    except Exception as error:
        stamp('final_snapshot_failure',error=str(error))
        reason+='; final snapshot failure: '+str(error)
        if status=='PASS':status='INCOMPLETE'
    try:
        retained_volume=safe_down(created)
    except Exception as error:
        stamp('cleanup_failure',error=str(error))
        reason+='; cleanup failure: '+str(error)
        if status=='PASS':status='INCOMPLETE'
    stamp('FINAL_VERDICT',verdict=status,reason=reason,project=PROJECT,retained_volume=retained_volume)
    return status, reason


def main():
    created=False; container=None; baseline=None; status='BLOCKED'; reason='not started'
    try:
        check(Path(COMPOSE).is_file(),'verify compose missing')
        check(not cmd(['git','status','--porcelain','--untracked-files=all']),'checkout not clean')
        head=cmd(['git','rev-parse','HEAD']);track=cmd(['git','rev-parse','origin/main'])
        remote=cmd(['git','ls-remote','--exit-code','origin','refs/heads/main']).split()[0]
        check(head==track==remote,'checkout stale or remote advanced')
        cmd(['git','merge-base','--is-ancestor',SOURCE,'HEAD'])
        check(not cmd(['git','diff','--name-only',SOURCE,'HEAD','--','simulator','replicator','mma2composer','MMA2','OSJS','deploy']),
              'product/compose changed since checkpoint')
        check(not any(os.getenv(k) for k in ['DOCKER_HOST','DOCKER_CONTEXT','COMPOSE_PROJECT_NAME','MCS_VERIFY_OSJS_PORT']),
              'Docker or Compose override present')
        check(d('context','show')=='default','Docker context not default')
        hostname=socket.gethostname()
        info=d('info','--format','{{.Name}} {{.DockerRootDir}}')
        check(info==hostname+' /var/lib/docker','daemon not verified sandbox-local: '+info)
        check(not d('ps','-aq'),'Docker daemon already has containers')
        check(not d('volume','ls','-q'),'Docker daemon already has volumes')
        networks=d('network','ls','--format','{{.Name}}').splitlines()
        check(sorted(networks)==['bridge','host','none'],'unexpected Docker networks: '+str(networks))
        config=dc('config')
        check(('127.0.0.1' in config and '18219' in config and 'internal: true' in config and
               'network_mode: service:mma2' in config and not re.search(r'(?m)^\s*privileged:\s*true',config)),
              'verify Compose topology contradicts reviewed template')
        check(not dc('ps','-aq'),'project already exists')
        check(d('volume','inspect',PROJECT+'_verify-data',ok=(1,))=='','project volume already exists')
        # One and only one application startup. No subsequent Compose up/start.
        created=True
        dc('up','-d','--build',timeout=900)
        ids={s:dc('ps','-q',s) for s in ['mma2','modbus-simulator-runtime','modbus-replicator-runtime','osjs-shell']}
        check(all(ids.values()),'missing test container')
        for service,ident in ids.items():
            label=d('inspect',ident,'--format','{{index .Config.Labels "com.docker.compose.project"}}')
            check(label==PROJECT,service+' has wrong project label')
            running=d('inspect',ident,'--format','{{.State.Running}}')
            check(running=='true',service+' not running')
        container=ids['osjs-shell']; mma=ids['mma2']
        for name in ('simulator','replicator'):
            socketpath='/data/run/modbus-'+name+'.sock'
            for attempt in range(60):
                present=d('exec',container,'node','-e',"process.stdout.write(require('fs').existsSync(process.argv[1])?'yes':'no')",socketpath)
                if present=='yes':break
                time.sleep(1)
            else:raise Gate(name+' socket not ready in 60s')
        baseline=snap(container,'baseline')
        for p,key in [('config/mma2/config.yaml','listeners'),('config/simulator/devices.yaml','devices'),('config/replicator/devices.yaml','devices')]:
            check(baseline[p] is not None and re.fullmatch(r'\s*'+key+r':\s*\[\]\s*',baseline[p]['text']) is not None,
                  'nonempty baseline '+p)
        check(not owners(baseline),'baseline has reservations')
        check(baseline[FILES[4]] is None and baseline[FILES[5]] is None,'stale restart artifacts')
        # Read-only test-container observation, bounded monitor writes only its own /tmp.
        d('exec','-d',container,'node','-e',MONITOR)
        for attempt in range(30):
            if d('exec',container,'node','-e',"process.stdout.write(require('fs').existsSync('/tmp/em003v-monitor.ready')?'yes':'no')")=='yes':break
            time.sleep(0.1)
        else:raise Gate('restart observer not ready')
        sim={'version':1,'request_id':'em003v-sim-create','operation':'apply','payload':{'document':{'devices':[
             {'name':'EM003V-SIM','enabled':True,'mma2':{'port':15020,'unit_id':1,'fc1':{'start':0,'count':0},
              'fc2':{'start':0,'count':0},'fc3':{'start':10,'count':4},'fc4':{'start':0,'count':0}},
              'random_runtime':{'fc1_interval_ms':0,'fc2_interval_ms':0,'fc3_interval_ms':0,'fc4_interval_ms':0}}]}}}
        rep={'version':1,'request_id':'em003v-rep-create','operation':'apply','payload':{'document':{'devices':[
             {'name':'EM003V-REP','enabled':True,'endpoint':'127.0.0.1:15020','unit_id':1,
              'pull_blocks':[{'function':3,'start':10,'count':4,'scan_rate_ms':500}],
              'destination':{'port':15021,'unit_id':1,'auto_port':False,'auto_unit_id':False}}]}}}
        with concurrent.futures.ThreadPoolExecutor(max_workers=2) as pool:
            a=pool.submit(ipc,container,'simulator',sim)
            b=pool.submit(ipc,container,'replicator',rep)
            sa,ea,ra=a.result();sb,eb,rb=b.result()
        check(max(sa,sb)<=min(ea,eb),'request intervals did not overlap; concurrency unproven')
        check(abs(sa-sb)<1.0,'start skew exceeds one second')
        after=snap(container,'post_concurrent_apply')
        check(owners(after)=={(15020,1,'simulator'),(15021,1,'replicator')},'lost/foreign owner mutation',True)
        configtext=after['config/mma2/config.yaml']['text']
        check(configtext.count('listen: 0.0.0.0:15020')==1 and configtext.count('listen: 0.0.0.0:15021')==1,
              'effective MMA2 configuration lost/duplicated a listener',True)
        check(configtext.count('holding_registers:')==2,'effective config lost FC3 area',True)
        for name,expected in [('simulator','EM003V-SIM'),('replicator','EM003V-REP')]:
            load={'version':1,'request_id':'em003v-'+name+'-load','operation':'load','payload':{}}
            _,_,answer=ipc(container,name,load)
            devices=answer['result']['document']['devices']
            check(len(devices)==1 and devices[0]['name']==expected,name+' document not preserved',True)
        events=monitor(container)
        requests=[e for e in events if e['kind']=='request']
        acks=[e for e in events if e['kind']=='ack']
        stamp('restart_observations',events=events)
        check(len(requests)==2 and len(acks)==2,'two restart request/ack pairs not observable')
        for req,ack in zip(requests,acks):
            found=re.search(r'(?m)^config_sha256:\s*([0-9a-f]{64})\s*$',req['body'])
            check(found is not None and ack['body']==found.group(1) and req['time']<=ack['time'],
                  'restart ACK SHA mismatch or sequence incorrect',True)
        check(requests[0]['time']<=acks[0]['time']<=requests[1]['time']<=acks[1]['time'],
              'restart transactions interleaved',True)
        logtext=logs(mma)
        check(logtext.count('restart request consumed fingerprint=')==2,'MMA2 supervisor restart count not two',True)
        check(after[FILES[4]] is None and after[FILES[5]] is None,'restart artifacts not cleared',True)
        stamp('concurrent_verify',result='PASS',simulator=ra,replicator=rb,owner_set=sorted(owners(after)))
        # Restore only test-only state by each producer's own authorized apply API.
        for name in ('replicator','simulator'):
            empty={'version':1,'request_id':'em003v-'+name+'-restore','operation':'apply','payload':{'document':{'devices':[]}}}
            ipc(container,name,empty)
        restored=snap(container,'post_restore')
        check(all(restored[p]==baseline[p] for p in FILES),
              'baseline bytes/hashes not restored by producer APIs',True)
        logtext=logs(mma)
        check(logtext.count('restart request consumed fingerprint=')==4,'cleanup restart count not four',True)
        d('exec',container,'node','-e',"require('fs').writeFileSync('/tmp/em003v-monitor.stop','stop')")
        stamp('verify_result',verdict='PASS',scope='one concurrent synthetic apply per producer + API restoration')
        status='PASS';reason='all runtime, restart and restoration checks passed'
    except ProductFailure as e:
        status='FAIL';reason=str(e);stamp('failure',kind='product',reason=reason)
    except (Gate,Exception) as e:
        status='BLOCKED';reason=str(e);stamp('failure',kind='environment_or_evidence',reason=reason)
    finally:
        status, reason=finalize(created, container, status, reason)
    return 0 if status=='PASS' else 1

if __name__=='__main__':
    sys.exit(main())
