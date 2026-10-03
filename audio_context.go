package gpui

import "github.com/hajimehoshi/ebiten/v2/audio"

// audioSampleRate is the sample rate of the context Run creates.
const audioSampleRate = 48000

// ensureAudio creates the Ebiten audio context once. Ebiten v2.10 does not
// create one itself, and the audio package needs a context before it can
// make a player. Run and BindMobile call this, so a page that plays sound
// through the host finds a context. Serve has none.
func ensureAudio() {
	if audio.CurrentContext() == nil {
		audio.NewContext(audioSampleRate)
	}
}
