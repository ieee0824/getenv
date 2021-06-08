package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ieee0824/getenv"
)

func main() {
	fmt.Println(getenv.Duration("ANY_ENV"))
	fmt.Println(getenv.Duration("ANY_ENV", 60))
	fmt.Println(getenv.Duration("ANY_ENV", "120s"))
	fmt.Println(getenv.Duration("ANY_ENV", 60*time.Second))
	fmt.Println(getenv.Duration("ANY_ENV", "1h30m20s"))

	os.Setenv("ANY_ENV", "60h")
	fmt.Println(getenv.Duration("ANY_ENV"))
	fmt.Println(getenv.Duration("ANY_ENV", 60))
	fmt.Println(getenv.Duration("ANY_ENV", "120s"))
	fmt.Println(getenv.Duration("ANY_ENV", 60*time.Second))
	fmt.Println(getenv.Duration("ANY_ENV", "1h30m20s"))

	os.Setenv("ANY_ENV", "60h")
	fmt.Println(getenv.Duration("ANY_ENV"))
	fmt.Println(getenv.Duration("ANY_ENV", 60))
	fmt.Println(getenv.Duration("ANY_ENV", "120s"))
	fmt.Println(getenv.Duration("ANY_ENV", 60*time.Second))
	fmt.Println(getenv.Duration("ANY_ENV", "1h30m20s"))

	os.Setenv("GETENV_DUMP_MODE", "dotenv")
	os.Setenv("ANY_ENV", "true")
	getenv.Bool("ANY_ENV", false)
	getenv.Duration("ANY_ENV")
	getenv.Duration("ANY_ENV", 60)
	getenv.Duration("ANY_ENV", "120s")
	getenv.Int("INT_ENV")
	getenv.Int("INT_ENV", -10)
	getenv.String("STR_ENV")
	getenv.String("STR_ENV", "")
	getenv.String("STR_ENV", "hoge")
}
