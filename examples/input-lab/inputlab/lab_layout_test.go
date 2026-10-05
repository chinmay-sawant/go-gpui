package inputlab

import (
	"testing"
)

func TestFieldWidth(t *testing.T) {
	app, _ := newTest(t)
	if w := boxByID(t, app, "b-name").W; w < 280 {
		t.Fatalf("b-name w=%v", w)
	}
}

func TestLoginButtonBelowPassword(t *testing.T) {
	app, _ := newTest(t)
	pw := boxByID(t, app, "si-password")
	btn := boxByID(t, app, "si-login")
	if btn.Y < pw.Y+pw.H {
		t.Fatalf("login y=%v overlaps password bottom %v", btn.Y, pw.Y+pw.H)
	}
	if btn.X != pw.X {
		t.Fatalf("login x=%v password x=%v", btn.X, pw.X)
	}
}

func TestMoreLinkVisible(t *testing.T) {
	app, _ := newTest(t)
	more := boxByID(t, app, "i-more")
	if more.H <= 0 || more.W <= 0 {
		t.Fatalf("more box=%+v", more)
	}
}
