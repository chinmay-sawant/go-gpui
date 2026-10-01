// Package clipboard reads and writes the desktop clipboard.
// The last text is also kept in memory, for a build with no clipboard
// program and for a paste that comes back before the desktop answers.
package clipboard

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"sync"
	"time"
)

const wait = 200 * time.Millisecond

var (
	mu  sync.Mutex
	mem string
)

// Write stores text and tries the desktop clipboard.
func Write(text string) {
	mu.Lock()
	mem = text
	mu.Unlock()

	writeOS(text)
}

// Read returns the desktop clipboard, or the last Write when that fails.
func Read() string {
	if text, ok := readOS(); ok {
		return text
	}

	mu.Lock()
	defer mu.Unlock()

	return mem
}

func writeOS(text string) {
	if os.Getenv("WAYLAND_DISPLAY") != "" && look("wl-copy") {
		if _, err := run(text, "wl-copy", "--trim-newline"); err == nil {
			return
		}
	}

	if os.Getenv("DISPLAY") != "" && look("xclip") {
		_, _ = run(text, "xclip", "-selection", "clipboard")
	}
}

func readOS() (string, bool) {
	if os.Getenv("WAYLAND_DISPLAY") != "" && look("wl-paste") {
		if text, err := run("", "wl-paste", "-n"); err == nil {
			return text, true
		}
	}

	if os.Getenv("DISPLAY") != "" && look("xclip") {
		text, err := run("", "xclip", "-selection", "clipboard", "-o")
		if err == nil {
			return text, true
		}
	}

	return "", false
}

func look(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func run(stdin, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	if stdin != "" {
		cmd.Stdin = bytes.NewReader([]byte(stdin))
	}

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
