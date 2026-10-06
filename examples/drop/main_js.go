//go:build js

// Command drop draws the drag-and-drop example into the browser canvas.
// The browser delivers dropped files through the same Ebiten call as the
// desktop window.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/drop/drop"
)

func main() {
	app, err := drop.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := ownframe.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
