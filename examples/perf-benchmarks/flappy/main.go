//go:build !js

// Command flappy opens the Benchmark C tick animation in a native window.
// Pass -web to serve the last Redraw in a browser instead: Serve never
// ticks, so the browser shows a still picture.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/go-gpui/examples/perf-benchmarks/benchutil"
)

func main() {
	web, addr := benchutil.Flags("127.0.0.1:8133")
	flag.Parse()

	app, err := benchutil.NewFlap()
	if err != nil {
		log.Fatal(err)
	}

	if err := benchutil.ServeOrRun(context.Background(), app.Page(), *web, *addr); err != nil {
		log.Fatal(err)
	}
}
