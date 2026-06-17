package getenv

import (
	"os"
	"strings"
)

func String(key string, def ...string) string {
	Logger.Dump(key, def)
	var d string
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	return v
}

// StringSlice resolves key as a comma-separated list, e.g. SOME_ENV=a,b,c.
//
// Each element is trimmed of surrounding whitespace and empty elements are dropped, so
// "a, b," yields ["a", "b"]. An unset or empty value is treated the same: the default
// is returned when present, otherwise an empty (non-nil) slice.
func StringSlice(key string, def ...[]string) []string {
	Logger.Dump(key, def)
	var d []string
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		if len(d) == 0 {
			return []string{}
		}
		return d
	}

	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}
