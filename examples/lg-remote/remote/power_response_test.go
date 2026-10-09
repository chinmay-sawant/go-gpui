package remote

import (
	"testing"
)

type blockedPowerLink struct {
	fakeLink
	gate chan struct{}
}

func (f *blockedPowerLink) Exec(host, spec string) (string, error) {
	<-f.gate
	return "Sent", nil
}

func (f *blockedPowerLink) Wake() (string, error) {
	<-f.gate
	return "Wake sent", nil
}

func TestPowerAndWakeRespondBeforeNetwork(t *testing.T) {
	a := newTest(t, WithPhone(true))
	f := &blockedPowerLink{gate: make(chan struct{})}
	a.SetLink(f)
	a.SetAsync(true)
	defer close(f.gate)
	clickID(t, a, "power")
	if a.view.PowerOn {
		t.Fatal("power colour waits for network response")
	}
	clickID(t, a, "wake")
	if !a.view.PowerOn {
		t.Fatal("wake leaves power colour stale")
	}
}
