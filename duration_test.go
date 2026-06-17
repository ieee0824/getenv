package getenv

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		input string
		def   interface{}
		want  time.Duration
	}{
		{"", "", 0},
		{"", 0, 0},
		{"", nil, 0},
		{"", 60, 60 * time.Second},
		{"", time.Duration(60), time.Duration(60)},
		{"", "60s", 60 * time.Second},
		{"60s", nil, 60 * time.Second},
		{"a", nil, 0},
		{"a", 60, 60 * time.Second},
		// numeric default types are interpreted as seconds
		{"", int64(60), 60 * time.Second},
		{"", int32(60), 60 * time.Second},
		{"", float64(1.5), 1500 * time.Millisecond},
		// invalid string default falls back to 0 (logged, not silent panic)
		{"", "30", 0},
		{"", "30sec", 0},
		// unsupported default type falls back to 0
		{"", true, 0},
		{"", uint(60), 0},
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

// TestDurationEnvParseErrorDoesNotLogRawValue ensures an invalid env value falls back
// to the default without leaking the raw value into the log (time.ParseDuration would
// otherwise embed it in its error message).
func TestDurationEnvParseErrorDoesNotLogRawValue(t *testing.T) {
	const raw = "garbage-SECRET-duration"
	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })
	t.Setenv("DUR_ENV", raw)

	got := Duration("DUR_ENV", 5*time.Second)

	assert.Equal(t, 5*time.Second, got, "should fall back to default")
	assert.NotContains(t, buf.String(), raw, "raw env value must not be logged")
	assert.Contains(t, buf.String(), "DUR_ENV", "key should be logged")
}
