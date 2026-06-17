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

# Duration の単位

`Duration` のデフォルト値は複数の型を受け付けます. **整数(`int`/`int32`/`int64`)と `float64` は「秒」として解釈**され, `time.Duration` はそのまま(ナノ秒)使われます. そのため `getenv.Duration("T", 60)` は 60 秒ですが, `getenv.Duration("T", time.Duration(60))` は 60 *ナノ秒* です(秒で渡すなら `getenv.Duration("T", 60*time.Second)`). 文字列は `time.ParseDuration` でパースされます(例: `"90s"`, `"1h30m"`). 不正な文字列や未対応の型はログ出力して `0` にフォールバックします.

# Bool

`Bool` は `true`/`t`/`1`/`yes`/`on`/`y` を真, `false`/`f`/`0`/`no`/`off`/`n` を偽として解釈します(大文字小文字無視・前後空白 trim). 値がセットされていても認識できない場合は, 無音で false を返さずデフォルト値を維持します(警告ログを出力).

# StringSlice

`StringSlice` はカンマ区切りの値(例: `SOME_ENV=a,b,c`)を分割し, 各要素を trim して空要素を除去します(`"a, b,"` → `["a", "b"]`). 未設定または空文字列のときは, デフォルトがあればそれを, 無ければ空スライスを返します.

# dotenv ダンプ

`GETENV_DUMP_MODE=dotenv` を設定すると、各アクセサが読み取った環境変数ごとに `KEY=` 行を出力します. `.env` テンプレートの生成に便利です.

デフォルト値は既定で**マスク**されます(`KEY=`). デフォルト値に誤って秘密情報を渡してもダンプ先に漏れないようにするためです. デフォルト値も出力したい場合は `getenv.Logger.DumpValues(true)` を呼び出してください. 出力先は `getenv.Logger.SetWriter(w)` で変更できます.
