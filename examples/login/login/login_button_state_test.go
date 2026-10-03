package login_test

import (
	"context"
	"image/color"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/login/login"
)

var (
	blueButton     = color.RGBA{R: 0x1a, G: 0x56, B: 0xdb, A: 0xff}
	disabledButton = color.RGBA{R: 0xe3, G: 0xde, B: 0xd4, A: 0xff}
)

// buttonFill samples the flat fill inside the button edge, away from the label.
func buttonFill(t *testing.T, app *login.App) color.RGBA {
	t.Helper()

	btn := boxByID(t, app, "login")
	img := decodePNG(t, app)

	return rgbaAt(img, int(btn.X)+4, int(btn.Y+btn.H/2))
}

func TestButtonGrayUntilBothFields(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if got := buttonFill(t, app); got != disabledButton {
		t.Fatalf("empty fill = %+v, want %+v", got, disabledButton)
	}

	click(t, ctx, app, "email", "")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	if app.View().Ready {
		t.Fatal("ready after email alone")
	}

	if got := buttonFill(t, app); got != disabledButton {
		t.Fatalf("email-only fill = %+v, want %+v", got, disabledButton)
	}

	click(t, ctx, app, "password", "")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	if !app.View().Ready {
		t.Fatal("not ready after both fields")
	}

	if got := buttonFill(t, app); got != blueButton {
		t.Fatalf("filled fill = %+v, want %+v", got, blueButton)
	}
}
