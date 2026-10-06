//go:build js

// Command telegram draws the Telegram-like demo into the browser canvas.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/telegram/telegram"
)

func main() {
	app, err := telegram.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := ownframe.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
