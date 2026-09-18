import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';
import {createToolkit} from './toolkit-renderer';

// UMIG-003: one OS.js-owned window with a self-contained fixture renderer.
// Real Simulator/Replicator adapters belong to the later UMIG-004/005 stages.
const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'MCSModbusToolkitWindow',
    title: 'MCS Modbus Toolkit',
    dimension: {width: 960, height: 640},
    position: 'center'
  });
  let toolkit = null;

  win.on('destroy', () => {
    if (toolkit) {
      toolkit.destroy();
      toolkit = null;
    }
    proc.destroy();
  });
  win.render($content => {
    toolkit = createToolkit(document);
    $content.appendChild(toolkit.element);
  });

  return proc;
};

osjs.register(applicationName, register);
