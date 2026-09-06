// OS.js CLI configuration for the Nameless isolated prototype (OSJS-002).
// Adds local package source (src/packages) to discovery.
const path = require('path');

module.exports = {
  discover: [
    path.resolve(__dirname, '../packages')
  ],
  tasks: []
};
