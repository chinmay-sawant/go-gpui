package player

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/go-gpui/examples/music"
)

func TestLoadStartsNewAudio(t *testing.T) {
	srv, _ := cannedServer(t)

	voice := &fakeVoice{dur: time.Minute}
	engine := music.NewEngine(fakeResolver{clip: testClip})
	engine.Open = func([]byte) (music.Voice, error) { return voice, nil }

	app, err := NewWith(srv.URL, engine)
	if err != nil {
		t.Fatalf("NewWith: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	waitLoaded(t, app)

	if err := app.Load(context.Background(), "canned"); err != nil {
		t.Fatalf("Load: %v", err)
	}

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if !app.audio.Loading() {
		t.Fatal("live Load did not request new audio")
	}

	waitLoaded(t, app)

	if app.View().Now.Title != "Canned One" {
		t.Fatalf("Now = %q", app.View().Now.Title)
	}
}
