package getenv

import (
	"time"
	"os"
)

func Duration(key string, def ...interface{}) time.Duration {
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

	v := os.Getenv(key)
	if v == "" {
		return d
	}
	d, _ = time.ParseDuration(v)

	return d
}


