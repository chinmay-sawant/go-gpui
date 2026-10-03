package player

import (
	"context"
	"strconv"
	"strings"
)

// search loads the query box. An empty box or a failed load leaves the view.
func (a *App) search(ctx context.Context) error {
	term := strings.TrimSpace(a.page.FormValue("q"))
	if term == "" {
		return nil
	}

	_ = a.Load(ctx, term)

	return nil
}

// actionIndex reads the numeric suffix of an action name.
func actionIndex(action, prefix string) (int, bool) {
	i, err := strconv.Atoi(strings.TrimPrefix(action, prefix))
	if err != nil || i < 0 {
		return 0, false
	}

	return i, true
}
