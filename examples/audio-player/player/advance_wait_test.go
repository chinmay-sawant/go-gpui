package player

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/chinmay-sawant/go-gpui/examples/music"
)

// gateResolver answers the first request at once and holds every later one
// until release is closed, so a test can observe a resolve in flight.
type gateResolver struct {
	mu      sync.Mutex
	calls   int
	release chan struct{}
	clip    music.Clip
}

func (g *gateResolver) Resolve(ctx context.Context, _ string, _ int) (music.Clip, error) {
	g.mu.Lock()
	g.calls++
	first := g.calls == 1
	g.mu.Unlock()

	if first {
		return g.clip, nil
	}

	select {
	case <-g.release:
		return g.clip, nil
	case <-ctx.Done():
		return music.Clip{}, ctx.Err()
	}
}

func TestAdvanceWaitsForTheResolve(t *testing.T) {
	voice := &fakeVoice{dur: time.Minute}
	resolver := &gateResolver{release: make(chan struct{}), clip: testClip}
	engine := music.NewEngine(resolver)
	engine.Open = func([]byte) (music.Voice, error) { return voice, nil }

	app, err := NewWith("", engine)
	if err != nil {
		t.Fatalf("NewWith: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	waitLoaded(t, app)

	voice.pos = voice.dur
	voice.playing = false

	// The first tick advances to the next track and starts a blocked resolve.
	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	now := app.View().Now.Title

	// While that resolve is in flight the queue must not step again.
	if err := app.Tick(context.Background()); err != nil {
		t.Fatalf("Tick: %v", err)
	}

	if app.View().Now.Title != now {
		t.Fatalf("the queue advanced during a resolve: %q", app.View().Now.Title)
	}

	close(resolver.release)
	waitLoaded(t, app)

	if !voice.playing {
		t.Fatal("the next track did not start after the resolve")
	}
}
