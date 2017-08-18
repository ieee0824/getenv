package getenv

import (
	"os"
	"strconv"
	"log"
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

func Int32(key string, def ...int32) (int32) {
	var d int32
	if len(def) != 0 {
		d = def[0]
	}
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	i32, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Printf("parse error: input: %v, %v\n", v, err.Error())
		return d
	}
	return int32(i32)
}

func Int64(key string, def ...int64) (int64) {
	var d int64
	if len(def) != 0 {
		d = def[0]
	}
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("parse error: input: %v, %v\n", v, err.Error())
		return d
	}
	return int64(i64)
}

func Int16(key string, def ...int16) (int16) {
	var d int16
	if len(def) != 0 {
		d = def[0]
	}
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	i16, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		log.Printf("parse error: input: %v, %v\n", v, err.Error())
		return d
	}
	return int16(i16)
}