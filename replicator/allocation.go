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

// SuggestDestination returns the first unreserved (port, unit_id) pair,
// starting at the Replicator default destination port.
func (s Store) SuggestDestination() (DestinationSuggestion, error) {
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return DestinationSuggestion{}, err
	}
	occupied := make(map[reservationKey]bool, len(owners.Reservations))
	for _, entry := range owners.Reservations {
		occupied[reservationKey{port: entry.Port, unitID: entry.UnitID}] = true
	}
	port, unit, err := firstAvailable(occupied, DefaultDestinationPort, 0, true, true)
	if err != nil {
		return DestinationSuggestion{}, err
	}
	return DestinationSuggestion{Port: port, UnitID: unit, Owner: ProducerReplicator, Status: "AVAILABLE"}, nil
}

// InspectDestination reports ownership of the exact (port, unit_id) pair.
func (s Store) InspectDestination(port, unitID uint16) (DestinationSuggestion, error) {
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return DestinationSuggestion{}, err
	}
	for _, entry := range owners.Reservations {
		if entry.Port == port && entry.UnitID == unitID {
			status := "OWNED"
			if entry.Owner != ProducerReplicator {
				status = "IN USE"
			}
			return DestinationSuggestion{Port: port, UnitID: unitID, Owner: entry.Owner, Status: status}, nil
		}
	}
	return DestinationSuggestion{Port: port, UnitID: unitID, Owner: ProducerReplicator, Status: "AVAILABLE"}, nil
}

// CheckDocumentOwnership is the authoritative Save & Apply guard. Ownership is
// scoped to the exact (port, unit_id) pair: the same Unit ID on another port is
// valid, and another Unit ID on the same port is also valid.
func (s Store) CheckDocumentOwnership(doc Document) error {
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return err
	}
	foreign := make(map[reservationKey]string)
	for _, entry := range owners.Reservations {
		if entry.Owner == ProducerReplicator {
			continue
		}
		foreign[reservationKey{port: entry.Port, unitID: entry.UnitID}] = entry.Owner
	}
	seen := make(map[reservationKey]string)
	for _, device := range doc.Devices {
		if !device.Enabled {
			continue
		}
		key := reservationKey{port: device.Destination.Port, unitID: device.Destination.UnitID}
		if owner := foreign[key]; owner != "" {
			return fmt.Errorf("device %q: %w: destination (%d,%d) owned by %q", device.Name, mma2composer.ErrReservationOwnedByOther, key.port, key.unitID, owner)
		}
		if other := seen[key]; other != "" {
			return fmt.Errorf("device %q: destination (%d,%d) already assigned to Replicator device %q", device.Name, key.port, key.unitID, other)
		}
		seen[key] = device.Name
	}
	return nil
}

// ResolveDocumentDestinations applies automatic choices while treating foreign
// (port, unit_id) reservations as immutable. Existing Replicator reservations
// are rebuildable by the same producer and become candidates again.
func (s Store) ResolveDocumentDestinations(doc Document) (Document, error) {
	if err := ValidateDocument(doc); err != nil {
		return Document{}, err
	}
	owners, err := mma2composer.New(s.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		return Document{}, err
	}
	occupied := make(map[reservationKey]bool)
	foreignOwner := make(map[reservationKey]string)
	for _, entry := range owners.Reservations {
		if entry.Owner == ProducerReplicator {
			continue
		}
		key := reservationKey{port: entry.Port, unitID: entry.UnitID}
		occupied[key] = true
		foreignOwner[key] = entry.Owner
	}

	resolved := doc
	for i := range resolved.Devices {
		destination := &resolved.Devices[i].Destination
		if !destination.AutoPort && !destination.AutoUnitID {
			manualKey := reservationKey{port: destination.Port, unitID: destination.UnitID}
			if owner := foreignOwner[manualKey]; owner != "" {
				return Document{}, fmt.Errorf("device %q: %w: destination (%d,%d) owned by %q", resolved.Devices[i].Name, mma2composer.ErrReservationOwnedByOther, destination.Port, destination.UnitID, owner)
			}
		}
		port, unit, resolveErr := firstAvailable(
			occupied,
			destination.Port,
			destination.UnitID,
			destination.AutoPort,
			destination.AutoUnitID,
		)
		if resolveErr != nil {
			return Document{}, fmt.Errorf("device %q: %w", resolved.Devices[i].Name, resolveErr)
		}
		key := reservationKey{port: port, unitID: unit}
		if owner := foreignOwner[key]; owner != "" {
			return Document{}, fmt.Errorf("device %q: %w: destination (%d,%d) owned by %q", resolved.Devices[i].Name, mma2composer.ErrReservationOwnedByOther, port, unit, owner)
		}
		if occupied[key] {
			return Document{}, fmt.Errorf("device %q: destination (%d,%d) is already assigned in this Replicator document", resolved.Devices[i].Name, port, unit)
		}
		destination.Port = port
		destination.UnitID = unit
		destination.Owner = ProducerReplicator
		destination.Status = "AVAILABLE"
		occupied[key] = true
	}
	return resolved, nil
}

func firstAvailable(occupied map[reservationKey]bool, requestedPort, requestedUnit uint16, autoPort, autoUnit bool) (uint16, uint16, error) {
	if !autoPort && requestedPort == 0 {
		return 0, 0, fmt.Errorf("manual destination port must be > 0")
	}
	startPort := requestedPort
	if startPort == 0 {
		startPort = DefaultDestinationPort
	}

	if !autoPort && !autoUnit {
		key := reservationKey{port: startPort, unitID: requestedUnit}
		if occupied[key] {
			return 0, 0, fmt.Errorf("destination (%d,%d) is already reserved", startPort, requestedUnit)
		}
		return startPort, requestedUnit, nil
	}

	if !autoPort && autoUnit {
		for unit := uint16(1); unit <= 255; unit++ {
			if !occupied[reservationKey{port: startPort, unitID: unit}] {
				return startPort, unit, nil
			}
		}
		return 0, 0, fmt.Errorf("no free Unit ID remains on port %d", startPort)
	}

	if autoPort && !autoUnit {
		for port := uint32(startPort); port <= 65535; port++ {
			p := uint16(port)
			if !occupied[reservationKey{port: p, unitID: requestedUnit}] {
				return p, requestedUnit, nil
			}
		}
		return 0, 0, errors.New("no free destination port remains")
	}

	for port := uint32(startPort); port <= 65535; port++ {
		p := uint16(port)
		for unit := uint16(1); unit <= 255; unit++ {
			if !occupied[reservationKey{port: p, unitID: unit}] {
				return p, unit, nil
			}
		}
	}
	return 0, 0, errors.New("no free destination reservation remains")
}
