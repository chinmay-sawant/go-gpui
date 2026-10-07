package scene

import "context"

// onKeyDown routes a key press: the history screen and the scene's own
// shortcuts first, then the core input adapter. The handler never draws;
// the tick paints the result.
func (s *Scene) onKeyDown(ctx context.Context, key string) error {
	if s.hist.open {
		return s.historyKey(ctx, key)
	}

	switch key {
	case "h":
		return s.openHistory(ctx)
	case "t":
		return s.toggleTheme(ctx)
	}

	s.model.Down(key)

	return nil
}

// onKeyUp releases a key in the core adapter.
func (s *Scene) onKeyUp(_ context.Context, key string) error {
	if s.hist.open {
		return nil
	}

	s.model.Up(key)

	return nil
}

// historyKey routes a key press on the score screen.
func (s *Scene) historyKey(ctx context.Context, key string) error {
	switch key {
	case "h", "escape":
		return s.closeHistory(ctx)
	case "arrowleft", "bracketleft", "pageup":
		return s.pageHistory(ctx, -1)
	case "arrowright", "bracketright", "pagedown":
		return s.pageHistory(ctx, +1)
	case "t":
		return s.toggleTheme(ctx)
	}

	return nil
}
