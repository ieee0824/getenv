# getenv

[English](./README.md) | 日本語

Go言語で環境変数を取得時に任意の型に変換して取得することができます.  
また環境変数に値が存在しないときはデフォルト値を設定しておいて取得することができます.

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

# dotenv ダンプ

`GETENV_DUMP_MODE=dotenv` を設定すると、各アクセサが読み取った環境変数ごとに `KEY=` 行を出力します. `.env` テンプレートの生成に便利です.

デフォルト値は既定で**マスク**されます(`KEY=`). デフォルト値に誤って秘密情報を渡してもダンプ先に漏れないようにするためです. デフォルト値も出力したい場合は `getenv.Logger.DumpValues(true)` を呼び出してください. 出力先は `getenv.Logger.SetWriter(w)` で変更できます.
