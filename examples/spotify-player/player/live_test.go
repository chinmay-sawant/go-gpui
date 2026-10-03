package player

import (
	"context"
	"testing"
)

func TestLiveApplyStartsAudio(t *testing.T) {
	app, voice := newFakeApp(t)
	redraw(t, app)

	app.apply("test", sampleTracks(), nil, nil)

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	waitAudio(t, app)

	if !voice.playing {
		t.Fatal("live data did not start the voice")
	}

	if !app.View().Playing {
		t.Fatal("live data did not set Playing")
	}
}

func TestEndedTrackAdvances(t *testing.T) {
	app, voice := newFakeApp(t)
	redraw(t, app)

	click(t, app, "play")
	waitAudio(t, app)

	voice.pos = voice.dur
	voice.playing = false

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	// The fake Open reuses one voice, so the next track starts at zero.
	voice.pos = 0
	waitAudio(t, app)

	v := app.View()
	if !v.Tracks[1].Active || v.Now.Title != v.Tracks[1].Title {
		t.Fatalf("now = %+v", v.Now)
	}
}
