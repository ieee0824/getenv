package getenv

import (
	"testing"
	"time"
	"os"
)

func TestParseDuration(t *testing.T) {
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
		if test.want != parseDuration(test.input) {
			t.Fatalf("want %v, but %v:", test.want, parseDuration(test.input))
		}
	}
}

func TestDuration(t *testing.T) {
	tests := []struct {
		input string
		def interface{}
		want time.Duration
	}{
		{"", "", 0},
		{"", 0, 0},
		{"", nil, 0},
		{"", 60, 60 * time.Second},
		{"", time.Duration(60), time.Duration(60)},
		{"", "60s", 60 * time.Second},
		{"60s", nil, 60*time.Second},
		{"a", nil, 0},
		{"a", 60, 0},
	}

	key := "TEST_DURATION"
	for _, test := range tests {
		os.Setenv(key, test.input)
		get := Duration(key, test.def)
		if test.want != get {
			t.Fatalf("want %v, but %v:", test.want, get)
		}
		os.Unsetenv(key)
	}
}