//go:build !js

// Command audio-player opens the Aurora audio player example.
// Pass -web to serve the picture in a browser instead.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/audio-player/player"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8118", "listen address for -web")
	flag.Parse()

	app, err := player.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	loadCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	if err := app.Load(loadCtx, ""); err != nil {
		log.Printf("audio-player: %v", err)
	}

	cancel()

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
