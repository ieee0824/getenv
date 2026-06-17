package getenv

import (
	"log"
	"os"
	"strings"
)

// convertStringToBoolean parses common boolean representations, case-insensitively
// and with surrounding whitespace trimmed. The second return value reports whether s
// was a recognised boolean; for unrecognised input it is false so the caller can
// decide how to proceed (e.g. keep a default).
func convertStringToBoolean(s string) (bool, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "t", "1", "yes", "on", "y":
		return true, true
	case "false", "f", "0", "no", "off", "n":
		return false, true
	default:
		return false, false
	}
}

// Bool resolves key as a boolean. Recognised truthy values are true/t/1/yes/on/y and
// falsy values are false/f/0/no/off/n (case-insensitive, trimmed). If the value is set
// but unrecognised, the default is kept (and a warning is logged) rather than silently
// returning false.
func Bool(key string, def ...bool) bool {
	Logger.Dump(key, def)
	var d bool
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	b, recognised := convertStringToBoolean(v)
	if !recognised {
		// Do not log the raw value, which may be sensitive; the key is enough.
		log.Printf("getenv: unrecognised bool for %q; keeping default %v", key, d)
		return d
	}
	return b
}
