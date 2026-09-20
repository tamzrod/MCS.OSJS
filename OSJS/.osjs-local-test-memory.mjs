import { createRequire } from 'module';
const require = createRequire(import.meta.url);

const OSjsService = require('./dist/osjs.js').default;

async function run() {
  const service = new OSjsService({ root: '/home/sysadmin/apps/MCS.OSJS-jr' });
  
  console.log('📋 Testing basic Memory device...\n');

  // Step 1: Create synthetic FC3 memory  
  console.log('Step 1: Creating synthetic FC3 memory device...');
  try {
    const mem = await service.resourceFactory.getServiceInstance('@osjs/resource-manager/memory', 'fc3');
    console.log(`✓ Memory service ID: ${mem.id}`);
    return [mem, service];
  } catch(err) {
    console.error('✗ Failed:', err.message);
    process.exit(1);
  }
}

run();
