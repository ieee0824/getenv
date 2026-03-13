package getenv

import (
	"os"
	"time"
)

func Duration(key string, def ...interface{}) time.Duration {
	Logger.Dump(key, def)
	var d time.Duration
	if len(def) != 0 {
		if i, ok := def[0].(int); ok {
			d = time.Duration(int64(i)) * time.Second
		} else if dr, ok := def[0].(time.Duration); ok {
			d = dr
		} else if s, ok := def[0].(string); ok {
			d, _ = time.ParseDuration(s)
		}
	}

	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	parsed, err := time.ParseDuration(v)
	if err != nil {
		return d
	}

	return parsed
}
