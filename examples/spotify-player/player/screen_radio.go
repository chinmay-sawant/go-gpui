package player

import (
	"context"
	"strings"
)

// RadioData is the data behind the Radio screen.
type RadioData struct {
	Featured Card
	Stations []Card
}

// defaultRadio returns the offline radio screen.
func defaultRadio() RadioData {
	return RadioData{
		Featured: Card{Title: "Nova Waves Radio", Sub: "Nova Waves, Aster Field and more", Cover: "card-0"},
		Stations: []Card{
			{Index: 0, Title: "Nova Waves Radio", Sub: "Nova Waves, Aster Field and more", Cover: "pick-0"},
			{Index: 1, Title: "Indie Coast Radio", Sub: "Indie Coast, The Far Coast and more", Cover: "pick-1"},
			{Index: 2, Title: "Lo-Fi Nights Radio", Sub: "Lo-Fi Nights, Mono Arcade and more", Cover: "pick-2"},
			{Index: 3, Title: "Synthwave FM", Sub: "Synthwave FM, Neon Horizon and more", Cover: "pick-3"},
			{Index: 4, Title: "Acoustic Mornings Radio", Sub: "Acoustic Mornings, Paper Trails and more", Cover: "pick-4"},
			{Index: 5, Title: "Deep Focus Radio", Sub: "Deep Focus, Blue Hour and more", Cover: "pick-5"},
		},
	}
}

// clickRadio applies one action on the radio screen.
func (a *App) clickRadio(ctx context.Context, action string) error {
	switch {
	case action == "radio-play":
		a.RadioTogglePlay(ctx)
	case strings.HasPrefix(action, "radio-st-"):
		a.RadioSelectStation(ctx, slot(action, "radio-st-"))
	}

	return nil
}

// RadioTogglePlay mirrors the now-bar play button for the radio screen.
func (a *App) RadioTogglePlay(ctx context.Context) {
	a.view.Playing = !a.view.Playing

	if a.view.Playing {
		if a.audio.Duration() > 0 {
			a.audio.Play()
		} else {
			a.playNow(ctx)
		}

		return
	}

	a.audio.Pause()
}

// RadioSelectStation makes station i the now-playing item and starts it.
func (a *App) RadioSelectStation(ctx context.Context, i int) {
	if i < 0 || i >= len(a.view.Radio.Stations) {
		return
	}

	for j := range a.view.Radio.Stations {
		a.view.Radio.Stations[j].Active = j == i
	}

	a.setNow(a.view.Radio.Stations[i])
	a.playNow(ctx)
}
