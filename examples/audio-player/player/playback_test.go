package player

import (
	"testing"
	"time"
)

func TestPlayStartsFakeVoice(t *testing.T) {
	app, voice := newAudioApp(t)

	clickBox(t, app, "play") // pause
	clickBox(t, app, "play") // play now
	waitLoaded(t, app)

	if !voice.playing || !app.audio.Playing() {
		t.Fatalf("voice playing = %v, engine = %v", voice.playing, app.audio.Playing())
	}
}

func TestSeekClickSeeksVoice(t *testing.T) {
	app, voice := newAudioApp(t)
	waitLoaded(t, app)

	clickBox(t, app, "seek-6")

	want := time.Duration(float64(voice.dur) * 0.7)
	if got := app.audio.Position(); got != want {
		t.Fatalf("Position = %v, want %v", got, want)
	}
}

func TestVolumeClickSetsVoice(t *testing.T) {
	app, voice := newAudioApp(t)
	waitLoaded(t, app)

	clickBox(t, app, "volume-2")

	if voice.volume != 0.3 {
		t.Fatalf("volume = %v, want 0.3", voice.volume)
	}
}
