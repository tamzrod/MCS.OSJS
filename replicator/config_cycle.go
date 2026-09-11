package replicator

import (
	"fmt"
	"strings"
)

// runConfigCycle executes one already-composed device mapping. It intentionally
// does not touch shared MMA2 configuration; document apply owns that structural
// lifecycle so poll cycles cannot rebuild other Replicator reservations.
func runConfigCycle(cfg Config) (CycleResult, error) {
	if err := validateCycleMapping(cfg); err != nil {
		return CycleResult{}, err
	}
	payload, err := ReadSourceRange(cfg.Source)
	if err != nil {
		return CycleResult{}, err
	}
	if sourcePayloadCount(payload) != int(cfg.Destination.Count) {
		return CycleResult{}, fmt.Errorf("source returned %d values; destination expects %d", sourcePayloadCount(payload), cfg.Destination.Count)
	}
	if err := sendDestination(cfg.Destination, payload); err != nil {
		return CycleResult{}, err
	}
	return CycleResult{
		Function:         payload.Function,
		SourceStart:      payload.Start,
		DestinationArea:  strings.ToLower(strings.TrimSpace(cfg.Destination.Area)),
		DestinationStart: cfg.Destination.Start,
		Count:            cfg.Destination.Count,
	}, nil
}
