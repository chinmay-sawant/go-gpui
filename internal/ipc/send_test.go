package ipc_test

import (
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/ipc"
)

func TestSendReachesBothListeners(t *testing.T) {
	t.Parallel()

	var a, b string
	ipc.Listen("send-both", func(s string) { a = s })
	ipc.Listen("send-both", func(s string) { b = s })
	ipc.Send("send-both", "hi")

	if a != "hi" || b != "hi" {
		t.Fatalf("a=%q b=%q", a, b)
	}
}

func TestCancelStopsOneListener(t *testing.T) {
	t.Parallel()

	var a, b int
	stop := ipc.Listen("send-cancel", func(string) { a++ })
	ipc.Listen("send-cancel", func(string) { b++ })
	stop()
	stop()
	ipc.Send("send-cancel", "x")

	if a != 0 || b != 1 {
		t.Fatalf("a=%d b=%d", a, b)
	}
}

func TestPanickingListenerDoesNotBlockTheNext(t *testing.T) {
	t.Parallel()

	var got string
	ipc.Listen("send-panic", func(string) { panic("nope") })
	ipc.Listen("send-panic", func(s string) { got = s })
	ipc.Send("send-panic", "live")

	if got != "live" {
		t.Fatalf("got %q", got)
	}
}
