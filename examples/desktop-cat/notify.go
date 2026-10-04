//go:build !js && !android && !ios

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/chinmay-sawant/go-gpui/examples/desktop-cat/cat"
)

func startNotifications(ctx context.Context, addr string, inbox *cat.Inbox) (func(), error) {
	if addr == "off" {
		return func() {}, nil
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("notification server: %w", err)
	}
	server := &http.Server{Handler: inbox.Handler(), ReadHeaderTimeout: 3 * time.Second, ReadTimeout: 5 * time.Second, WriteTimeout: 5 * time.Second}
	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Print(err)
		}
	}()
	log.Printf("desktop cat notifications: http://%s", listener.Addr())
	stop := func() { _ = server.Close() }
	go func() { <-ctx.Done(); stop() }()
	return stop, nil
}

func readNotifications(inbox *cat.Inbox) {
	decoder := json.NewDecoder(os.Stdin)
	for {
		var n cat.Notification
		if err := decoder.Decode(&n); err != nil {
			return
		}
		if err := inbox.Submit(n); err != nil {
			log.Print(err)
		}
	}
}
