package getenv

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToBoolean(t *testing.T) {
	tests := []struct {
		input      string
		want       bool
		wantParsed bool
	}{
		{"true", true, true},
		{"t", true, true},
		{"True", true, true},
		{"TRUE", true, true},
		{"T", true, true},
		{"1", true, true},
		{"yes", true, true},
		{"on", true, true},
		{"y", true, true},
		{"  TRUE  ", true, true},
		{"false", false, true},
		{"f", false, true},
		{"False", false, true},
		{"FALSE", false, true},
		{"0", false, true},
		{"no", false, true},
		{"off", false, true},
		{"n", false, true},
		// unrecognised values report parsed=false
		{"hoge", false, false},
		{"2", false, false},
		{"", false, false},
	}

	for _, test := range tests {
		got, parsed := convertStringToBoolean(test.input)
		assert.Equal(t, test.want, got, "value for input %q", test.input)
		assert.Equal(t, test.wantParsed, parsed, "parsed for input %q", test.input)
	}
}

func TestBoolUnrecognisedKeepsDefault(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	t.Setenv("FLAG", "maybe-SECRET")

	assert.True(t, Bool("FLAG", true), "unrecognised value should keep default true")
	assert.False(t, Bool("FLAG", false), "unrecognised value should keep default false")

	out := buf.String()
	assert.Contains(t, out, "FLAG", "key should be logged")
	assert.NotContains(t, out, "maybe-SECRET", "raw value must not be logged")
}

func TestBoolExplicitOverridesDefault(t *testing.T) {
	t.Setenv("FLAG", "off")
	assert.False(t, Bool("FLAG", true), "explicit falsy should override default true")

	t.Setenv("FLAG", "yes")
	assert.True(t, Bool("FLAG", false), "explicit truthy should override default false")
}
