package tv

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/ws"
)

func TestVolumeAvoidsBlockingStatusRead(t *testing.T) {
	var requests atomic.Int32
	c := testClient(t, func(req map[string]any) map[string]any {
		requests.Add(1)
		if req["uri"] == "ssap://audio/getVolume" {
			time.Sleep(80 * time.Millisecond)
		}
		return map[string]any{"returnValue": true, "volume": 20}
	})
	start := time.Now()
	if _, err := c.Exec("vol:Up"); err != nil {
		t.Fatal(err)
	}
	if requests.Load() != 1 {
		t.Fatal("volume waits for another network round trip")
	}
	if took := time.Since(start); took > 50*time.Millisecond {
		t.Fatalf("volume %.1fms", float64(took)/float64(time.Millisecond))
	}
	t.Logf("volume dispatch %s", time.Since(start))
}

func TestIdlePointerReusesHealthySocket(t *testing.T) {
	var requests atomic.Int32
	seen := make(chan string, 2)
	url := socketURL(t, func(b []byte) []byte { seen <- string(b); return nil })
	c := testClient(t, func(req map[string]any) map[string]any {
		requests.Add(1)
		time.Sleep(80 * time.Millisecond)
		return map[string]any{"socketPath": url}
	})
	var err error
	c.in, err = ws.Dial(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	c.pointerAt = time.Now().Add(-3 * time.Second)
	start := time.Now()
	if err := c.Button("LEFT"); err != nil {
		t.Fatal(err)
	}
	if requests.Load() != 0 {
		t.Fatal("idle pointer reopened a healthy connection")
	}
	select {
	case <-seen:
	case <-time.After(time.Second):
		t.Fatal("TV received no button")
	}
	t.Logf("pointer dispatch %s", time.Since(start))
}
