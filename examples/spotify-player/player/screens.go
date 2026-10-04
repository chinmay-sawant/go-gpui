package player

import (
	"context"
	"strings"
)

// screenClick hands an action to the screen that owns its prefix. It reports
// false when no screen claimed the action.
func (a *App) screenClick(ctx context.Context, action string) (bool, error) {
	switch {
	case strings.HasPrefix(action, "search-"):
		return true, a.clickSearch(ctx, action)
	case strings.HasPrefix(action, "library-"):
		return true, a.clickLibrary(ctx, action)
	case strings.HasPrefix(action, "liked-"):
		return true, a.clickLiked(ctx, action)
	case strings.HasPrefix(action, "browse-"):
		return true, a.clickBrowse(ctx, action)
	case strings.HasPrefix(action, "radio-"):
		return true, a.clickRadio(ctx, action)
	case strings.HasPrefix(action, "queue-"):
		return true, a.clickQueue(ctx, action)
	}

	return false, nil
}
