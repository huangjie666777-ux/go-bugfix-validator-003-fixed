package validator

import (
	"strings"
	"testing"
)

func TestFQDNLengthLimit(t *testing.T) {
	maxFQDN := strings.Repeat(strings.Repeat("a", 63)+".", 3) + strings.Repeat("a", 61)

	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{name: "maximum length", value: maxFQDN, valid: true},
		{name: "maximum length with root dot", value: maxFQDN + ".", valid: true},
		{name: "overlong without root dot", value: maxFQDN + "a", valid: false},
		{name: "overlong with root dot", value: maxFQDN + "a.", valid: false},
	}

	validate := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.value, "fqdn")
			if tt.valid && err != nil {
				t.Fatalf("expected %d-character FQDN to be valid: %v", len(tt.value), err)
			}
			if !tt.valid && err == nil {
				t.Fatalf("expected %d-character FQDN to be rejected", len(tt.value))
			}
		})
	}
}
