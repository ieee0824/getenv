package getenv

import (
	"strings"
	"os"
)

func convertStringToBoolean(s string) bool {
	s = strings.ToLower(s)
	switch s {
	case "true", "t", "1":
		return true
	default:
		return false
	}
}

func Bool(key string, def ...bool) bool {
	var d bool
	if len(def) != 0 {
		d = def[0]
	} else {
		d = false
	}
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	return convertStringToBoolean(v)
}
