package getenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice(t *testing.T) {
	tests := []struct {
		name   string
		def    []string
		want   []string
		envKey string
		envVal string
	}{
		{
			name: "empty",
			want: []string{},
		},
		{
			name: "val empty, def not empty",
			def:  []string{"a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name:   "val not empty, def empty",
			envKey: "SOME_ENV",
			envVal: "a,b,c",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "val not empty, def not empty",
			envKey: "SOME_ENV",
			envVal: "a,b,c",
			def:    []string{"d", "e", "f"},
			want:   []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.envKey != "" {
				t.Setenv(test.envKey, test.envVal)
			}

			got := StringSlice(test.envKey, test.def)
			assert.Equal(t, test.want, got)
		})
	}
}
