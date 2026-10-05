//go:build js

// Command input-lab draws the combined input window on canvas.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/input-lab/inputlab"
)

func main() {
	app, err := inputlab.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := gpui.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
