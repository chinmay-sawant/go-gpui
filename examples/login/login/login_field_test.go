package login_test

import (
	"context"
	"testing"
)

func TestFieldHitBoxStaysInsideTheCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)
	card := boxByID(t, app, "card")
	email := boxByID(t, app, "email")

	if email.X < card.X || email.Y < card.Y {
		t.Fatalf("email origin = %.1f, %.1f, card origin = %.1f, %.1f", email.X, email.Y, card.X, card.Y)
	}

	if email.X+email.W > card.X+card.W+1 || email.Y+email.H > card.Y+card.H+1 {
		t.Fatalf("email extends outside the card")
	}

	click(t, ctx, app, "email", "")
	if got := app.Page().FocusedField(); got != "email" {
		t.Fatalf("focus = %q", got)
	}
}
