package simulator

import (
	"fmt"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// bootstrapPersistenceTransitions captures a complete initial snapshot from the
// currently running, unsealed memory before the first persistence-enabled
// structural restart. This prevents the first restart from entering SEALED state
// with no restorable snapshot.
//
// It runs only for an OFF -> ON persistence transition. To avoid silently
// migrating memory identity/layout during bootstrap, the existing running device
// must have the same name, Port/UnitID and FC1-FC4 ranges as the edited device.
// Any read/write failure aborts Save & Apply before compose/restart.
func bootstrapPersistenceTransitions(dataRoot string, previous, edited Document) error {
	if dataRoot == "" {
		return fmt.Errorf("appliance data root is unavailable")
	}

	beforeByName := make(map[string]DeviceDefinition, len(previous.Devices))
	for _, d := range previous.Devices {
		beforeByName[d.Name] = d
	}

	for _, after := range edited.Devices {
		if !persistenceConfigured(after) {
			continue
		}
		before, ok := beforeByName[after.Name]
		if ok && persistenceConfigured(before) {
			continue
		}
		if !ok {
			return fmt.Errorf("device %q cannot enable persistence before it has a running memory to snapshot", after.Name)
		}
		if !before.Enabled {
			return fmt.Errorf("device %q cannot enable persistence while its previous runtime is disabled", after.Name)
		}
		if before.MMA2.Port != after.MMA2.Port || before.MMA2.UnitID != after.MMA2.UnitID {
			return fmt.Errorf("device %q cannot change Port/Unit ID while enabling persistence; enable persistence first, then change identity separately", after.Name)
		}
		if before.MMA2.FC1 != after.MMA2.FC1 || before.MMA2.FC2 != after.MMA2.FC2 ||
			before.MMA2.FC3 != after.MMA2.FC3 || before.MMA2.FC4 != after.MMA2.FC4 {
			return fmt.Errorf("device %q cannot change memory ranges while enabling persistence; enable persistence first, then change ranges separately", after.Name)
		}

		key := mma2composer.PersistenceMemoryKey{Port: after.MMA2.Port, UnitID: after.MMA2.UnitID}
		rules, err := persistenceOwnedRulesForMemory(memoryFromMMA2Params(after.MMA2))
		if err != nil {
			return fmt.Errorf("device %q derive persistence areas: %w", after.Name, err)
		}
		configs, err := mma2composer.PersistenceSnapshotConfigs(key, rules)
		if err != nil {
			return fmt.Errorf("device %q build snapshot config: %w", after.Name, err)
		}
		if len(configs) == 0 {
			return fmt.Errorf("device %q has persistence enabled but no configured memory area to snapshot", after.Name)
		}
		adapter, err := mma2composer.NewPersistenceFilesystemAdapter(dataRoot, configs)
		if err != nil {
			return fmt.Errorf("device %q open snapshot store: %w", after.Name, err)
		}
		reader := persistenceAreaReader(before)
		for _, cfg := range configs {
			payload, err := reader.ReadPersistenceArea(key, cfg.Area, cfg.Kind, cfg.Start, cfg.Count)
			if err != nil {
				return fmt.Errorf("device %q capture initial %s snapshot: %w", after.Name, cfg.Area, err)
			}
			if err := adapter.WritePersistenceBytes(key, cfg.Area, 0, payload); err != nil {
				return fmt.Errorf("device %q persist initial %s snapshot: %w", after.Name, cfg.Area, err)
			}
		}
	}
	return nil
}
