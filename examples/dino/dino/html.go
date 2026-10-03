package dino

import (
	_ "embed"
	"strings"
)

//go:embed template/index.html
var indexHTML string

//go:embed template/styles.css
var stylesCSS string

// buildHTML assembles the page. The stylesheet is inlined at the marker
// because the engine reads one HTML string and never fetches a linked
// stylesheet.
func buildHTML() string {
	return strings.Replace(indexHTML, "<!-- stylesheet -->", "<style>"+stylesCSS+"</style>", 1)
}
