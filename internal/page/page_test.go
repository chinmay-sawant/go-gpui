package page_test

import (
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestNewRejectsEmptyHTML(t *testing.T) {
	t.Parallel()

	_, err := page.New(page.Config{Width: 320, Height: 400})
	if err != page.ErrEmptyHTML {
		t.Fatalf("err = %v", err)
	}
}

func TestSetSizeClamps(t *testing.T) {
	t.Parallel()

	screen, err := page.New(page.Config{
		HTML:      "<p>Hi</p>",
		Width:     480,
		Height:    640,
		MinWidth:  320,
		MinHeight: 400,
		MaxWidth:  800,
		MaxHeight: 900,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetSize(10, 10)
	width, height := screen.Size()
	if width != 320 || height != 400 {
		t.Fatalf("min clamp = %d x %d", width, height)
	}

	screen.SetSize(5000, 5000)
	width, height = screen.Size()
	if width != 800 || height != 900 {
		t.Fatalf("max clamp = %d x %d", width, height)
	}
}
