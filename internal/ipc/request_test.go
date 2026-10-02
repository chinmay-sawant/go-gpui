package ipc_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/ipc"
)

func TestRequestRoundTrip(t *testing.T) {
	t.Parallel()

	want := errors.New("nope")
	ipc.Handle("request-echo", func(ctx context.Context, s string) (string, error) {
		if ctx == nil || s == "bad" {
			return "", want
		}

		return "ok:" + s, nil
	})

	got, err := ipc.Request(context.Background(), "request-echo", "ping")
	if err != nil || got != "ok:ping" {
		t.Fatalf("got %q err %v", got, err)
	}

	got, err = ipc.Request(context.Background(), "request-echo", "bad")
	if got != "" || !errors.Is(err, want) {
		t.Fatalf("got %q err %v", got, err)
	}
}

func TestRequestWithoutHandler(t *testing.T) {
	t.Parallel()

	_, err := ipc.Request(context.Background(), "request-none", "x")
	if err != ipc.ErrNoHandler {
		t.Fatalf("err = %v", err)
	}
}

func TestRequestRejectsBadContext(t *testing.T) {
	t.Parallel()

	called := false
	ipc.Handle("request-ctx", func(context.Context, string) (string, error) {
		called = true
		return "ran", nil
	})

	_, err := ipc.Request(nil, "request-ctx", "x")
	if err == nil || !strings.Contains(err.Error(), "nil context") {
		t.Fatalf("nil ctx err = %v", err)
	}

	dead, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = ipc.Request(dead, "request-ctx", "x")
	if err != context.Canceled || called {
		t.Fatalf("err=%v called=%v", err, called)
	}
}
