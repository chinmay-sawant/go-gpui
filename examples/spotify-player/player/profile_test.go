package player

import (
	"context"
	"testing"
)

func TestAvatarOpensProfileAndKeepsPlaying(t *testing.T) {
	app, voice := newFakeApp(t)
	redraw(t, app)

	click(t, app, "play")
	waitAudio(t, app)

	before := app.View()

	click(t, app, "avatar")

	v := app.View()
	if v.Nav != "profile" {
		t.Fatalf("Nav = %q", v.Nav)
	}

	if v.Playing != before.Playing || v.Now != before.Now || v.Progress != before.Progress {
		t.Fatalf("profile changed playback: %+v", v)
	}

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if !voice.playing {
		t.Fatal("profile stopped the audio")
	}

	redraw(t, app)

	if _, ok := findBox(app.Boxes(), "pcard-0"); !ok {
		t.Fatal("profile did not draw public playlists")
	}

	if _, ok := findBox(app.Boxes(), "row-0"); ok {
		t.Fatal("home tracklist drew under profile")
	}

	click(t, app, "nav-home")

	if v = app.View(); v.Nav != "home" {
		t.Fatalf("Nav = %q after home", v.Nav)
	}

	redraw(t, app)

	if _, ok := findBox(app.Boxes(), "row-0"); !ok {
		t.Fatal("home tracklist did not come back")
	}
}
