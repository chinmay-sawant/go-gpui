//go:build js

// Command devtools draws the devtools overlay example into the browser
// canvas. The browser may keep F12 and Ctrl+Shift+I for itself, so
// Config.DevTools is the way in there.
package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/devtools/devtools"
)

func main() {
	app, err := devtools.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := ownframe.Run(context.Background(), app.Page()); err != nil {
		log.Fatal(err)
	}
}
