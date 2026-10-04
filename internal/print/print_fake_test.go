package print

import (
	"context"
	"os/exec"
	"testing"
)

type call struct {
	name string
	args []string
}

// fakeRunner records commands and reports the named tools as installed.
func fakeRunner(t *testing.T, present map[string]bool) *[]call {
	t.Helper()

	calls := &[]call{}

	look, run := lookPath, runCommand
	t.Cleanup(func() {
		lookPath, runCommand = look, run
	})

	lookPath = func(name string) (string, error) {
		if present[name] {
			return "/bin/" + name, nil
		}

		return "", exec.ErrNotFound
	}

	runCommand = func(_ context.Context, path string, args ...string) error {
		*calls = append(*calls, call{name: path, args: args})

		return nil
	}

	return calls
}
