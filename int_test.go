package getenv

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIntParseErrorDoesNotLogRawValue ensures a failed parse falls back to the
// default and never writes the raw environment value to the log (the strconv error
// would otherwise embed it).
func TestIntParseErrorDoesNotLogRawValue(t *testing.T) {
	const raw = "garbage-SECRET-value"
	tests := []struct {
		name string
		typ  string
		call func() interface{}
		want interface{}
	}{
		{"Int", "int", func() interface{} { return Int("NUM_ENV", -10) }, -10},
		{"Int32", "int32", func() interface{} { return Int32("NUM_ENV", -10) }, int32(-10)},
		{"Int64", "int64", func() interface{} { return Int64("NUM_ENV", -10) }, int64(-10)},
		{"Int16", "int16", func() interface{} { return Int16("NUM_ENV", -10) }, int16(-10)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			t.Cleanup(func() { log.SetOutput(os.Stderr) })
			t.Setenv("NUM_ENV", raw)

			got := tt.call()

			assert.Equal(t, tt.want, got, "should fall back to default")
			out := buf.String()
			assert.NotContains(t, out, raw, "raw env value must not be logged")
			assert.Contains(t, out, "NUM_ENV", "key should be logged")
			assert.Contains(t, out, "as "+tt.typ, "type should be logged")
			assert.Contains(t, out, "invalid syntax", "parse reason should be logged")
		})
	}
}
