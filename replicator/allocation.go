package replicator

import (
	"errors"
	"fmt"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

type reservationKey struct {
	port   uint16
	unitID uint16
}

// DestinationSuggestion is the small ownership-aware result exposed to the UI.
type DestinationSuggestion struct {
	Port   uint16 `json:"port"`
	UnitID uint16 `json:"unit_id"`
	Owner  string `json:"owner"`
	Status string `json:"status"`
}

// SuggestDestination returns the first destination whose port and Unit ID are
// both unused in the shared MMA2 ownership document.
func (s Store) SuggestDestination() (DestinationSuggestion, error) {
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return DestinationSuggestion{}, err
	}
	occupiedPairs := make(map[reservationKey]bool, len(owners.Reservations))
	occupiedPorts := make(map[uint16]bool, len(owners.Reservations))
	occupiedUnits := make(map[uint16]bool, len(owners.Reservations))
	for _, entry := range owners.Reservations {
		occupiedPairs[reservationKey{port: entry.Port, unitID: entry.UnitID}] = true
		occupiedPorts[entry.Port] = true
		occupiedUnits[entry.UnitID] = true
	}
	port, unit, err := firstAvailable(occupiedPairs, occupiedPorts, occupiedUnits, DefaultDestinationPort, 0, true, true)
	if err != nil {
		return DestinationSuggestion{}, err
	}
	return DestinationSuggestion{Port: port, UnitID: unit, Owner: ProducerReplicator, Status: "AVAILABLE"}, nil
}

// InspectDestination reports current shared ownership. A destination is IN USE
// when either its port or its Unit ID belongs to another producer.
func (s Store) InspectDestination(port, unitID uint16) (DestinationSuggestion, error) {
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return DestinationSuggestion{}, err
	}
	for _, entry := range owners.Reservations {
		if entry.Port == port && entry.Owner != ProducerReplicator {
			return DestinationSuggestion{Port: port, UnitID: unitID, Owner: entry.Owner, Status: "IN USE"}, nil
		}
	}
	for _, entry := range owners.Reservations {
		if entry.UnitID == unitID && entry.Owner != ProducerReplicator {
			return DestinationSuggestion{Port: port, UnitID: unitID, Owner: entry.Owner, Status: "IN USE"}, nil
		}
	}
	for _, entry := range owners.Reservations {
		if entry.Port == port && entry.UnitID == unitID {
			return DestinationSuggestion{Port: port, UnitID: unitID, Owner: entry.Owner, Status: "OWNED"}, nil
		}
	}
	return DestinationSuggestion{Port: port, UnitID: unitID, Owner: ProducerReplicator, Status: "AVAILABLE"}, nil
}

// CheckDocumentOwnership is the authoritative Save & Apply guard. Port and
// Unit ID are independently exclusive against foreign producers.
func (s Store) CheckDocumentOwnership(doc Document) error {
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return err
	}
	foreignPortOwner := make(map[uint16]string)
	foreignUnitOwner := make(map[uint16]string)
	for _, entry := range owners.Reservations {
		if entry.Owner == ProducerReplicator {
			continue
		}
		foreignPortOwner[entry.Port] = entry.Owner
		foreignUnitOwner[entry.UnitID] = entry.Owner
	}
	seenPorts := make(map[uint16]string)
	seenUnits := make(map[uint16]string)
	for _, device := range doc.Devices {
		if !device.Enabled {
			continue
		}
		if owner := foreignPortOwner[device.Destination.Port]; owner != "" {
			return fmt.Errorf("device %q: %w: destination port %d owned by %q", device.Name, mma2composer.ErrReservationOwnedByOther, device.Destination.Port, owner)
		}
		if owner := foreignUnitOwner[device.Destination.UnitID]; owner != "" {
			return fmt.Errorf("device %q: %w: destination Unit ID %d owned by %q", device.Name, mma2composer.ErrReservationOwnedByOther, device.Destination.UnitID, owner)
		}
		if other := seenPorts[device.Destination.Port]; other != "" {
			return fmt.Errorf("device %q: destination port %d already assigned to Replicator device %q", device.Name, device.Destination.Port, other)
		}
		if other := seenUnits[device.Destination.UnitID]; other != "" {
			return fmt.Errorf("device %q: destination Unit ID %d already assigned to Replicator device %q", device.Name, device.Destination.UnitID, other)
		}
		seenPorts[device.Destination.Port] = device.Name
		seenUnits[device.Destination.UnitID] = device.Name
	}
	return nil
}

