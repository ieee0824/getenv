package getenv

import "os"

func String(key string, def ...string) string {
	Logger.Dump(key, def)
	var d string
	if len(def) != 0 {
		d = def[0]
	}
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	return v
}
