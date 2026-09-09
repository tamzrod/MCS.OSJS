'use strict';

const UNAVAILABLE = 'UNAVAILABLE';
const NEUTRAL = '\u2014';
const KNOWN = ['RUNNING', 'WAITING', 'STOPPED', 'ERROR'];

const statusWord = ({ device, appliedName, runtimeStatus, runtimeStatusError, field }) => {
  if (!device) return NEUTRAL;
  if (!appliedName || appliedName !== device.name) return NEUTRAL;
  if (!runtimeStatus) return runtimeStatusError ? UNAVAILABLE : NEUTRAL;
  const value = runtimeStatus[field];
  return KNOWN.includes(value) ? value : UNAVAILABLE;
};

module.exports = { statusWord };
