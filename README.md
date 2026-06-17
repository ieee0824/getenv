# getenv

English | [日本語](./README.jp.md)

This is a package for getting environment variables with arbitrary types.

[![test](https://github.com/ieee0824/getenv/actions/workflows/test.yml/badge.svg)](https://github.com/ieee0824/getenv/actions/workflows/test.yml)

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
```

# Duration units

`Duration` accepts a default of several types. **Integer defaults (`int`, `int32`, `int64`) and `float64` are interpreted as seconds**, while a `time.Duration` default is used as-is (a nanosecond count). So `getenv.Duration("T", 60)` is 60 seconds, but `getenv.Duration("T", time.Duration(60))` is 60 *nanoseconds* — pass `getenv.Duration("T", 60*time.Second)` instead. String defaults are parsed with `time.ParseDuration` (e.g. `"90s"`, `"1h30m"`); an invalid string or an unsupported default type is logged and falls back to `0`.

# dotenv dump

Setting `GETENV_DUMP_MODE=dotenv` makes every accessor emit a `KEY=` line for each
environment variable it reads, which is handy for generating a `.env` template.

Default values are **masked** (`KEY=`) so that secrets accidentally passed as a default
are not leaked to the dump target. Call `getenv.Logger.DumpValues(true)` to include the
default values, and `getenv.Logger.SetWriter(w)` to redirect the output.