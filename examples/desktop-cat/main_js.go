//go:build js && wasm

package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/desktop-cat/cat"
)

func main() {
	page, err := cat.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := ownframe.RunWithOptions(context.Background(), page, ownframe.WindowOptions{Transparent: true}); err != nil {
		log.Fatal(err)
	}
}
