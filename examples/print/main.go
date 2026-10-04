//go:build !js

// Command print opens the PDF export example in a native window.
// Pass -web to serve the picture and GET /pdf in a browser instead.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	print "github.com/chinmay-sawant/go-gpui/examples/print/print"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8125", "listen address for -web")
	flag.Parse()

	app, err := print.New(*webMode)
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
