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
	v := os.Getenv(key)
	if v == "" {
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
	v := os.Getenv(key)
	if v == "" && len(d) == 0 {
		return []string{}
	} else if v == "" && len(d) != 0 {
		return d
	}

	return strings.Split(v, ",")
}
