package remote

import (
	"context"
	"testing"
)

func TestPhonePadAndKeys(t *testing.T) {
	app := newTest(t, WithPhone(true))
	fake := &fakeLink{}
	app.SetLink(fake)

	clickID(t, app, "tab-pad")
	if app.view.Panel != "pad" {
		t.Fatal(app.view.Panel)
	}

	clickID(t, app, "pad")
	if err := app.onTick(context.Background()); err != nil {
		t.Fatal(err)
	}
	if fake.spec != "click:" {
		t.Fatalf("click %s", fake.spec)
	}

	if err := app.onSwipe(context.Background(), 20, -8); err != nil {
		t.Fatal(err)
	}
	if fake.spec != "move:20,-8" {
		t.Fatalf("move %s", fake.spec)
	}

	clickID(t, app, "tab-remote")
	clickID(t, app, "up")
	if fake.spec != "button:UP" {
		t.Fatalf("up %s", fake.spec)
	}
}

func TestPowerWakesWhenOff(t *testing.T) {
	app := newTest(t)
	app.view.PowerOn = false
	fake := &fakeLink{}
	app.SetLink(fake)
	clickID(t, app, "power")
	if fake.wake != 1 {
		t.Fatalf("wake %d", fake.wake)
	}
	if app.Status() != "Waking the TV" {
		t.Fatalf("status %s", app.Status())
	}
}

func TestInputAndHDMISpecs(t *testing.T) {
	if got, _ := actionOf("input"); got != "hub:" {
		t.Fatal(got)
	}
	if got, _ := actionOf("hdmi1"); got != "hdmi:1" {
		t.Fatal(got)
	}
	if got, _ := actionOf("netflix"); got != "app:NETFLIX|netflix" {
		t.Fatal(got)
	}
}
