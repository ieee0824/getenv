package getenv

import (
	"log"
	"os"
	"time"
)

// Duration resolves key as a time.Duration.
//
// The optional default may be one of:
//   - int / int32 / int64 — interpreted as a number of SECONDS
//   - float64             — interpreted as a (possibly fractional) number of SECONDS
//   - time.Duration       — used as-is, i.e. a nanosecond count
//     (pass 60*time.Second, NOT time.Duration(60) which is 60ns)
//   - string              — parsed with time.ParseDuration ("90s", "1h30m", ...)
//
// An unparseable string default, or a default of an unsupported type, is logged and
// falls back to 0. When the environment variable is set, its value is parsed with
// time.ParseDuration and falls back to the default (with a log) on error.
func Duration(key string, def ...interface{}) time.Duration {
	Logger.Dump(key, def)

	var d time.Duration
	if len(def) != 0 && def[0] != nil {
		switch v := def[0].(type) {
		case int:
			d = time.Duration(v) * time.Second
		case int32:
			d = time.Duration(v) * time.Second
		case int64:
			d = time.Duration(v) * time.Second
		case float64:
			d = time.Duration(float64(time.Second) * v)
		case time.Duration:
			d = v
		case string:
			parsed, err := time.ParseDuration(v)
			if err != nil {
				log.Printf("getenv: invalid default duration %q for %q: %v", v, key, err)
			} else {
				d = parsed
			}
		default:
			log.Printf("getenv: unsupported default type %T for %q; using 0", def[0], key)
		}
	}

	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	parsed, err := time.ParseDuration(v)
	if err != nil {
		// time.ParseDuration embeds the raw value in its error message, so log only
		// the key to avoid leaking a potentially sensitive environment value.
		log.Printf("getenv: failed to parse value for %q as duration", key)
		return d
	}

	return parsed
}
