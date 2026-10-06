// Command perf-stress opens a dashboard heavier than the large and flappy
// benchmarks: sidebar, header, cards, a 280-row grid, form controls, emoji,
// and a tick that moves a progress bar and an equalizer without a Redraw.
// Pass -web to serve a still picture in a browser instead.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui"
)

func main() {
	web := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8134", "listen address for -web")
	flag.Parse()

	app, err := newApp(280)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if *web {
		if err := gpui.Serve(ctx, app.page, *addr); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := gpui.Run(ctx, app.page); err != nil {
		log.Fatal(err)
	}
}
