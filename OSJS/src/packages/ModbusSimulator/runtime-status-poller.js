'use strict';

const createStatusPoller = ({requestStatus, onStatus, onUnavailable, interval = 1000, setTimer = setTimeout, clearTimer = clearTimeout}) => {
  let selectedName = null;
  let timer = null;
  let generation = 0;
  let stopped = false;

  const clear = () => {
    if (timer !== null) clearTimer(timer);
    timer = null;
  };

  const request = token => {
    if (stopped || token !== generation || !selectedName) return;
    const name = selectedName;
    Promise.resolve(requestStatus(name))
      .then(status => {
        if (!stopped && token === generation && name === selectedName) onStatus(name, status);
      })
      .catch(error => {
        if (!stopped && token === generation && name === selectedName) onUnavailable(name, error);
      })
      .then(() => {
        if (!stopped && token === generation && name === selectedName) {
          timer = setTimer(() => request(token), interval);
        }
      });
  };

  const select = (name, force = false) => {
    const next = name || null;
    if (!force && next === selectedName) return;
    generation++;
    selectedName = next;
    clear();
    if (!stopped && selectedName) request(generation);
  };

  return {
    select,
    refresh: () => select(selectedName, true),
    stop: () => {
      stopped = true;
      generation++;
      clear();
    }
  };
};

const createRuntimeMessageController = ({publish}) => {
  let condition = null;
  let visible = false;

  const update = (next, message) => {
    if (next === condition) return;
    const recovering = condition !== null && next === null;
    condition = next;
    if (next !== null) {
      publish(message, true);
      visible = true;
    } else if (recovering && visible) {
      publish('Simulator runtime status recovered.', false);
      visible = false;
    }
  };

  return {
    operationMessage: () => {
      condition = null;
      visible = false;
    },
    status: status => {
      if (status && status.device_status === 'ERROR') {
        const detail = status.raw_ingest_error || 'No diagnostic detail was returned.';
        update(`simulator:${detail}`, `Simulator error: ${detail}`);
      } else {
        update(null, '');
      }
    },
    unavailable: error => {
      const detail = error && (error.message || String(error)) || 'Unknown runtime error.';
      update(`unavailable:${detail}`, `Simulator status unavailable: ${detail}`);
    }
  };
};

module.exports = {createStatusPoller, createRuntimeMessageController};
