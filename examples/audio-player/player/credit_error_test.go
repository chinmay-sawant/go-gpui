package player

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chinmay-sawant/go-gpui/examples/music"
)

func TestCreditShowsAudioError(t *testing.T) {
	engine := music.NewEngine(fakeResolver{clip: testClip})
	engine.Open = func([]byte) (music.Voice, error) { return nil, errors.New("boom") }

	app, err := NewWith("", engine)
	if err != nil {
		t.Fatalf("NewWith: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	for i := 0; i < 2000 && app.audio.Err() == nil; i++ {
		if err := app.Tick(context.Background()); err != nil {
			t.Fatalf("Tick: %v", err)
		}

		time.Sleep(time.Millisecond)
	}

	if app.audio.Err() == nil {
		t.Fatal("the open error never surfaced")
	}

	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if got := app.View().Credit; got != "Audio unavailable" {
		t.Fatalf("Credit = %q (loading=%v err=%v)", got, app.audio.Loading(), app.audio.Err())
	}
}
