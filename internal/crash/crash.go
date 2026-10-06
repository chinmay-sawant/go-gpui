// Package crash writes a panic report to a text file on this machine.
// The file is not uploaded.
package crash

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

var (
	mu     sync.Mutex
	custom string
)

// SetDir sets the folder Write uses. An empty dir uses the default.
func SetDir(dir string) {
	mu.Lock()
	custom = dir
	mu.Unlock()
}

// Write creates the folder and writes one UTF-8 report.
// The name starts with a UTC timestamp, and a later call in that
// same second gets a different name.
func Write(title, reason string) (string, error) {
	folder := location()
	if err := os.MkdirAll(folder, 0o755); err != nil {
		return "", err
	}

	body := fmt.Sprintf(
		"title: %s\nreason: %s\ngo: %s\n%s",
		title,
		reason,
		runtime.Version(),
		debug.Stack(),
	)

	return create(folder, body)
}

func location() string {
	mu.Lock()
	dir := custom
	mu.Unlock()
	if dir != "" {
		return dir
	}

	config, err := os.UserConfigDir()
	if err != nil {
		return filepath.Join(os.TempDir(), "ownframe-crashes")
	}

	return filepath.Join(config, "ownframe", "crashes")
}

func create(folder, body string) (string, error) {
	stamp := time.Now().UTC().Format("20060102-150405")
	for n := 0; n < 1000; n++ {
		name := stamp + ".txt"
		if n > 0 {
			name = fmt.Sprintf("%s-%d.txt", stamp, n)
		}

		path := filepath.Join(folder, name)
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if errors.Is(err, os.ErrExist) {
			continue
		}

		if err != nil {
			return "", err
		}

		_, werr := f.WriteString(body)
		cerr := f.Close()
		if werr != nil {
			return "", werr
		}

		return path, cerr
	}

	return "", fmt.Errorf("crash: no free name")
}
