# getenv

This is a package for getting environment variables with arbitrary types.

[![CircleCI](https://circleci.com/gh/ieee0824/getenv.svg?style=shield)](https://circleci.com/gh/ieee0824/getenv)

# example


```
import (
	"fmt"
	"github.com/ieee0824/getenv"
	"time"
	"os"
)


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
```

```
0s
1m0s
2m0s
1m0s
1h30m20s
60h0m0s
60h0m0s
60h0m0s
60h0m0s
60h0m0s
18h12m16s
18h12m16s
18h12m16s
18h12m16s
18h12m16s
```