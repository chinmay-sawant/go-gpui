package player

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick applies the action of the clicked box. Click redraws after this.
func (a *App) onClick(ctx context.Context, box gpui.Box) error {
	action := box.Action

	switch {
	case action == "play":
		a.view.Playing = !a.view.Playing

		if a.view.Playing {
			if a.audio.Duration() > 0 {
				a.audio.Play()
			} else {
				a.playNow(ctx)
			}
		} else {
			a.audio.Pause()
		}
	case action == "next":
		a.step(1)
		a.playNow(ctx)
	case action == "prev":
		a.step(-1)
		a.playNow(ctx)
	case strings.HasPrefix(action, "row-"):
		a.selectRow(slot(action, "row-"))
		a.playNow(ctx)
	case strings.HasPrefix(action, "pick-"):
		a.selectPick(slot(action, "pick-"))
		a.playNow(ctx)
	case strings.HasPrefix(action, "card-"):
		a.selectShelf(slot(action, "card-"))
		a.playNow(ctx)
	case strings.HasPrefix(action, "list-"):
		a.selectList(slot(action, "list-"))
	case action == "nav-home":
		a.view.Nav = "home"
	case action == "nav-search":
		a.view.Nav = "search"
	case action == "nav-profile":
		a.view.Nav = "profile"
	case action == "heart":
		a.view.Liked = !a.view.Liked
	case action == "shuffle":
		a.view.Shuffle = !a.view.Shuffle
	case action == "repeat":
		a.view.Repeat = !a.view.Repeat
	case action == "mute":
		a.toggleMute()
		a.audio.SetVolume(float64(a.view.Volume) / 100)
	case strings.HasPrefix(action, "seek-"):
		a.setProgress(slot(action, "seek-")*10 + 10)
		a.audio.SeekFraction(float64(a.view.Progress) / 100)
	case strings.HasPrefix(action, "volume-"):
		a.view.Volume = slot(action, "volume-")*10 + 10
		a.audio.SetVolume(float64(a.view.Volume) / 100)
	case action == "search":
		return a.searchNow(ctx)
	}

	a.page.SetData(a.view)

	return nil
}
