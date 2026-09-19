import './index.scss';
import osjs from 'osjs';
import {name as applicationName} from './metadata.json';
import {createToolkit} from './toolkit-renderer';
import {createMemoryContract} from './memory-contract';
import {createMemoryTransport} from './memory-transport';
import {createMemoryEditor} from './memory-editor';

// One OS.js-owned window. Only Memory uses a live adapter; Replicator and
// Diagnostics remain fixture-only. No legacy Simulator application imports.
const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'MCSModbusToolkitWindow',
    title: 'MCS Modbus Toolkit',
    dimension: {width: 960, height: 640},
    position: 'center'
  });
  const transport = createMemoryTransport(proc);
  const memory = createMemoryContract(transport.send);
  let toolkit = null;
  let editor = null;

  win.on('destroy', () => {
    if (editor) { editor.destroy(); editor = null; }
    transport.close();
    if (toolkit) { toolkit.destroy(); toolkit = null; }
    proc.destroy();
  });
  win.render($content => {
    toolkit = createToolkit(document);
    const shadow = toolkit.element.shadowRoot;
    const memoryRoot = shadow.getElementById('simulator-root');
    // Remove fixture DOM before the Toolkit is mounted. A failed load must
    // never display sample values as if they were canonical configuration.
    memoryRoot.replaceChildren();
    shadow.querySelector('.subtitle').textContent = 'Memory runtime / Replicator + Diagnostics preview';
    editor = createMemoryEditor(document, memoryRoot, memory);
    $content.appendChild(toolkit.element);
  });

  return proc;
};

osjs.register(applicationName, register);
