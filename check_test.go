package getenv

import "testing"

func TestExists(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
		want bool
	}{
		{
			name: "key not exists",
			want: false,
		},
		{
			name: "key exists",
			key:  "SOME_ENV",
			val:  "some_value",
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.val != "" {
				t.Setenv(test.key, test.val)
			}

			got := Exists(test.key)
			if got != test.want {
				t.Fatalf("name %s, want %v, but %v:", test.name, test.want, got)
			}
		})
	}
}
