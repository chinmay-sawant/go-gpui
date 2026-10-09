package remote

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"
)

func TestRapidCommandsKeepOrder(t *testing.T) {
	a := newTest(t)
	a.SetLink(&fakeLink{})
	a.SetAsync(true)
	done := make(chan int, 10)
	for i := range 10 {
		a.later(func() { done <- i })
	}
	for want := range 10 {
		select {
		case got := <-done:
			if got != want {
				t.Fatalf("got %d want %d", got, want)
			}
		case <-time.After(time.Second):
			t.Fatal("command worker stalled")
		}
	}
}

func TestRepeatedBluetoothResultsAreDisplayed(t *testing.T) {
	drainBridge(t)
	a := newTest(t, WithPhone(true))
	clickID(t, a, "mode")
	ctx := context.Background()
	for range 2 {
		clickID(t, a, "volup")
		bridge.SetBluetooth("Pair this phone in the TV Bluetooth list.")
		if err := a.onTick(ctx); err != nil {
			t.Fatal(err)
		}
		if a.Status() != "Pair this phone in the TV Bluetooth list." {
			t.Fatal(a.Status())
		}
	}
	drainBridge(t)
	for i := range 32 {
		if !bridge.Enqueue(fmt.Sprint(i)) {
			t.Fatal("early queue rejection")
		}
	}
	if bridge.Enqueue("overflow") {
		t.Fatal("queue silently discarded a command")
	}
	if bridge.Take() != "0" {
		t.Fatal("oldest command was lost")
	}
	drainBridge(t)
}
