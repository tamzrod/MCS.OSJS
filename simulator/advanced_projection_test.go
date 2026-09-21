package simulator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func TestAdvancedProjectionMerge(t *testing.T) {
	inheritedPolicy := &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{ID: "inherited"}}}
	matchingExtra := map[string]interface{}{
		"state_sealing": map[string]interface{}{"enabled": false},
		"rbe":           map[string]interface{}{"coils": []interface{}{}},
		"custom":        "inherited",
	}
	cfg := EffectiveMMA2Config{Listeners: []MMA2Listener{
		{Listen: "0.0.0.0:1501", Memory: []MMA2Memory{{UnitID: 1, Extra: map[string]interface{}{"custom": "wrong-port"}}}},
		{Listen: "0.0.0.0:1502", Memory: []MMA2Memory{
			{UnitID: 1, Extra: map[string]interface{}{"custom": "wrong-unit"}},
			{UnitID: 2, Policy: inheritedPolicy, Extra: matchingExtra},
		}},
	}}

	t.Run("projects omitted fields from matching memory", func(t *testing.T) {
		params := MMA2Params{Port: 1502, UnitID: 2}
		if err := projectAdvancedSettings(&params, cfg); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(params.Policy, inheritedPolicy) {
			t.Fatalf("policy = %#v, want %#v", params.Policy, inheritedPolicy)
		}
		if !reflect.DeepEqual(params.StateSealing, matchingExtra["state_sealing"]) {
			t.Fatalf("state_sealing = %#v", params.StateSealing)
		}
		if !reflect.DeepEqual(params.RBE, matchingExtra["rbe"]) {
			t.Fatalf("rbe = %#v", params.RBE)
		}
		if params.Extra["custom"] != "inherited" {
			t.Fatalf("custom extension = %#v", params.Extra["custom"])
		}
	})

	t.Run("preserves explicit fields and extensions", func(t *testing.T) {
		explicitPolicy := &mma2composer.Policy{}
		params := MMA2Params{
			Port:         1502,
			UnitID:       2,
			Policy:       explicitPolicy,
			StateSealing: map[string]interface{}{},
			RBE:          map[string]interface{}{},
			Extra:        map[string]interface{}{"custom": "explicit"},
		}
		if err := projectAdvancedSettings(&params, cfg); err != nil {
			t.Fatal(err)
		}
		if params.Policy != explicitPolicy {
			t.Fatal("explicit policy was replaced")
		}
		if len(params.StateSealing) != 0 || len(params.RBE) != 0 {
			t.Fatalf("explicit empty advanced values changed: sealing=%#v rbe=%#v", params.StateSealing, params.RBE)
		}
		if params.Extra["custom"] != "explicit" {
			t.Fatalf("explicit extension = %#v", params.Extra["custom"])
		}
	})

	t.Run("rejects malformed recognized projection", func(t *testing.T) {
		malformed := EffectiveMMA2Config{Listeners: []MMA2Listener{{
			Listen: "0.0.0.0:1502",
			Memory: []MMA2Memory{{UnitID: 2, Extra: map[string]interface{}{"rbe": true}}},
		}}}
		params := MMA2Params{Port: 1502, UnitID: 2}
		err := projectAdvancedSettings(&params, malformed)
		if err == nil || !strings.Contains(err.Error(), "project rbe") {
			t.Fatalf("malformed recognized projection error = %v", err)
		}
	})
}
