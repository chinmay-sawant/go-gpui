package scene

import (
	_ "embed"
	"strings"
)

//go:embed template/game.html
var gameHTML string

//go:embed template/history.html
var historyHTML string

//go:embed template/styles.css
var stylesCSS string

// pageHTML inlines the one stylesheet into a template and fills the
// palette rules. The engine reads one HTML string and never fetches a
// linked stylesheet.
func pageHTML(tpl string) string {
	css := strings.Replace(stylesCSS, "/* palette */", paletteRules(), 1)

	return strings.Replace(tpl, "<!-- stylesheet -->", "<style>"+css+"</style>", 1)
}
