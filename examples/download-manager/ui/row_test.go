package ui

import "testing"

func TestStateTextReadsWithoutColor(t *testing.T) {
	cases := []struct {
		state State
		want  string
	}{
		{StateQueued, "… Queued"},
		{StateRunning, "▼ Running"},
		{StatePaused, "‖ Paused"},
		{StateCompleted, "✓ Completed"},
		{StateFailed, "✗ Failed"},
		{StateCancelled, "⊘ Cancelled"},
	}

	for _, c := range cases {
		row := Row{State: c.state}

		if got := row.StateText(); got != c.want {
			t.Errorf("StateText(%v) = %q, want %q", c.state, got, c.want)
		}
	}
}

func TestRowProgress(t *testing.T) {
	cases := []struct {
		done, total int64
		want        string
	}{
		{0, 100, "0%"},
		{55, 100, "55%"},
		{100, 100, "100%"},
		{150, 100, "100%"},
		{10, -1, "--"},
		{10, 0, "--"},
	}

	for _, c := range cases {
		row := Row{Done: c.done, Total: c.total}

		if got := row.Progress(); got != c.want {
			t.Errorf("Progress(%d/%d) = %q, want %q", c.done, c.total, got, c.want)
		}
	}
}

func TestRowButtons(t *testing.T) {
	running := Row{State: StateRunning}
	if !running.CanPause() || !running.CanCancel() || running.CanResume() || running.CanRetry() {
		t.Fatal("running buttons wrong")
	}

	paused := Row{State: StatePaused}
	if !paused.CanResume() || !paused.CanCancel() || paused.CanPause() {
		t.Fatal("paused buttons wrong")
	}

	failed := Row{State: StateFailed}
	if !failed.CanRetry() || !failed.CanRemove() || failed.CanCancel() {
		t.Fatal("failed buttons wrong")
	}

	completed := Row{State: StateCompleted}
	if !completed.CanRetry() || !completed.CanRemove() || completed.CanCancel() {
		t.Fatal("completed buttons wrong")
	}
}
