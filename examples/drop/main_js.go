//go:build js

// Command drop draws the drag-and-drop example into the browser canvas.
// The browser delivers dropped files through the same Ebiten call as the
// desktop window.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/drop/drop"
)

func main() {
	app, err := drop.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := gpui.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
