package mma2composer

import (
	"strings"
	"testing"
)

func TestFC43EveryExplicitFieldUsesASCIIByteLimit(t *testing.T) {
	for _, field := range []string{"vendor_name", "product_code", "major_minor_revision"} {
		for _, invalid := range []interface{}{"", "é", strings.Repeat("a", 245), nil, 123} {
			if err := ValidateFC43(map[string]interface{}{field: invalid}); err == nil {
				t.Fatalf("%s accepted %#v", field, invalid)
			}
		}
		if err := ValidateFC43(map[string]interface{}{field: strings.Repeat("a", 244)}); err != nil {
			t.Fatal(err)
		}
	}
	if err := ValidateFC43(map[string]interface{}{}); err != nil {
		t.Fatal(err)
	}
	if err := ValidateFC43(nil); err != nil {
		t.Fatal(err)
	}
}
