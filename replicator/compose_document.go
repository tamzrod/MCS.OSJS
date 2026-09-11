package replicator

import (
	"fmt"
	"net"
	"strconv"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// ComposeDocumentDestinations replaces only Replicator-owned MMA2 reservations
// with the enabled destinations from doc. Foreign reservations are preserved
// verbatim and first-come ownership is checked before any shared artifact write.
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
		runtimeCfg, err := device.runtimeConfig()
		if err != nil {
			return Document{}, mma2composer.EffectiveConfig{}, fmt.Errorf("device %q: %w", device.Name, err)
		}
		memory, err := destinationMemory(runtimeCfg.Destination)
		if err != nil {
			return Document{}, mma2composer.EffectiveConfig{}, fmt.Errorf("device %q: %w", device.Name, err)
		}
		id := fmt.Sprintf("replicator-%d-%d", runtimeCfg.Destination.ListenerPort, runtimeCfg.Destination.UnitID)
		listen := net.JoinHostPort("0.0.0.0", strconv.Itoa(int(runtimeCfg.Destination.ListenerPort)))
		cfg = mma2composer.AddMemory(cfg, id, listen, memory)
		owners.Reservations = append(owners.Reservations, mma2composer.OwnershipEntry{
			Port: runtimeCfg.Destination.ListenerPort, UnitID: runtimeCfg.Destination.UnitID, Owner: ProducerReplicator,
		})
		device.Destination.Owner = ProducerReplicator
		device.Destination.Status = "OWNED"
	}
	if err := composer.Commit(cfg, owners); err != nil {
		return Document{}, mma2composer.EffectiveConfig{}, err
	}
	return resolved, cfg, nil
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
