package configvalidate

import (
	"strings"
	"testing"
)

const rbeFixture = `rbe:
  tcp:
    listen: "127.0.0.1:19001"
listeners:
  - id: lab
    listen: "127.0.0.1:1502"
    memory:
      - unit_id: 1
        holding_registers: {start: 0, count: 10}
        rbe:
          holding_registers:
            - {id: 1, name: watched, start: 0, count: 1}
`

func TestYAMLRBEValidation(t *testing.T) {
	if err := YAML([]byte(rbeFixture)); err != nil {
		t.Fatal(err)
	}
	for _, fixture := range []string{
		strings.Replace(rbeFixture, "id: 1,", "id: 0,", 1),
		strings.Replace(rbeFixture, "id: 1,", "id: 256,", 1),
		strings.Replace(rbeFixture, "name: watched, start: 0", "name: watched, start: 10", 1),
		strings.Replace(rbeFixture, "19001", "1502", 1),
		rbeFixture + "notify: {}\n",
		rbeFixture + "notify: {influx: {url: http://localhost}}\n",
		strings.Replace(rbeFixture, "  tcp:", "  influx: {}\n  tcp:", 1),
		strings.TrimPrefix(rbeFixture, "rbe:\n  tcp:\n    listen: \"127.0.0.1:19001\"\n"),
	} {
		if err := YAML([]byte(fixture)); err == nil {
			t.Fatalf("invalid configuration accepted: %s", fixture)
		}
	}
}
