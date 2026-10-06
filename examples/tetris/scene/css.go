package scene

import (
	"strconv"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// paletteRules are the template's cell colour rules. Each rule reads a
// custom property with the light default, so a theme sheet can restyle
// the page without touching geometry.
func paletteRules() string {
	p := paletteFor(false)
	var b strings.Builder

	b.WriteString(".cell{background:var(--cell," + p.cell.hex() + ");}\n")

	for k := game.PieceI; k <= game.PieceL; k++ {
		cl := string(letter(k))
		b.WriteString(".k-" + cl + "{background:var(--k-" + cl + "," + p.kinds[k].hex() + ");}\n")
	}

	b.WriteString(".k-g{background:var(--k-g," + p.ghostCell.hex() + ");}\n")

	return b.String()
}

// themeCSS is the stylesheet for one theme: the custom properties that
// every template rule reads.
func themeCSS(dark bool) string {
	p := paletteFor(dark)
	var b strings.Builder

	b.WriteString(":root{")
	b.WriteString("--bg:" + p.bg.hex() + ";")
	b.WriteString("--panel:" + p.panel.hex() + ";")
	b.WriteString("--board:" + p.board.hex() + ";")
	b.WriteString("--cell:" + p.cell.hex() + ";")
	b.WriteString("--ink:" + p.ink.hex() + ";")
	b.WriteString("--muted:" + p.muted.hex() + ";")
	b.WriteString("--btn:" + p.btn.hex() + ";")
	b.WriteString("--over:rgba(" + strconv.Itoa(int(p.over.r)) + "," +
		strconv.Itoa(int(p.over.g)) + "," + strconv.Itoa(int(p.over.b)) + "," +
		strconv.FormatFloat(p.overAlpha, 'g', 2, 64) + ");")

	for k := game.PieceI; k <= game.PieceL; k++ {
		cl := string(letter(k))
		b.WriteString("--k-" + cl + ":" + p.kinds[k].hex() + ";")
	}

	b.WriteString("--k-g:" + p.ghostCell.hex() + ";}")
	b.WriteString("html,body{background:var(--bg);}")

	return b.String()
}
