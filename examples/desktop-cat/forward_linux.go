//go:build linux && !android

package main

import (
	"context"
	"encoding/json"
	"io"

	"github.com/chinmay-sawant/ownframe/examples/desktop-cat/cat"
)

func forwardNotifications(ctx context.Context, inbox *cat.Inbox, pipe io.Writer) {
	encoder := json.NewEncoder(pipe)
	latest, version := inbox.Latest()
	if version > 1 {
		if err := encoder.Encode(latest); err != nil {
			return
		}
	}
	for {
		select {
		case <-ctx.Done():
			return
		case n := <-inbox.Updates:
			if err := encoder.Encode(n); err != nil {
				return
			}
		}
	}
}

func nativeEndpoint(addr string) string {
	if addr == "off" {
		return ""
	}
	return "http://" + addr + "/notify"
}
