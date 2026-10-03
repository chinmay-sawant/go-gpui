package player

import (
	"context"
	"testing"
)

func TestEndedAdvancesQueue(t *testing.T) {
	app, voice := newAudioApp(t)
	waitLoaded(t, app)

	before := app.View().Now.Title
	voice.pos = voice.dur
	voice.playing = false

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if app.View().Now.Title == before {
		t.Fatalf("Now stayed %q", before)
	}

	waitLoaded(t, app)

	if !voice.playing {
		t.Fatal("the next track is not playing")
	}
}

func TestRepeatRestartsVoice(t *testing.T) {
	app, voice := newAudioApp(t)
	waitLoaded(t, app)

	clickBox(t, app, "repeat")
	voice.pos = voice.dur
	voice.playing = false

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if voice.pos != 0 || !voice.playing {
		t.Fatalf("repeat left pos = %v, playing = %v", voice.pos, voice.playing)
	}
}

func TestNextRequestsNewAudio(t *testing.T) {
	app, _ := newAudioApp(t)
	waitLoaded(t, app)

	clickBox(t, app, "next")

	if !app.audio.Loading() {
		t.Fatal("next did not request audio")
	}
}

func TestShuffleAdvances(t *testing.T) {
	app, voice := newAudioApp(t)
	waitLoaded(t, app)

	clickBox(t, app, "shuffle")
	voice.pos = voice.dur
	voice.playing = false

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if !app.audio.Loading() {
		t.Fatal("shuffle did not request audio")
	}

	if !app.View().Playing {
		t.Fatal("shuffle did not keep playing intent")
	}
}
