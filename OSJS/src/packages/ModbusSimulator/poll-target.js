'use strict';

const appliedPollTargets = (persistedDevices = []) =>
  (Array.isArray(persistedDevices) ? persistedDevices : []).map(
    d => d && typeof d.name === 'string' ? d.name : null
  );

const pollTargetFor = (targets, selectedIndex) => {
  if (selectedIndex === null || selectedIndex < 0) return null;
  return (targets && targets[selectedIndex]) || null;
};

module.exports = {appliedPollTargets, pollTargetFor};