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
	values, err := ReadSourceRange(cfg.Source)
	if err != nil {
		return CycleResult{}, err
	}
	if len(values.Values) != int(cfg.Destination.Count) {
		return CycleResult{}, fmt.Errorf("source returned %d registers; destination expects %d", len(values.Values), cfg.Destination.Count)
	}
	if err := sendDestination(cfg.Destination, values.Values); err != nil {
		return CycleResult{}, err
	}
	return CycleResult{
		Function:         values.Function,
		SourceStart:      values.Start,
		DestinationArea:  strings.ToLower(strings.TrimSpace(cfg.Destination.Area)),
		DestinationStart: cfg.Destination.Start,
		Count:            cfg.Destination.Count,
	}, nil
}
