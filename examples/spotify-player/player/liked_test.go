package player

import "testing"

func TestLikedScreenNavAndBoxes(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)

	click(t, app, "nav-liked")

	if v := app.View(); v.Nav != "liked" {
		t.Fatalf("Nav = %q", v.Nav)
	}

	redraw(t, app)

	for _, id := range []string{"liked-play", "liked-heart", "liked-row-0", "liked-row-7"} {
		if _, ok := findBox(app.Boxes(), id); !ok {
			t.Fatalf("no box %s", id)
		}
	}

	if _, ok := findBox(app.Boxes(), "row-0"); ok {
		t.Fatal("home tracklist drew under liked")
	}
}

func TestLikedScreenRowSelect(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)

	click(t, app, "nav-liked")
	redraw(t, app)

	want := app.View().LikedSongs.Tracks[2]

	click(t, app, "liked-row-2")

	v := app.View()
	if v.Now.Title != want.Title {
		t.Fatalf("Now = %q want %q", v.Now.Title, want.Title)
	}

	if v.Progress != 0 || v.Elapsed != "0:00" {
		t.Fatalf("progress = %d %q", v.Progress, v.Elapsed)
	}

	if !v.LikedSongs.Tracks[2].Active || v.LikedSongs.Tracks[0].Active {
		t.Fatal("active row did not move")
	}
}

func TestLikedScreenPlayAndHeart(t *testing.T) {
	app, voice := newFakeApp(t)
	redraw(t, app)

	click(t, app, "nav-liked")
	redraw(t, app)

	before := app.View().Liked

	click(t, app, "liked-heart")

	if app.View().Liked == before {
		t.Fatal("heart did not toggle")
	}

	click(t, app, "liked-play")
	waitAudio(t, app)

	if !app.View().Playing || !voice.playing {
		t.Fatalf("playing = %v voice = %v", app.View().Playing, voice.playing)
	}
}
