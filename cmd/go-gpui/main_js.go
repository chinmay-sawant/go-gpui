//go:build js

// Command go-gpui draws the login screen into the browser canvas.
package main

import "log"

import "github.com/chinmay-sawant/go-gpui/internal/window"

func main() {
	app := mustApp()

	if err := window.Run(app); err != nil {
		log.Fatal(err)
	}
}
