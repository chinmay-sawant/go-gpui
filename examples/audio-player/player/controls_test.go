package player

import "testing"

func TestClickSeekAndVolumeMap(t *testing.T) {
	app := newApp(t)

	clickBox(t, app, "seek-6")

	v := app.View()
	if v.Progress != 70 {
		t.Fatalf("Progress = %d, want 70", v.Progress)
	}

	clickBox(t, app, "volume-2")

	if v = app.View(); v.Volume != 30 {
		t.Fatalf("Volume = %d, want 30", v.Volume)
	}

	assertReplayable(t, app)
}

func TestClickBarHitsSegment(t *testing.T) {
	app := newApp(t)
	was := app.View().Progress

	clickBox(t, app, "seek")

	v := app.View()
	if v.Progress == was || v.Progress < 10 || v.Progress > 100 {
		t.Fatalf("Progress = %d, was %d", v.Progress, was)
	}

	assertReplayable(t, app)
}

func TestClickNavAndList(t *testing.T) {
	app := newApp(t)

	clickBox(t, app, "nav-library")
	clickBox(t, app, "list-3")

	v := app.View()
	if v.Nav != "library" {
		t.Fatalf("Nav = %q", v.Nav)
	}

	if !v.Playlists[3].Active || v.Playlists[0].Active {
		t.Fatal("the playlist highlight did not move")
	}

	assertReplayable(t, app)
}
