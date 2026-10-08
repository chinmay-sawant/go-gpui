package remote

import (
	"context"
	"testing"
)

func TestPhoneRemoteFits(t *testing.T) {
	app := newTest(t, WithPhone(true))
	app.Page().SetSize(1080, 2400)
	if err := app.Page().Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	var bottom float64
	for _, box := range app.Page().Boxes() {
		if box.Y+box.H > bottom {
			bottom = box.Y + box.H
		}
	}

	t.Logf("bottom %v", bottom)
	if bottom > 2400 {
		t.Fatalf("page is %v tall", bottom)
	}
}
