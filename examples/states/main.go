//go:build !js

// Command states opens the CSS states example in a native window.
// Pass -web to serve the picture in a browser instead.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/states/states"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8107", "listen address for -web")
	flag.Parse()

	app, err := states.New()
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
