package replicator

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"github.com/tamzrod/MCS.OSJS/mma2raw"
)

const ProducerReplicator = "replicator"

type CycleResult struct {
	Function         uint8
	SourceStart      uint16
	DestinationArea  string
	DestinationStart uint16
	Count            uint16
}

func (s Store) RunOnce() (CycleResult, error) {
	cfg, err := s.Load()
	if err != nil {
		return CycleResult{}, err
	}
	if err := validateCycleMapping(cfg); err != nil {
		return CycleResult{}, err
	}
	if err := s.composeDestination(cfg.Destination); err != nil {
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
	return CycleResult{Function: payload.Function, SourceStart: payload.Start, DestinationArea: strings.ToLower(strings.TrimSpace(cfg.Destination.Area)), DestinationStart: cfg.Destination.Start, Count: cfg.Destination.Count}, nil
}

func validateCycleMapping(cfg Config) error {
	if cfg.Source.Count != cfg.Destination.Count {
		return fmt.Errorf("source.count and destination.count must match for 1:1 replication")
	}
	if cfg.Source.Function < 1 || cfg.Source.Function > 4 {
		return fmt.Errorf("replication supports FC1, FC2, FC3, and FC4 only")
	}
	area := strings.ToLower(strings.TrimSpace(cfg.Destination.Area))
	want := fmt.Sprintf("fc%d", cfg.Source.Function)
	if area != want {
		return fmt.Errorf("cross-area mapping is not supported: source fc%d to destination %s", cfg.Source.Function, area)
	}
	return nil
}

func sourcePayloadCount(payload RegisterValues) int {
	if payload.Function == 1 || payload.Function == 2 {
		return len(payload.Bits)
	}
	return len(payload.Values)
}

func (s Store) composeDestination(dst DestinationConfig) error {
	composer := mma2composer.New(s.Root, ProducerReplicator)
	cfg, err := composer.LoadEffective()
	if err != nil { return err }
	owners, err := composer.LoadOwners()
	if err != nil { return err }
	if err := mma2composer.Collision(dst.ListenerPort, dst.UnitID, ProducerReplicator, owners); err != nil { return err }
	cfg, owners = composer.DropProducerReservations(cfg, owners)
	mem, err := destinationMemory(dst)
	if err != nil { return err }
	id := fmt.Sprintf("replicator-%d-%d", dst.ListenerPort, dst.UnitID)
	listen := net.JoinHostPort("0.0.0.0", strconv.Itoa(int(dst.ListenerPort)))
	cfg = mma2composer.AddMemory(cfg, id, listen, mem)
	owners.Reservations = append(owners.Reservations, mma2composer.OwnershipEntry{Port: dst.ListenerPort, UnitID: dst.UnitID, Owner: ProducerReplicator})
	return composer.Commit(cfg, owners)
}

func destinationMemory(dst DestinationConfig) (mma2composer.Memory, error) {
	mem := mma2composer.Memory{UnitID: dst.UnitID}
	area := &mma2composer.Area{Start: dst.Start, Count: dst.Count}
	switch strings.ToLower(strings.TrimSpace(dst.Area)) {
	case "fc1": mem.Coils = area
	case "fc2": mem.DiscreteInputs = area
	case "fc3": mem.HoldingRegs = area
	case "fc4": mem.InputRegs = area
	default: return mma2composer.Memory{}, fmt.Errorf("destination area must be fc1, fc2, fc3, or fc4")
	}
	mem.Policy = &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{ID: "replicator-fc-access", SourceIP: []string{"0.0.0.0/0", "::/0", "127.0.0.1", "::1"}, AllowFC: []uint8{1, 2, 3, 4, 5, 6, 15, 16}}}}
	return mem, nil
}

func sendDestination(dst DestinationConfig, payload RegisterValues) error {
	area, err := rawArea(dst.Area)
	if err != nil { return err }
	addr := net.JoinHostPort("127.0.0.1", strconv.Itoa(int(dst.ListenerPort)))
	client := mma2raw.NewClient(addr, dst.UnitID, map[mma2raw.Area]mma2raw.Range{area: {Start: dst.Start, Count: dst.Count}})
	values := mma2raw.Values{}
	if area == mma2raw.Coils || area == mma2raw.DiscreteInputs { values.Bits = payload.Bits } else { values.Registers = payload.Values }
	if err := client.Send(area, values); err != nil { return fmt.Errorf("destination raw ingest: %w", err) }
	return nil
}

func rawArea(area string) (mma2raw.Area, error) {
	switch strings.ToLower(strings.TrimSpace(area)) {
	case "fc1": return mma2raw.Coils, nil
	case "fc2": return mma2raw.DiscreteInputs, nil
	case "fc3": return mma2raw.HoldingRegisters, nil
	case "fc4": return mma2raw.InputRegisters, nil
	default: return 0, fmt.Errorf("destination area must be fc1, fc2, fc3, or fc4")
	}
}
