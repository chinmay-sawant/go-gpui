//go:build js

// Command platform draws the platform example into the browser canvas.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/platform/platform"
)

func main() {
	app, err := platform.New()
	if err != nil {
		log.Fatal(err)
	}

	// On wasm, ownframe.Run draws into the browser canvas and blocks.
	if err := ownframe.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
