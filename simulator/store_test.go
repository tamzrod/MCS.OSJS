package simulator

import (
	"os"
	"path/filepath"
	"testing"
)

func validDevice() DeviceDefinition {
	return DeviceDefinition{
		Name:    "Sim-PLC-1",
		Enabled: true,
		MMA2: MMA2Params{
			Port:   5020,
			UnitID: 1,
			FC1:    Area{Start: 0, Count: 64},
			FC2:    Area{Start: 0, Count: 64},
			FC3:    Area{Start: 0, Count: 100},
			FC4:    Area{Start: 0, Count: 100},
		},
		RandomRuntime: RandomRuntimeParams{
			FC1IntervalMS: 1000,
			FC2IntervalMS: 5000,
			FC3IntervalMS: 60000,
			FC4IntervalMS: 10000,
		},
	}
}

func TestConfigRootFromEnvRefusesInventedPath(t *testing.T) {
	t.Setenv("OSJS_DATA_DIR", "")
	if _, err := ConfigRootFromEnv(); err == nil {
		t.Fatal("expected error when OSJS_DATA_DIR is unset")
	}
}

func TestRoundTripExactValues(t *testing.T) {
	root := t.TempDir()
	t.Setenv("OSJS_DATA_DIR", root)
	gotRoot, err := ConfigRootFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if gotRoot != root {
		t.Fatalf("root %q != %q", gotRoot, root)
	}

	s := Store{Root: gotRoot}
	want := validDevice()
	if err := s.SaveOne(want); err != nil {
		t.Fatal(err)
	}

	// Persistence is under the verified mount, not MMA2/ or OSJS source.
	path := s.DevicesPath()
	rel := "config/simulator/devices.yaml"
	if path != filepath.Join(root, rel) {
		t.Fatalf("devices path %q is not under mount %s/%s", path, root, rel)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}

	doc, err := s.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Devices) != 1 {
		t.Fatalf("got %d devices", len(doc.Devices))
	}
	got := doc.Devices[0]
	if got != want {
		t.Fatalf("round-trip mismatch\nwant %#v\ngot  %#v", want, got)
	}
}

func TestInvalidInputsDoNotReplaceLastValid(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	want := validDevice()
	if err := s.SaveOne(want); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(s.DevicesPath())
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name string
		mut  func(*DeviceDefinition)
	}{
		{"port 0", func(d *DeviceDefinition) { d.MMA2.Port = 0 }},
		{"unit_id 256", func(d *DeviceDefinition) { d.MMA2.UnitID = 256 }},
		{"fc3 start+count overflow", func(d *DeviceDefinition) { d.MMA2.FC3 = Area{Start: 65530, Count: 20} }},
		{"fc1 interval 0", func(d *DeviceDefinition) { d.RandomRuntime.FC1IntervalMS = 0 }},
		{"empty name", func(d *DeviceDefinition) { d.Name = "  " }},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bad := validDevice()
			tc.mut(&bad)
			if err := s.SaveOne(bad); err == nil {
				t.Fatalf("expected validation error for %s", tc.name)
			}
			after, err := os.ReadFile(s.DevicesPath())
			if err != nil {
				t.Fatal(err)
			}
			if string(after) != string(before) {
				t.Fatalf("invalid %s replaced last valid definition", tc.name)
			}
			doc, err := s.Load()
			if err != nil {
				t.Fatal(err)
			}
			if doc.Devices[0] != want {
				t.Fatalf("loaded definition changed after rejected %s", tc.name)
			}
		})
	}
}
