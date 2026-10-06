//go:build !js

// Command normal opens the Benchmark A desktop form in a native window.
// Pass -web to serve the picture in a browser instead.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/chinmay-sawant/ownframe/examples/perf-benchmarks/benchutil"
)

func main() {
	web, addr := benchutil.Flags("127.0.0.1:8131")
	flag.Parse()

	app, err := benchutil.NewNormal()
	if err != nil {
		log.Fatal(err)
	}

	if err := benchutil.ServeOrRun(context.Background(), app.Page(), *web, *addr); err != nil {
		log.Fatal(err)
	}
}
