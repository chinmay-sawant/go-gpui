package window

import (
	"strings"
	"testing"
)

func TestDevPanelShowsRelayoutCounts(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.commits, s.skipped = 3, 2

	joined := strings.Join(s.devLines(), "\n")
	if !strings.Contains(joined, "relayouts 3  skipped 2") {
		t.Fatalf("panel lines = %q", joined)
	}
}
