package player

import "testing"

func TestBoxesCoverTheControls(t *testing.T) {
	app := newApp(t)

	for _, id := range []string{
		"play", "seek", "volume", "queue-0", "queue-5",
		"recent-0", "nav-home", "list-0", "seek-0", "seek-9", "volume-9",
	} {
		if _, ok := findBox(app.Boxes(), id); !ok {
			t.Fatalf("no %s box", id)
		}
	}
}
