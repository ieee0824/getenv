package getenv

import "os"

func Exists(key string) bool {
	_, ok := os.LookupEnv(key)

	return ok
}
