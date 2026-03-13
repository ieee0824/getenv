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

// env format
// SOME_ENV=a,b,c,d,e,f
func StringSlice(key string, def ...[]string) []string {
	Logger.Dump(key, def)
	var d []string
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(d) == 0 {
			return []string{}
		}
		return d
	}

	return strings.Split(v, ",")
}
