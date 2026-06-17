package getenv

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sync"
)

var Logger = &EnvLogger{
	w: os.Stdout,
}

type EnvLogger struct {
	mu         sync.Mutex
	w          io.Writer
	dumpValues bool
}

func (l *EnvLogger) SetWriter(w io.Writer) {
	l.mu.Lock()
	l.w = w
	l.mu.Unlock()
}

// DumpValues controls whether the dotenv dump output includes the
// (developer-provided) default values.
//
// It is disabled by default so that secrets accidentally placed in default values
// are not written to the dump target. Enable it explicitly only when the defaults
// are known to be non-sensitive.
func (l *EnvLogger) DumpValues(enabled bool) {
	l.mu.Lock()
	l.dumpValues = enabled
	l.mu.Unlock()
}

func (l *EnvLogger) Dump(key string, def ...interface{}) {
	switch os.Getenv("GETENV_DUMP_MODE") {
	case "dotenv":
		if err := l.dumpDotEnv(key, def...); err != nil {
			// A logging side effect must never crash the host program, so the
			// write error is reported best-effort instead of panicking.
			log.Printf("getenv: dump error: %v", err)
		}
	}
}

func (l *EnvLogger) dumpDotEnv(key string, def ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Build the whole record first and emit it with a single Write so concurrent
	// dumps cannot interleave into corrupted lines.
	var b bytes.Buffer
	b.WriteString(key)
	b.WriteByte('=')
	if l.dumpValues {
		b.WriteString(formatDefault(def...))
	}
	b.WriteByte('\n')

	_, err := l.w.Write(b.Bytes())
	return err
}

// formatDefault renders the first default value for dotenv output. Slice defaults
// render their first element, mirroring the historical behaviour. It returns an
// empty string when there is no usable default.
func formatDefault(def ...interface{}) string {
	if len(def) == 0 || def[0] == nil {
		return ""
	}
	v := def[0]
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		s := reflect.ValueOf(v)
		if s.Len() == 0 {
			return ""
		}
		v = s.Index(0).Interface()
	}
	return fmt.Sprintf("%v", v)
}
