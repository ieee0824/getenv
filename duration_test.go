package getenv

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		input string
		want time.Duration
	}{
		{"1h0m0s", 1*time.Hour},
		{"1h30m20s", 1*time.Hour + 30*time.Minute + 20*time.Second},
		{"60", 60*time.Second},
		{"", 0},
		{"a", 0},
		{"00af00", 0},
		{"0h30m", 30*time.Minute},
		{"20s", 20*time.Second},
	}

	for _, test := range tests {
		if test.want != Duration(test.input) {
			t.Fatalf("want %v, but %v:", test.want, Duration(test.input))
		}
	}
}