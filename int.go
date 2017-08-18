package getenv

import (
	"os"
	"strconv"
)

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
