package music

import (
	"hash/fnv"
	"math"
)

// Demo tune constants: a short generated WAV so playback works offline.
const (
	demoRate    = 48000
	demoSeconds = 16
	demoStep    = 0.32 // seconds per arpeggio step
)

// DemoTune returns a Clip holding a generated tune. The query picks the key,
// so different tracks sound different.
func DemoTune(query string) Clip {
	return Clip{
		Title:   "Demo tone",
		Creator: "go-gpui",
		License: "generated",
		Data:    wavTune(seed(query)),
	}
}

// seed folds a query into a small deterministic key offset.
func seed(query string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(query))

	return int(h.Sum32() % 12)
}

// wavTune renders an arpeggio over a four-chord loop to 16-bit stereo PCM.
func wavTune(root int) []byte {
	frames := demoRate * demoSeconds
	pcm := make([]int16, frames*2)
	step := int(demoStep * demoRate)
	chords := [4][3]int{{0, 4, 7}, {-4, 0, 3}, {-1, 2, 7}, {-3, 2, 5}}

	for i := 0; i < frames; i++ {
		chord := chords[i/step%4]
		note := chord[i/step%3]
		t := float64(i%demoRate) / demoRate
		env := math.Exp(-3 * t)
		freq := 220 * math.Pow(2, float64(root+note)/12)
		v := math.Sin(2*math.Pi*freq*float64(i)/demoRate) * env * 0.35
		s := int16(v * math.MaxInt16)
		pcm[i*2] = s
		pcm[i*2+1] = s
	}

	return wavBytes(pcm)
}
