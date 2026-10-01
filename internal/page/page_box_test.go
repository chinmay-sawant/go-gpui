package page_test

import (
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func boxCenter(t *testing.T, screen *page.Page, id string) (float64, float64) {
	t.Helper()

	for _, box := range screen.Boxes() {
		if box.ID != id || box.W <= 0 || box.H <= 0 {
			continue
		}

		return box.X + box.W/2, box.Y + box.H/2
	}

	t.Fatalf("no box id=%q", id)

	return 0, 0
}

func boxText(t *testing.T, screen *page.Page, id string) string {
	t.Helper()

	for _, box := range screen.Boxes() {
		if box.ID == id {
			return box.Text
		}
	}

	t.Fatalf("no box id=%q", id)

	return ""
}
