//go:build js

// Command telegram draws the Telegram-like demo into the browser canvas.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/telegram/telegram"
)

func main() {
	app, err := telegram.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := gpui.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
