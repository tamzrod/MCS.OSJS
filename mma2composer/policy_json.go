package mma2composer

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

func (policy Policy) MarshalJSON() ([]byte, error) {
	data, err := yaml.Marshal(policy)
	if err != nil {
		return nil, err
	}
	var fields map[string]interface{}
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	return json.Marshal(fields)
}

func (policy *Policy) UnmarshalJSON(data []byte) error {
	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	encoded, err := yaml.Marshal(fields)
	if err != nil {
		return err
	}
	type plainPolicy Policy
	return yaml.Unmarshal(encoded, (*plainPolicy)(policy))
}
