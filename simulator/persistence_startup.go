package simulator

import (
	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// persistenceStartupContext builds and runs the startup persistence restore for
// every persistence-enabled memory of the given devices, using the real
// appliance data root and the real MMA2 Raw Ingest v1 endpoint. It returns the
// observed per-memory outcomes. A build error (for example a device whose rules
// cannot form a plan) is returned so the caller can fail closed; per-memory
// restore failures are reported in each result's deterministic classification
// rather than as a hard error, so one memory's failure does not silently drop
// the others and the memory stays sealed.
func persistenceStartupContext(dataRoot string, devices []DeviceDefinition) ([]mma2composer.PersistenceStartupResult, error) {
	if dataRoot == "" {
		return nil, nil
	}
	var results []mma2composer.PersistenceStartupResult
	for _, device := range devices {
		if !persistenceConfigured(device) {
			continue
		}
		key := mma2composer.PersistenceMemoryKey{Port: device.MMA2.Port, UnitID: device.MMA2.UnitID}
		mem := memoryFromMMA2Params(device.MMA2)
		rules, err := persistenceOwnedRulesForMemory(mem)
		if err != nil {
			return nil, err
		}
		if len(rules) == 0 {
			continue
		}
		configs, err := mma2composer.PersistenceSnapshotConfigs(key, rules)
		if err != nil {
			return nil, err
		}
		adapter, err := mma2composer.NewPersistenceFilesystemAdapter(dataRoot, configs)
		if err != nil {
			return nil, err
		}
		writer := NewPersistenceRawIngestWriter("127.0.0.1", device.MMA2.Port, device.MMA2.UnitID)
		// Persistence-enabled memory starts sealed; restore proceeds only sealed.
		startup := mma2composer.PersistenceStartupMemory{
			Key:     key,
			Enabled: true,
			Sealed:  true,
			Rules:   rules,
			Extra:   mem.Extra,
		}
		result, err := mma2composer.RestorePersistenceAtStartup(startup, adapter, writer)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
