package getenv

import (
	"testing"
	"time"
	"os"
)

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