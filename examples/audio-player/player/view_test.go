package player

import (
	"image"
	"image/color"
	"testing"
)

func TestDefaultView(t *testing.T) {
	v := DefaultView()

	if len(v.Queue) != 6 {
		t.Fatalf("Queue = %d tracks", len(v.Queue))
	}

	if v.Now.Title != v.Queue[0].Title || !v.Queue[0].Active {
		t.Fatalf("Now = %q, Queue[0] = %q", v.Now.Title, v.Queue[0].Title)
	}

	if v.Progress < 0 || v.Progress > 100 {
		t.Fatalf("Progress = %d", v.Progress)
	}

	if len(v.Recent) != 4 || len(v.Playlists) != 6 {
		t.Fatalf("Recent = %d, Playlists = %d", len(v.Recent), len(v.Playlists))
	}

	if v.Status != offlineStatus || !v.Playing {
		t.Fatalf("Status = %q, Playing = %v", v.Status, v.Playing)
	}

	for _, track := range v.Queue {
		if track.Cover == "" || track.Length == "" {
			t.Fatalf("track %d is incomplete: %+v", track.Index, track)
		}
	}
}

// hasColor reports whether img holds a pixel within tol of want.
func hasColor(img image.Image, want color.RGBA, tol int) bool {
	for y := 0; y < img.Bounds().Dy(); y += 2 {
		for x := 0; x < img.Bounds().Dx(); x += 2 {
			r, g, b, _ := img.At(x, y).RGBA()

			if near(int(r>>8), int(want.R), tol) &&
				near(int(g>>8), int(want.G), tol) &&
				near(int(b>>8), int(want.B), tol) {
				return true
			}
		}
	}

	return false
}

func near(a, b, tol int) bool {
	if a < b {
		a, b = b, a
	}

	return a-b <= tol
}
