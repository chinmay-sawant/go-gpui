//go:build !js

// Command lg-remote is a Wi-Fi remote for an LG webOS TV.
// Pass -web to serve the same page in a browser.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/lg-remote/remote"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8131", "listen address for -web")
	flag.Parse()

	app, err := remote.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if *webMode {
		if err := ownframe.Serve(ctx, app.Page(), *addr); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := ownframe.Run(ctx, app.Page()); err != nil {
		log.Fatal(err)
	}
}
