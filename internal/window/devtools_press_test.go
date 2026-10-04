package window

import "testing"

// TestDevPanelTabsSwitch checks a click on a tab hit changes the active tab
// and clears the content scroll.
func TestDevPanelTabsSwitch(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.dev.scroll = 40

	s.devRun(devHit{act: devActTab, arg: int(devTabFrame)})
	if s.dev.tab != devTabFrame {
		t.Fatalf("tab = %v, want Frame", s.dev.tab)
	}

	if s.dev.scroll != 0 {
		t.Fatalf("scroll = %v, want 0 after a tab switch", s.dev.scroll)
	}
}

// TestDevPanelRunActions checks the toggle, JSON, and op rows.
func TestDevPanelRunActions(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.display = devTestDisplay()

	s.devRun(devHit{act: devActOpsToggle})
	if !s.dev.ops {
		t.Fatal("the toggle row did not turn the outlines on")
	}

	s.devRun(devHit{act: devActJSON, key: "rect"})
	if !s.dev.collapsed["rect"] {
		t.Fatal("the JSON row did not collapse rect")
	}

	s.devRun(devHit{act: devActOp, arg: 1})
	if !s.dev.haveOp || s.dev.opPick != 1 {
		t.Fatalf("op pick = %v/%d, want op 1", s.dev.haveOp, s.dev.opPick)
	}

	s.devRun(devHit{act: devActOp, arg: 1})
	if s.dev.haveOp {
		t.Fatal("the same op row did not clear the pick")
	}
}