// ResolveDocumentDestinations applies automatic choices while treating foreign
// port and Unit ID ownership as immutable. Existing Replicator reservations are
// rebuildable by this same producer and are reserved again from the edited doc.
func (s Store) ResolveDocumentDestinations(doc Document) (Document, error) {
	if err := ValidateDocument(doc); err != nil {
		return Document{}, err
	}
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return Document{}, err
	}
	occupiedPairs := make(map[reservationKey]bool)
	occupiedPorts := make(map[uint16]bool)
	occupiedUnits := make(map[uint16]bool)
	foreignPortOwner := make(map[uint16]string)
	foreignUnitOwner := make(map[uint16]string)
	for _, entry := range owners.Reservations {
		if entry.Owner == ProducerReplicator {
			continue
		}
		occupiedPairs[reservationKey{port: entry.Port, unitID: entry.UnitID}] = true
		occupiedPorts[entry.Port] = true
		occupiedUnits[entry.UnitID] = true
		foreignPortOwner[entry.Port] = entry.Owner
		foreignUnitOwner[entry.UnitID] = entry.Owner
	}

	resolved := doc
	for i := range resolved.Devices {
		destination := &resolved.Devices[i].Destination
		if !destination.AutoPort {
			if owner := foreignPortOwner[destination.Port]; owner != "" {
				return Document{}, fmt.Errorf("device %q: %w: destination port %d owned by %q", resolved.Devices[i].Name, mma2composer.ErrReservationOwnedByOther, destination.Port, owner)
			}
		}
		if !destination.AutoUnitID {
			if owner := foreignUnitOwner[destination.UnitID]; owner != "" {
				return Document{}, fmt.Errorf("device %q: %w: destination Unit ID %d owned by %q", resolved.Devices[i].Name, mma2composer.ErrReservationOwnedByOther, destination.UnitID, owner)
			}
		}
		port, unit, resolveErr := firstAvailable(
			occupiedPairs,
			occupiedPorts,
			occupiedUnits,
			destination.Port,
			destination.UnitID,
			destination.AutoPort,
			destination.AutoUnitID,
		)
		if resolveErr != nil {
			return Document{}, fmt.Errorf("device %q: %w", resolved.Devices[i].Name, resolveErr)
		}
		destination.Port = port
		destination.UnitID = unit
		destination.Owner = ProducerReplicator
		destination.Status = "AVAILABLE"
		occupiedPairs[reservationKey{port: port, unitID: unit}] = true
		occupiedPorts[port] = true
		occupiedUnits[unit] = true
	}
	return resolved, nil
}

func firstAvailable(occupiedPairs map[reservationKey]bool, occupiedPorts map[uint16]bool, occupiedUnits map[uint16]bool, requestedPort, requestedUnit uint16, autoPort, autoUnit bool) (uint16, uint16, error) {
	if !autoPort && requestedPort == 0 {
		return 0, 0, fmt.Errorf("manual destination port must be > 0")
	}
	startPort := requestedPort
	if startPort == 0 {
		startPort = DefaultDestinationPort
	}

	if !autoPort && occupiedPorts[startPort] {
		return 0, 0, fmt.Errorf("destination port %d is already reserved", startPort)
	}
	if !autoUnit && occupiedUnits[requestedUnit] {
		return 0, 0, fmt.Errorf("destination Unit ID %d is already reserved", requestedUnit)
	}
	if !autoPort && !autoUnit {
		key := reservationKey{port: startPort, unitID: requestedUnit}
		if occupiedPairs[key] {
			return 0, 0, fmt.Errorf("destination (%d,%d) is already reserved", startPort, requestedUnit)
		}
		return startPort, requestedUnit, nil
	}

	if !autoPort && autoUnit {
		for unit := uint16(1); unit <= 255; unit++ {
			if !occupiedUnits[unit] && !occupiedPairs[reservationKey{port: startPort, unitID: unit}] {
				return startPort, unit, nil
			}
		}
		return 0, 0, fmt.Errorf("no free Unit ID remains for port %d", startPort)
	}

	if autoPort && !autoUnit {
		for port := uint32(startPort); port <= 65535; port++ {
			p := uint16(port)
			if occupiedPorts[p] {
				continue
			}
			if !occupiedPairs[reservationKey{port: p, unitID: requestedUnit}] {
				return p, requestedUnit, nil
			}
		}
		return 0, 0, errors.New("no free destination port remains")
	}

	for port := uint32(startPort); port <= 65535; port++ {
		p := uint16(port)
		if occupiedPorts[p] {
			continue
		}
		for unit := uint16(1); unit <= 255; unit++ {
			if occupiedUnits[unit] {
				continue
			}
			if !occupiedPairs[reservationKey{port: p, unitID: unit}] {
				return p, unit, nil
			}
		}
	}
	return 0, 0, errors.New("no free destination reservation remains")
}
