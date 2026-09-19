import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';
import {createToolkit} from './toolkit-renderer';
import {createMemoryContract} from './memory-contract';
import {createMemoryTransport} from './memory-transport';
import {createMemoryEditor} from './memory-editor';
import {createReplicatorContract} from './replicator-contract';
import {createReplicatorTransport} from './replicator-transport';
import {createReplicatorEditor} from './replicator-editor';
import {createDiagnosticsEditor} from './diagnostics-editor';

// One OS.js Toolkit window. Diagnostics reads only the existing canonical
// Memory/Replicator contracts; no third socket, service-control API or fixture.
const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'MCSModbusToolkitWindow',
    title: 'MCS Modbus Toolkit',
    dimension: {width: 960, height: 640},
    position: 'center'
  });
  const memoryTransport = createMemoryTransport(proc);
  const memory = createMemoryContract(memoryTransport.send);
  const replicatorTransport = createReplicatorTransport(proc);
  const replicator = createReplicatorContract(replicatorTransport.send);
  let toolkit = null;
  let editor = null;
  let replicatorEditor = null;
  let diagnosticsEditor = null;

  win.on('destroy', () => {
    if (editor) { editor.destroy(); editor = null; }
    if (replicatorEditor) { replicatorEditor.destroy(); replicatorEditor = null; }
    if (diagnosticsEditor) { diagnosticsEditor.destroy(); diagnosticsEditor = null; }
    memoryTransport.close();
    replicatorTransport.close();
    if (toolkit) { toolkit.destroy(); toolkit = null; }
    proc.destroy();
  });
  win.render($content => {
    toolkit = createToolkit(document);
    const shadow = toolkit.element.shadowRoot;
    const memoryRoot = shadow.getElementById('simulator-root');
    const replicatorRoot = shadow.getElementById('replicator-root');
    const diagnosticsRoot = shadow.getElementById('panel-diagnostics');
    // No fixture fallback even when a canonical load or status fails.
    memoryRoot.replaceChildren();
    replicatorRoot.replaceChildren();
    diagnosticsRoot.replaceChildren();
    shadow.querySelector('.subtitle').textContent = 'Memory + Replicator / read-only Diagnostics';
    editor = createMemoryEditor(document, memoryRoot, memory);
    replicatorEditor = createReplicatorEditor(document, replicatorRoot, replicator);
    diagnosticsEditor = createDiagnosticsEditor(document, diagnosticsRoot, memory, replicator);
    $content.appendChild(toolkit.element);
  });

  return proc;
};

osjs.register(applicationName, register);
