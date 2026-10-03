package player

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick applies one control action to the view.
func (a *App) onClick(ctx context.Context, box gpui.Box) error {
	action := box.Action

	switch {
	case action == "play":
		a.togglePlay(ctx)
	case action == "next":
		a.step(1)
		a.playNow(ctx)
	case action == "prev":
		a.step(-1)
		a.playNow(ctx)
	case action == "shuffle":
		a.view.Shuffle = !a.view.Shuffle
	case action == "repeat":
		a.view.Repeat = !a.view.Repeat
	case action == "heart":
		a.view.Liked = !a.view.Liked
	case strings.HasPrefix(action, "seek-"):
		if i, ok := actionIndex(action, "seek-"); ok {
			a.seekTo(i)
			a.seekAudio()
		}
	case strings.HasPrefix(action, "volume-"):
		if i, ok := actionIndex(action, "volume-"); ok {
			a.setVolume(i)
			a.volumeAudio()
		}
	case action == "search":
		return a.search(ctx)
	case strings.HasPrefix(action, "nav-"):
		a.view.Nav = strings.TrimPrefix(action, "nav-")
	case strings.HasPrefix(action, "queue-"):
		if i, ok := actionIndex(action, "queue-"); ok {
			a.selectQueue(i)
			a.playNow(ctx)
		}
	case strings.HasPrefix(action, "recent-"):
		if i, ok := actionIndex(action, "recent-"); ok && i < len(a.view.Recent) {
			a.selectQueue(a.view.Recent[i].Index)
			a.playNow(ctx)
		}
	case strings.HasPrefix(action, "list-"):
		a.selectPlaylist(action)
	default:
		return nil
	}

	a.page.SetData(a.view)

	return nil
}
