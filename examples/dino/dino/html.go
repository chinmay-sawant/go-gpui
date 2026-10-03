package dino

import (
	"fmt"
	"strings"
)

// buildHTML assembles the scene: the sky marks, the ground, the dinosaur,
// the obstacle slot parts, and the text.
func buildHTML() string {
	var b strings.Builder

	b.WriteString(`<!DOCTYPE html><html><head><meta charset="utf-8">`)
	b.WriteString(`<title>Dino Run</title><style>`)
	b.WriteString(styles)
	b.WriteString(`</style></head><body><div class="scene">`)
	b.WriteString(cloudHTML)
	b.WriteString(`<div class="ground"></div>`)
	b.WriteString(pebbleHTML())
	b.WriteString(dinoHTML)
	b.WriteString(obstacleHTML())
	b.WriteString(textHTML)
	b.WriteString(overlayHTML())
	b.WriteString(`</div></body></html>`)

	return b.String()
}

// pebbleHTML is the six ground marks along the path.
func pebbleHTML() string {
	var b strings.Builder

	for i := range pebbleParts {
		fmt.Fprintf(&b,
			`<div id="p%d" class="pebble" style="left:%dpx;top:242px;width:8px;height:4px"></div>`,
			i, 40+i*140)
	}

	return b.String()
}

// obstacleHTML is the four obstacle slots. Every part starts hidden, so
// the layout is empty until the first tick paints the active ones.
func obstacleHTML() string {
	var b strings.Builder

	for slot := range slotMax {
		for i, part := range obstacleParts {
			x := 210 + i*18 + slot*4
			y := 122 + i%3*22

			fmt.Fprintf(&b,
				`<div id="%s" class="hidden" style="left:%dpx;top:%dpx;width:24px;height:24px"></div>`,
				slotID(slot, part), x, y)
		}
	}

	return b.String()
}
