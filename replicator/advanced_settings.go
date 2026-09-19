package replicator

import (
	"fmt"
	"net"
	"strconv"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"gopkg.in/yaml.v3"
)

func cloneAdvancedValue(value interface{}) interface{} {
	switch typed := value.(type) {
	case map[string]interface{}:
		copied := make(map[string]interface{}, len(typed))
		for key, item := range typed {
			copied[key] = cloneAdvancedValue(item)
		}
		return copied
	case []interface{}:
		copied := make([]interface{}, len(typed))
		for index, item := range typed {
			copied[index] = cloneAdvancedValue(item)
		}
		return copied
	default:
		return value
	}
}

func inheritAdvanced(device *DeviceDefinition, cfg mma2composer.EffectiveConfig) error {
	for _, listener := range cfg.Listeners {
		_, port, err := net.SplitHostPort(listener.Listen)
		if err != nil || port != strconv.Itoa(int(device.Destination.Port)) {
			continue
		}
		for _, memory := range listener.Memory {
			if memory.UnitID != device.Destination.UnitID {
				continue
			}
			body, err := yaml.Marshal(memory)
			if err != nil {
				return err
			}
			var saved map[string]interface{}
			if err := yaml.Unmarshal(body, &saved); err != nil {
				return err
			}
			for _, key := range []string{"policy", "state_sealing", "rbe"} {
				if _, explicit := device.MMA2Advanced[key]; explicit {
					continue
				}
				if value, exists := saved[key]; exists {
					if device.MMA2Advanced == nil {
						device.MMA2Advanced = map[string]interface{}{}
					}
					device.MMA2Advanced[key] = value
				}
			}
			return nil
		}
	}
	return nil
}

func applyAdvanced(memory *mma2composer.Memory, advanced map[string]interface{}) error {
	for key := range advanced {
		if key != "policy" && key != "state_sealing" && key != "rbe" {
			return fmt.Errorf("unsupported MMA2 advanced setting %q", key)
		}
	}
	body, err := yaml.Marshal(advanced)
	if err != nil {
		return err
	}
	var settings mma2composer.Memory
	if err := yaml.Unmarshal(body, &settings); err != nil {
		return err
	}
	if _, exists := advanced["policy"]; exists {
		memory.Policy = settings.Policy
	}
	memory.Extra = settings.Extra
	return nil
}
