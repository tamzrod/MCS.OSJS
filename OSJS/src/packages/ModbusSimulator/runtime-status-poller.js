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

module.exports = {createStatusPoller};
