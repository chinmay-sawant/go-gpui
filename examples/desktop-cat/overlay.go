//go:build !js && !android && !ios

package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/desktop-cat/cat"
)

func runOverlay(ctx context.Context, c *cat.Companion, margin int) error {
	log.Print("desktop cat: click for latest message, drag to move; Ctrl+C closes")
	return gpui.RunWithOptions(ctx, c.Page, gpui.WindowOptions{
		Transparent: true, Borderless: true, Floating: true,
		Interactive: c.Interactive, Draggable: c.Draggable,
		FixedSize: true, BottomRight: true, Margin: margin,
	})
}
