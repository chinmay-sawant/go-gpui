package window

// devPanelPress runs the action of the hit under the pointer. A click on
// empty panel space does nothing, so a pinned box stays pinned.
func (s *shell) devPanelPress(x, y int) {
	for _, hit := range s.dev.hits {
		if devInRect(hit.rect, float64(x), float64(y)) {
			s.devRun(hit)

			return
		}
	}
}

// devRun applies one panel hit.
func (s *shell) devRun(hit devHit) {
	switch hit.act {
	case devActTab:
		s.dev.tab = devTabAt(hit.arg)
		s.dev.scroll = 0
	case devActOpsToggle:
		s.dev.ops = !s.dev.ops
	case devActJSON:
		if s.dev.collapsed == nil {
			s.dev.collapsed = map[string]bool{}
		}

		s.dev.collapsed[hit.key] = !s.dev.collapsed[hit.key]
	case devActOp:
		if s.dev.haveOp && s.dev.opPick == hit.arg {
			s.dev.haveOp = false

			return
		}

		s.dev.opPick, s.dev.haveOp = hit.arg, true
	}
}

// devResize follows the pointer with the dock's left edge.
func (s *shell) devResize(x int) {
	w := float64(s.screenW - x)
	if w < devDockMin {
		w = devDockMin
	}

	if w > float64(s.screenW) {
		w = float64(s.screenW)
	}

	s.dev.dockW = w
}
