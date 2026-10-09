package remote

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"
)

func drainBridge(t *testing.T) {
	t.Helper()

	for bridge.Take() != "" {
	}
}

func TestBluetoothSendsHid(t *testing.T) {
	drainBridge(t)
	app := newTest(t, WithBluetooth(true))
	clickID(t, app, "mode")
	if app.Mode() != "Bluetooth" || bridge.Take() != "start" {
		t.Fatalf("mode %s", app.Mode())
	}

	clickID(t, app, "volup")
	if got := bridge.Take(); got != "con:233" {
		t.Fatalf("hid %s", got)
	}

	clickID(t, app, "hotstar")
	if app.Status() != "That control needs Wi-Fi" {
		t.Fatalf("status %s", app.Status())
	}
}

func TestWakeAndKeys(t *testing.T) {
	app := newTest(t)
	fake := &fakeLink{scan: "10.0.0.9"}
	app.SetLink(fake)
	clickID(t, app, "wake")
	if err := app.onTick(context.Background()); err != nil {
		t.Fatal(err)
	}

	if fake.wake != 1 || app.Status() != "Wake sent" {
		t.Fatalf("wake %d %s", fake.wake, app.Status())
	}

	for _, key := range append(faceKeys(), moreKeys()...) {
		if key.ID == "wake" {
			continue
		}

		if _, ok := actionOf(key.ID); !ok {
			t.Fatalf("missing %s", key.ID)
		}
	}
}

func TestChannelSpec(t *testing.T) {
	if got, ok := actionOf("chup"); !ok || got != "ch:Up" {
		t.Fatalf("channel %s", got)
	}

	if got, ok := actionOf("power"); !ok || got != "power:" {
		t.Fatalf("power %s", got)
	}
}
