package simulator

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

func (params MMA2Params) MarshalJSON() ([]byte, error) {
	data, err := yaml.Marshal(params)
	if err != nil {
		return nil, err
	}
	var fields map[string]interface{}
	if err := yaml.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	return json.Marshal(fields)
}

func (params *MMA2Params) UnmarshalJSON(data []byte) error {
	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	encoded, err := yaml.Marshal(fields)
	if err != nil {
		return err
	}
	type plainParams MMA2Params
	return yaml.Unmarshal(encoded, (*plainParams)(params))
}
