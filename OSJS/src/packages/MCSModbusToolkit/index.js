import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';

// UMIG-002: isolated OS.js-owned shell only. No Electron source or runtime
// services are imported here; the donor UI and adapters belong to later tasks.
const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'MCSModbusToolkitWindow',
    title: 'MCS Modbus Toolkit',
    dimension: {width: 960, height: 640},
    position: 'center'
  });

  win.on('destroy', () => proc.destroy());
  win.render($content => {
    const root = document.createElement('section');
    root.className = 'mcs-toolkit-placeholder';
    const heading = document.createElement('h1');
    heading.textContent = 'MCS Modbus Toolkit';
    const message = document.createElement('p');
    message.textContent = 'Toolkit shell staged. Memory, Replicator and Diagnostics are not connected yet.';
    root.append(heading, message);
    $content.appendChild(root);
  });

  return proc;
};

osjs.register(applicationName, register);
