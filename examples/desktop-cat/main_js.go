//go:build js && wasm

package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/desktop-cat/cat"
)

func main() {
	page, err := cat.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := gpui.RunWithOptions(context.Background(), page, gpui.WindowOptions{Transparent: true}); err != nil {
		log.Fatal(err)
	}
}
