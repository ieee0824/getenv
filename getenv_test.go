package getenv

import (
	"testing"
	"os"
)

func TestConvertStringToBoolean(t *testing.T) {
	tests := []struct {
		input string
		want bool
	}{
		{"true", true},
		{"t", true},
		{"True", true},
		{"TRUE", true},
		{"T", true},
		{"1", true},
		{"false", false},
		{"f", false},
		{"False", false},
		{"FALSE", false},
		{"0", false},
		{"hoge", false},
	}

	for _, test := range tests {
		got := convertStringToBoolean(test.input)
		if got != test.want {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}

