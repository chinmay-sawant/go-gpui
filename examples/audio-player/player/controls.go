package player

import (
	"context"
	"strconv"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick applies one control action to the view.
func (a *App) onClick(ctx context.Context, box gpui.Box) error {
	action := box.Action

	switch {
	case action == "play":
		a.view.Playing = !a.view.Playing
	case action == "next":
		a.step(1)
	case action == "prev":
		a.step(-1)
	case action == "shuffle":
		a.view.Shuffle = !a.view.Shuffle
	case action == "repeat":
		a.view.Repeat = !a.view.Repeat
	case action == "heart":
		a.view.Liked = !a.view.Liked
	case strings.HasPrefix(action, "seek-"):
		if i, ok := actionIndex(action, "seek-"); ok {
			a.seekTo(i)
		}
	case strings.HasPrefix(action, "volume-"):
		if i, ok := actionIndex(action, "volume-"); ok {
			a.setVolume(i)
		}
	case action == "search":
		return a.search(ctx)
	case strings.HasPrefix(action, "nav-"):
		a.view.Nav = strings.TrimPrefix(action, "nav-")
	case strings.HasPrefix(action, "queue-"):
		if i, ok := actionIndex(action, "queue-"); ok {
			a.selectQueue(i)
		}
	case strings.HasPrefix(action, "recent-"):
		if i, ok := actionIndex(action, "recent-"); ok && i < len(a.view.Recent) {
			a.selectQueue(a.view.Recent[i].Index)
		}
	case strings.HasPrefix(action, "list-"):
		a.selectPlaylist(action)
	default:
		return nil
	}

	a.page.SetData(a.view)

	return nil
}

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
