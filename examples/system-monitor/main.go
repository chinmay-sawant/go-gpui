//go:build !js

// Command system-monitor opens the system monitor example. The default
// dummy mode needs no permissions; -mode live reads the host instead.
package main

import (
	"flag"
	"log"
)

func main() {
	data := flag.String("data", "", "data directory (default: user config dir)")
	mode := flag.String("mode", "dummy", "data source: dummy or live")
	stress := flag.Bool("stress", false, "use the 10,000 process dummy fixture")
	seed := flag.Int64("seed", 1, "dummy fixture seed")
	flag.Parse()

	if err := run(*data, *mode, *stress, *seed); err != nil {
		log.Fatal(err)
	}
}
