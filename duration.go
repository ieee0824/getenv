package getenv

import (
	"regexp"
	"time"
	"strconv"
)

var (
	reAlp = regexp.MustCompile(`[a-zA-Z]+`)
	reNum = regexp.MustCompile(`[0-9]+`)
)

func trimArray(a []string) []string {
	ret := []string{}

	for _, e := range a {
		if len(e) != 0 {
			ret = append(ret, e)
		}
	}
	return ret
}

func Duration(s string) time.Duration {
	alp := reAlp.Copy()
	num := reNum.Copy()

	numNodes := trimArray(alp.Split(s, -1))
	alpNodes := trimArray(num.Split(s, -1))

	if len(alpNodes) == 0 {
		if len(numNodes) == 0 {
			return 0
		}

		s, err := strconv.Atoi(numNodes[0])
		if err != nil {
			return 0
		}
		return time.Duration(int64(s)) * time.Second
	}

	if len(alpNodes) != len(numNodes) {
		return 0
	}

	var ret time.Duration
	for i, n := range numNodes {
		t, err := strconv.Atoi(n)
		if err != nil {
			return 0
		}
		d := time.Duration(int64(t))
		switch alpNodes[i] {
		case "h":
			ret += d * time.Hour
		case "m":
			ret += d * time.Minute
		case "s":
			ret += d * time.Second
		default:
			return 0
		}
	}
	return ret
}

