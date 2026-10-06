package ipc_test

import (
	"context"
	"sync"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/ipc"
)

func TestConcurrentSendListenHandleRequest(t *testing.T) {
	t.Parallel()

	const channel = "race-bus"
	reply := func(_ context.Context, s string) (string, error) {
		return s, nil
	}
	ipc.Handle(channel, reply)

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for n := 0; n < 40; n++ {
				cancel := ipc.Listen(channel, func(string) {})
				ipc.Send(channel, "m")
				cancel()
				_, _ = ipc.Request(context.Background(), channel, "m")
				swap := ipc.Handle(channel, reply)
				swap()
				ipc.Handle(channel, reply)
			}
		}()
	}

	wg.Wait()
}
