//go:build !js

// Command reload opens the hot reload example in a native window.
// Pass -web to serve the picture in a browser instead, and -reload=false to
// turn the file watch off.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/reload/reload"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8123", "listen address for -web")
	watch := flag.Bool("reload", true, "watch index.html and redraw on a change")
	flag.Parse()

	app, err := reload.New(*watch)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if *webMode {
		if err := gpui.Serve(ctx, app.Page(), *addr); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := gpui.Run(ctx, app.Page()); err != nil {
		log.Fatal(err)
	}
}
