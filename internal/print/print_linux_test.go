//go:build linux && !android

package print

import (
	"context"
	"testing"
)

func TestPrintUsesLinuxPath(t *testing.T) {
	calls := fakeRunner(t, map[string]bool{"lp": true})

	if err := Print(context.Background(), "/tmp/r.pdf"); err != nil {
		t.Fatal(err)
	}

	if got := *calls; len(got) != 1 || got[0].name != "/bin/lp" {
		t.Fatalf("calls = %+v", got)
	}
}

func TestAvailableOnLinux(t *testing.T) {
	if !Available() {
		t.Fatal("Available() = false on Linux")
	}
}
