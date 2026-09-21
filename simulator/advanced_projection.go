package simulator

import "fmt"

// projectAdvancedSettings hydrates advanced settings omitted by a Simulator
// device from the matching effective MMA2 memory. Explicit device values win.
func projectAdvancedSettings(params *MMA2Params, cfg EffectiveMMA2Config) error {
	for _, listener := range cfg.Listeners {
		if listenPort(listener.Listen) != params.Port {
			continue
		}
		for _, memory := range listener.Memory {
			if memory.UnitID != params.UnitID {
				continue
			}
			if params.Policy == nil {
				params.Policy = memory.Policy
			}
			for key, value := range memory.Extra {
				switch key {
				case "state_sealing":
					if params.StateSealing != nil {
						continue
					}
					projected, ok := value.(map[string]interface{})
					if !ok {
						return fmt.Errorf("project state_sealing for (%d,%d): expected object, got %T", params.Port, params.UnitID, value)
					}
					params.StateSealing = projected
				case "rbe":
					if params.RBE != nil {
						continue
					}
					projected, ok := value.(map[string]interface{})
					if !ok {
						return fmt.Errorf("project rbe for (%d,%d): expected object, got %T", params.Port, params.UnitID, value)
					}
					params.RBE = projected
				default:
					if params.Extra == nil {
						params.Extra = make(map[string]interface{})
					}
					if _, explicit := params.Extra[key]; !explicit {
						params.Extra[key] = value
					}
				}
			}
			return nil
		}
	}
	return nil
}
