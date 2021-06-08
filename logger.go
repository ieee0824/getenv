package getenv

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

var Logger = &EnvLogger{
	w: os.Stdout,
}

type EnvLogger struct {
	w io.Writer
}

func (l *EnvLogger) SetWriter(w io.Writer) {
	l.w = w
}

func (l *EnvLogger) Dump(key string, def ...interface{}) {
	switch os.Getenv("GETENV_DUMP_MODE") {
	case "dotenv":
		if err := l.dumpDotEnv(key, def...); err != nil {
			panic(err)
		}
	}
}

func (l *EnvLogger) dumpDotEnv(key string, def ...interface{}) error {
	if _, err := l.w.Write([]byte(key)); err != nil {
		return err
	}
	if _, err := l.w.Write([]byte("=")); err != nil {
		return err
	}
	if len(def) == 0 {
		if _, err := l.w.Write([]byte("\n")); err != nil {
			return err
		}
		return nil
	}

	var v interface{}
	switch reflect.TypeOf(def[0]).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(def[0])
		if s.Len() == 0 {
			if _, err := l.w.Write([]byte("\n")); err != nil {
				return err
			}
			return nil
		}
		v = s.Index(0).Interface()
	default:
		v = def[0]
	}
	if _, err := l.w.Write([]byte(fmt.Sprintf("%v\n", v))); err != nil {
		return err
	}

	return nil
}
