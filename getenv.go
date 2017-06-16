package getenv

import (
	"os"
	"strings"
	"strconv"
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

func String(key string, def ...string) string {
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

func Int(key string, def ...int) int {
	var d int
	if len(def) != 0 {
		d = def[0]
	}
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return d
	}
	return i
}

