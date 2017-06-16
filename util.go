package getenv

import "strconv"

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}