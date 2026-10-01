//go:build js

// Command login draws the sign-in example into the browser canvas.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/login/login"
)

func main() {
	app, err := login.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := gpui.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
