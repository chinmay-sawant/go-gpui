package remote

import (
	"testing"
	"time"
)

type motionLink struct {
	fakeLink
	sent chan string
}

func (f *motionLink) Exec(host, spec string) (string, error) {
	f.sent <- spec
	return "Pointer", nil
}

func TestTrackpadCoalescesQueuedMovement(t *testing.T) {
	a := newTest(t, WithPhone(true))
	f := &motionLink{sent: make(chan string, 2)}
	a.SetLink(f)
	a.SetAsync(true)
	started, release := make(chan struct{}), make(chan struct{})
	a.later(func() { close(started); <-release })
	<-started
	a.setSensitivity(200)
	for range 100 {
		a.queueMotion("TV", 1, 2)
	}
	close(release)
	select {
	case got := <-f.sent:
		if got != "move:200,400" {
			t.Fatal(got)
		}
	case <-time.After(time.Second):
		t.Fatal("trackpad stalled")
	}
	select {
	case got := <-f.sent:
		t.Fatalf("stale movement queued: %s", got)
	default:
	}
}

func TestTrackpadPreservesFractionalMovement(t *testing.T) {
	a := newTest(t, WithPhone(true))
	f := &countedLink{}
	a.SetLink(f)
	a.setSensitivity(25)
	for range 16 {
		a.queueMotion("TV", .25, 0)
	}
	if f.calls != 1 || f.spec != "move:1,0" {
		t.Fatalf("calls %d move %s", f.calls, f.spec)
	}
}
