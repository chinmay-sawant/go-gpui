// Package assets contains the generated backpack-cat artwork for examples.
package assets

import "embed"

// Cats contains the transparent PNG character collection.
//
//go:embed "cat_images/*.png"
var Cats embed.FS
