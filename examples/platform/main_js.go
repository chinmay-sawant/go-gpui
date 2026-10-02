//go:build js

// Command platform draws the platform example into the browser canvas.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/platform/platform"
)

func main() {
	app, err := platform.New()
	if err != nil {
		log.Fatal(err)
	}

	// On wasm, gpui.Run draws into the browser canvas and blocks.
	if err := gpui.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
