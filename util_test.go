package getenv

import "testing"

func TestIsNumber(t *testing.T) {
	if !isNumber("1") {
		t.Fatalf("want %v, but %v:", true, isNumber("1"))
	} else if isNumber("a") {
		t.Fatalf("want %v, but %v:", false, isNumber("a"))
	}
}