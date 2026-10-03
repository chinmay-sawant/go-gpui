package player

import "testing"

func TestClickPlayToggles(t *testing.T) {
	app := newApp(t)
	was := app.View().Playing

	clickBox(t, app, "play")

	if app.View().Playing == was {
		t.Fatalf("Playing stayed %v", was)
	}

	assertReplayable(t, app)
}

func TestClickNextChangesTrack(t *testing.T) {
	app := newApp(t)
	before := app.View().Now.Title

	clickBox(t, app, "next")

	v := app.View()
	if v.Now.Title == before {
		t.Fatalf("Now stayed %q", before)
	}

	if v.Progress != 0 || v.Elapsed != "0:00" || v.Now.Title != v.Queue[1].Title {
		t.Fatalf("next left Now=%q progress=%d elapsed=%q", v.Now.Title, v.Progress, v.Elapsed)
	}

	assertReplayable(t, app)
}

func TestClickQueueRowSelects(t *testing.T) {
	app := newApp(t)

	clickBox(t, app, "queue-2")

	v := app.View()
	if v.Now.Title != v.Queue[2].Title || !v.Queue[2].Active {
		t.Fatalf("Now = %q, active = %v", v.Now.Title, v.Queue[2].Active)
	}

	if v.Queue[0].Active {
		t.Fatal("the old row is still active")
	}

	assertReplayable(t, app)
}
