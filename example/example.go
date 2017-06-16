package main

import (
	"fmt"
	"github.com/ieee0824/getenv"
	"time"
	"os"
)

func main() {
	fmt.Println(getenv.Duration("ANY_ENV"))
	fmt.Println(getenv.Duration("ANY_ENV", 60))
	fmt.Println(getenv.Duration("ANY_ENV","120s"))
	fmt.Println(getenv.Duration("ANY_ENV", 60 * time.Second))
	fmt.Println(getenv.Duration("ANY_ENV", "1h30m20s"))

	os.Setenv("ANY_ENV", "60h")
	fmt.Println(getenv.Duration("ANY_ENV"))
	fmt.Println(getenv.Duration("ANY_ENV", 60))
	fmt.Println(getenv.Duration("ANY_ENV","120s"))
	fmt.Println(getenv.Duration("ANY_ENV", 60 * time.Second))
	fmt.Println(getenv.Duration("ANY_ENV", "1h30m20s"))

	os.Setenv("ANY_ENV", "60h")
	fmt.Println(getenv.Duration("ANY_ENV"))
	fmt.Println(getenv.Duration("ANY_ENV", 60))
	fmt.Println(getenv.Duration("ANY_ENV","120s"))
	fmt.Println(getenv.Duration("ANY_ENV", 60 * time.Second))
	fmt.Println(getenv.Duration("ANY_ENV", "1h30m20s"))
}
