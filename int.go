package getenv

import (
	"errors"
	"log"
	"os"
	"strconv"
)

func Int(key string, def ...int) int {
	Logger.Dump(key, def)
	var d int
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("getenv: failed to parse %q as int: %v", key, parseErrReason(err))
		return d
	}
	return i
}

func Int32(key string, def ...int32) int32 {
	Logger.Dump(key, def)
	var d int32
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	i32, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Printf("getenv: failed to parse %q as int32: %v", key, parseErrReason(err))
		return d
	}
	return int32(i32)
}

func Int64(key string, def ...int64) int64 {
	Logger.Dump(key, def)
	var d int64
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("getenv: failed to parse %q as int64: %v", key, parseErrReason(err))
		return d
	}
	return int64(i64)
}

func Int16(key string, def ...int16) int16 {
	Logger.Dump(key, def)
	var d int16
	if len(def) != 0 {
		d = def[0]
	}
	v, ok := os.LookupEnv(key)
	if !ok {
		return d
	}
	i16, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		log.Printf("getenv: failed to parse %q as int16: %v", key, parseErrReason(err))
		return d
	}
	return int16(i16)
}

// parseErrReason returns the reason of a strconv error (e.g. "invalid syntax" or
// "value out of range") without the raw input value, which strconv.NumError would
// otherwise embed in its message. This keeps the (potentially sensitive) environment
// value out of the logs.
func parseErrReason(err error) error {
	var ne *strconv.NumError
	if errors.As(err, &ne) {
		return ne.Err
	}
	return err
}
