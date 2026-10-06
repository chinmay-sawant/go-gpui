//go:build !js && !android && !ios

package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/desktop-cat/cat"
)

func runOverlay(ctx context.Context, c *cat.Companion, margin int) error {
	log.Print("desktop cat: click for latest message, drag to move; Ctrl+C closes")
	return ownframe.RunWithOptions(ctx, c.Page, ownframe.WindowOptions{
		Transparent: true, Borderless: true, Floating: true,
		Interactive: c.Interactive, Draggable: c.Draggable,
		FixedSize: true, BottomRight: true, Margin: margin,
	})
}
