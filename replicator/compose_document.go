package replicator

import (
	"fmt"
	"net"
	"strconv"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// ComposeDocumentDestinations replaces only Replicator-owned MMA2 reservations
// with the enabled destinations from doc. Foreign reservations are preserved.
func (s Store) ComposeDocumentDestinations(doc Document) (Document, mma2composer.EffectiveConfig, error) {
	resolved, err := s.ResolveDocumentDestinations(doc)
	if err != nil {
		return Document{}, mma2composer.EffectiveConfig{}, err
	}
	composer := mma2composer.New(s.Root, ProducerReplicator)
	cfg, err := composer.LoadEffective()
	if err != nil {
		return Document{}, mma2composer.EffectiveConfig{}, err
	}
	owners, err := composer.LoadOwners()
	if err != nil {
		return Document{}, mma2composer.EffectiveConfig{}, err
	}

	for _, device := range resolved.Devices {
		if !device.Enabled {
			continue
		}
		if err := mma2composer.Collision(device.Destination.Port, device.Destination.UnitID, ProducerReplicator, owners); err != nil {
			return Document{}, mma2composer.EffectiveConfig{}, fmt.Errorf("device %q: %w", device.Name, err)
		}
	}

	cfg, owners = composer.DropProducerReservations(cfg, owners)
	for i := range resolved.Devices {
		device := &resolved.Devices[i]
		if !device.Enabled {
			device.Destination.Owner = ProducerReplicator
			device.Destination.Status = "AVAILABLE"
			continue
		}
		memory, err := destinationMemoryForBlocks(device.Destination.UnitID, device.blocks())
		if err != nil {
			return Document{}, mma2composer.EffectiveConfig{}, fmt.Errorf("device %q: %w", device.Name, err)
		}
		id := fmt.Sprintf("replicator-%d-%d", device.Destination.Port, device.Destination.UnitID)
		listen := net.JoinHostPort("0.0.0.0", strconv.Itoa(int(device.Destination.Port)))
		cfg = mma2composer.AddMemory(cfg, id, listen, memory)
		owners.Reservations = append(owners.Reservations, mma2composer.OwnershipEntry{
			Port: device.Destination.Port, UnitID: device.Destination.UnitID, Owner: ProducerReplicator,
		})
		device.Destination.Owner = ProducerReplicator
		device.Destination.Status = "OWNED"
	}
	if err := composer.Commit(cfg, owners); err != nil {
		return Document{}, mma2composer.EffectiveConfig{}, err
	}
	return resolved, cfg, nil
}

func unionArea(current *mma2composer.Area, start, count uint16) (*mma2composer.Area, error) {
	if current == nil {
		return &mma2composer.Area{Start: start, Count: count}, nil
	}
	lo := uint32(current.Start)
	hi := lo + uint32(current.Count)
	blockLo := uint32(start)
	blockHi := blockLo + uint32(count)
	if blockLo < lo {
		lo = blockLo
	}
	if blockHi > hi {
		hi = blockHi
	}
	span := hi - lo
	if span == 0 || span > 0xffff {
		return nil, fmt.Errorf("combined destination area span %d cannot be represented by MMA2", span)
	}
	return &mma2composer.Area{Start: uint16(lo), Count: uint16(span)}, nil
}

func destinationMemoryForBlocks(unitID uint16, blocks []PullBlock) (mma2composer.Memory, error) {
	if len(blocks) == 0 {
		return mma2composer.Memory{}, fmt.Errorf("at least one pull block is required")
	}
	memory := mma2composer.Memory{UnitID: unitID}
	for _, block := range blocks {
		var err error
		switch block.Function {
		case 1:
			memory.Coils, err = unionArea(memory.Coils, block.Start, block.Count)
		case 2:
			memory.DiscreteInputs, err = unionArea(memory.DiscreteInputs, block.Start, block.Count)
		case 3:
			memory.HoldingRegs, err = unionArea(memory.HoldingRegs, block.Start, block.Count)
		case 4:
			memory.InputRegs, err = unionArea(memory.InputRegs, block.Start, block.Count)
		default:
			return mma2composer.Memory{}, fmt.Errorf("unsupported pull block function %d", block.Function)
		}
		if err != nil {
			return mma2composer.Memory{}, err
		}
	}
	memory.Policy = &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{
		ID:       "replicator-fc-access",
		SourceIP: []string{"0.0.0.0/0", "::/0", "127.0.0.1", "::1"},
		AllowFC:  []uint8{1, 2, 3, 4, 5, 6, 15, 16},
	}}}
	return memory, nil
}

func documentDestinationPorts(doc Document) []uint16 {
	seen := make(map[uint16]bool)
	ports := make([]uint16, 0, len(doc.Devices))
	for _, device := range doc.Devices {
		if !device.Enabled || device.Destination.Port == 0 || seen[device.Destination.Port] {
			continue
		}
		seen[device.Destination.Port] = true
		ports = append(ports, device.Destination.Port)
	}
	return ports
}
