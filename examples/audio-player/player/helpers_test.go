package player

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/music"
)

// newApp builds a player over the fake engine and draws one frame.
func newApp(t *testing.T) *App {
	t.Helper()

	app, _ := newAudioApp(t)

	return app
}

// newAudioApp builds a player whose engine opens one fake voice.
func newAudioApp(t *testing.T) (*App, *fakeVoice) {
	t.Helper()

	voice := &fakeVoice{dur: 3 * time.Minute}
	engine := music.NewEngine(fakeResolver{clip: testClip})
	engine.Open = func([]byte) (music.Voice, error) { return voice, nil }

	app, err := NewWith("", engine)
	if err != nil {
		t.Fatalf("NewWith: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	return app, voice
}

// waitLoaded ticks until the engine holds a voice, so a resolve is done.
func waitLoaded(t *testing.T, app *App) {
	t.Helper()

	for i := 0; i < 2000; i++ {
		if err := app.Tick(context.Background()); err != nil {
			t.Fatalf("Tick: %v", err)
		}

		if !app.audio.Loading() && app.audio.Credit().Title != "" {
			return
		}

		time.Sleep(time.Millisecond)
	}

	t.Fatal("audio did not load")
}

func clickBox(t *testing.T, app *App, id string) {
	t.Helper()

	box, ok := findBox(app.Boxes(), id)
	if !ok {
		t.Fatalf("no %s box", id)
	}

	if err := app.Click(context.Background(), box.X+box.W/2, box.Y+box.H/2); err != nil {
		t.Fatalf("click %s: %v", id, err)
	}
}

// assertReplayable pins the display-list path after a click redraw.
func assertReplayable(t *testing.T, app *App) {
	t.Helper()

	if app.Page().Display() == nil {
		t.Fatal("the frame fell back to the bitmap path")
	}

	if len(app.PNG()) == 0 {
		t.Fatal("no PNG after the redraw")
	}
}

func findBox(boxes []gpui.Box, id string) (gpui.Box, bool) {
	for _, box := range boxes {
		if box.ID == id {
			return box, true
		}
	}

	return gpui.Box{}, false
}
