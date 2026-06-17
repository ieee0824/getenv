package getenv

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) { return 0, errors.New("write failed") }

// useDumpLogger points the global Logger at the given writer with dotenv dump mode
// enabled, and restores the defaults afterwards.
func useDumpLogger(t *testing.T, w io.Writer, dumpValues bool) {
	t.Helper()
	t.Setenv("GETENV_DUMP_MODE", "dotenv")
	Logger.SetWriter(w)
	Logger.DumpValues(dumpValues)
	t.Cleanup(func() {
		Logger.SetWriter(os.Stdout)
		Logger.DumpValues(false)
	})
}

func TestDumpMasksValuesByDefault(t *testing.T) {
	var buf bytes.Buffer
	useDumpLogger(t, &buf, false)

	String("API_KEY", "sk-live-secret")

	assert.Equal(t, "API_KEY=\n", buf.String())
	assert.NotContains(t, buf.String(), "sk-live-secret")
}

func TestDumpShowsValuesWhenEnabled(t *testing.T) {
	var buf bytes.Buffer
	useDumpLogger(t, &buf, true)

	String("STR_ENV", "hoge")

	assert.Equal(t, "STR_ENV=hoge\n", buf.String())
}

func TestDumpSliceMaskedByDefault(t *testing.T) {
	var buf bytes.Buffer
	useDumpLogger(t, &buf, false)

	StringSlice("HOSTS", []string{"secret1", "secret2", "secret3"})

	assert.Equal(t, "HOSTS=\n", buf.String())
	assert.NotContains(t, buf.String(), "secret")
}

func TestDumpSliceShownWhenEnabled(t *testing.T) {
	var buf bytes.Buffer
	useDumpLogger(t, &buf, true)

	StringSlice("HOSTS", []string{"a", "b", "c"})

	assert.Equal(t, "HOSTS=[a b c]\n", buf.String())
}

func TestDumpDoesNotPanicOnWriteError(t *testing.T) {
	useDumpLogger(t, errWriter{}, true)

	assert.NotPanics(t, func() {
		got := String("SOME_KEY", "default")
		assert.Equal(t, "default", got)
	})
}

// TestDumpConcurrent exercises concurrent dumps while the writer is swapped, to be
// run under `go test -race`. It also asserts that records are never interleaved.
func TestDumpConcurrent(t *testing.T) {
	var mu sync.Mutex
	lines := map[string]struct{}{}
	collector := writerFunc(func(p []byte) (int, error) {
		mu.Lock()
		lines[string(p)] = struct{}{}
		mu.Unlock()
		return len(p), nil
	})

	useDumpLogger(t, collector, true)

	var wg sync.WaitGroup
	const n = 200
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			String("KEY", "val")
		}()
	}
	// Concurrently swap the writer to surface any data race on Logger.w.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Logger.SetWriter(collector)
		}()
	}
	wg.Wait()

	// Every record arrives as a single, well-formed write; no interleaving.
	mu.Lock()
	defer mu.Unlock()
	for line := range lines {
		assert.Equal(t, "KEY=val\n", line)
		assert.Equal(t, 1, strings.Count(line, "\n"))
	}
}

type writerFunc func(p []byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) { return f(p) }
