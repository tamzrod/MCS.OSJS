package replicator

import (
	"fmt"
	"net"
	"strings"
)

// runConfigCycle executes one already-composed device mapping. It intentionally
// does not touch shared MMA2 configuration; document apply owns that structural
// lifecycle so poll cycles cannot rebuild other Replicator reservations.
func runConfigCycle(cfg Config) (CycleResult, error) {
	return runConfigCycleObserved(cfg, &CycleComms{})
}

func runConfigCycleObserved(cfg Config, comms *CycleComms) (CycleResult, error) {
	if err := validateCycleMapping(cfg); err != nil {
		return CycleResult{}, err
	}
	payload, err := readSourceObserved(cfg.Source, defaultModbusTimeout, comms)
	if err != nil {
		return CycleResult{}, err
	}
	if sourcePayloadCount(payload) != int(cfg.Destination.Count) {
		return CycleResult{}, fmt.Errorf("source returned %d values; destination expects %d", sourcePayloadCount(payload), cfg.Destination.Count)
	}
	if err := sendDestination(cfg.Destination, payload); err != nil {
		comms.MMA2 = observation("ERROR", "WRITE_FAILED", net.JoinHostPort("127.0.0.1", fmt.Sprint(cfg.Destination.ListenerPort)))
		comms.MMA2.Error = err.Error()
		return CycleResult{}, err
	}
	comms.MMA2 = observation("OK", "ACKNOWLEDGED", net.JoinHostPort("127.0.0.1", fmt.Sprint(cfg.Destination.ListenerPort)))
	comms.MMA2.LastSuccessAt = comms.MMA2.ObservedAt
	return CycleResult{
		Function:         payload.Function,
		SourceStart:      payload.Start,
		DestinationArea:  strings.ToLower(strings.TrimSpace(cfg.Destination.Area)),
		DestinationStart: cfg.Destination.Start,
		Count:            cfg.Destination.Count,
	}, nil
}
