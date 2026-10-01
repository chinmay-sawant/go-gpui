//go:build !js

// Command go-gpui opens the login screen in a native window.
// Pass -web to serve the older browser page instead.
package main

import (
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui/internal/web"
	"github.com/chinmay-sawant/go-gpui/internal/window"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8091", "listen address for -web")
	flag.Parse()

	app := mustApp()

	if *webMode {
		if err := web.Serve(app, *addr); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := window.Run(app); err != nil {
		log.Fatal(err)
	}
}
