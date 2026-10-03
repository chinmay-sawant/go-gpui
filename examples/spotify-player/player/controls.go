package player

import (
	"context"
	"strconv"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick applies the action of the clicked box. Click redraws after this.
func (a *App) onClick(ctx context.Context, box gpui.Box) error {
	action := box.Action

	switch {
	case action == "play":
		a.view.Playing = !a.view.Playing
	case action == "next":
		a.step(1)
	case action == "prev":
		a.step(-1)
	case strings.HasPrefix(action, "row-"):
		a.selectRow(slot(action, "row-"))
	case strings.HasPrefix(action, "pick-"):
		a.selectPick(slot(action, "pick-"))
	case strings.HasPrefix(action, "card-"):
		a.selectShelf(slot(action, "card-"))
	case strings.HasPrefix(action, "list-"):
		a.selectList(slot(action, "list-"))
	case action == "nav-home":
		a.view.Nav = "home"
	case action == "nav-search":
		a.view.Nav = "search"
	case action == "heart":
		a.view.Liked = !a.view.Liked
	case action == "shuffle":
		a.view.Shuffle = !a.view.Shuffle
	case action == "repeat":
		a.view.Repeat = !a.view.Repeat
	case action == "mute":
		a.toggleMute()
	case strings.HasPrefix(action, "seek-"):
		a.setProgress(slot(action, "seek-")*10 + 10)
	case strings.HasPrefix(action, "volume-"):
		a.view.Volume = slot(action, "volume-")*10 + 10
	case action == "search":
		return a.searchNow(ctx)
	}

	a.page.SetData(a.view)

	return nil
}

// slot parses the trailing number of an action like "row-3".
func slot(action, prefix string) int {
	n, err := strconv.Atoi(strings.TrimPrefix(action, prefix))
	if err != nil {
		return -1
	}

	return n
}

// searchNow runs a live search when the box has a non-empty value.
func (a *App) searchNow(ctx context.Context) error {
	q := strings.TrimSpace(a.page.FormValue("q"))
	if q == "" {
		return nil
	}

	a.view.Query = q

	if err := a.Load(ctx, q); err != nil {
		a.page.SetData(a.view)
	}

	return nil
}
