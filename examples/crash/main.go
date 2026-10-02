//go:build !js

// Command crash opens the crash-report example in a native window.
// Pass -web to serve the picture in a browser instead.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/crash/crash"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8114", "listen address for -web")
	flag.Parse()

	// SetCrashDir points Report and a recovered panic at crashes/ under the
	// working directory. An empty dir restores the per-user default.
	gpui.SetCrashDir("crashes")

	app, err := crash.New()
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

	// Run recovers a panic from the page, writes a report under crashes/,
	// and returns an error that names the report file. log.Fatal prints it.
	if err := gpui.Run(ctx, app.Page()); err != nil {
		log.Fatal(err)
	}
}
