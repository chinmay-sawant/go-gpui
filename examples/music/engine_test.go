package music

import (
	"context"
	"testing"
	"time"
)

// newFakeEngine returns an engine whose Open hands back one fake voice.
func newFakeEngine() (*Engine, *fakeVoice) {
	voice := &fakeVoice{dur: time.Minute}
	engine := NewEngine(&fakeResolver{clip: Clip{Title: "Free", Data: []byte("x")}})
	engine.Open = func([]byte) (Voice, error) { return voice, nil }

	return engine, voice
}

// waitVoice ticks the engine until its background resolve lands.
func waitVoice(t *testing.T, engine *Engine) *fakeVoice {
	t.Helper()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		engine.Update()
		if !engine.Loading() {
			voice, ok := engine.voice.(*fakeVoice)
			if !ok {
				t.Fatalf("no voice: %v", engine.Err())
			}

			return voice
		}

		time.Sleep(time.Millisecond)
	}

	t.Fatal("the engine never finished loading")

	return nil
}

func TestEngineStartsOnUpdate(t *testing.T) {
	t.Parallel()

	engine, voice := newFakeEngine()
	engine.SetVolume(0.5)
	engine.Want(context.Background(), "jazz", 2)

	if got := waitVoice(t, engine); got != voice {
		t.Fatal("wrong voice installed")
	}

	if !voice.playing {
		t.Fatal("the voice did not start")
	}

	if voice.volume != 0.5 {
		t.Fatalf("volume = %v", voice.volume)
	}

	if engine.Credit().Title != "Free" {
		t.Fatalf("credit = %+v", engine.Credit())
	}
}
