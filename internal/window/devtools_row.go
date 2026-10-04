package window

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// devSpan is one run on a panel line: text in one ink, or a filled swatch
// when swatch is positive.
type devSpan struct {
	text   string
	ink    color.RGBA
	swatch float64
}

// devLine is one panel line, drawn left to right.
type devLine struct {
	spans []devSpan
}

// devAct is what a panel row does when a click lands on it.
type devAct int

const (
	devActNone      devAct = iota
	devActTab              // select the tab named by arg
	devActOpsToggle        // flip the operation outlines
	devActJSON             // collapse or expand the JSON node named by key
	devActOp               // outline the operation at index arg
)

// devRow is one content line plus the click action it carries.
type devRow struct {
	line devLine
	act  devAct
	arg  int
	key  string
}

// devSwatchSize is the side of a colour chip in a panel line.
const devSwatchSize = 9.0

// devSwatchGap is the space between a swatch and the text after it.
const devSwatchGap = 6.0

// devText builds a one-span line.
func devText(s string, ink color.RGBA) devLine {
	return devLine{spans: []devSpan{{text: s, ink: ink}}}
}

// devPlain builds a line in the default panel ink.
func devPlain(s string) devLine {
	return devText(s, devFg)
}

// devSwatch builds a colour chip span.
func devSwatch(ink color.RGBA) devSpan {
	return devSpan{ink: ink, swatch: devSwatchSize}
}

// plain is the line without colour, for labels and tests.
func (l devLine) plain() string {
	out := ""
	for _, span := range l.spans {
		out += span.text
	}

	return out
}

// width is the measured width of the line, swatches included.
func (l devLine) width() float64 {
	w := 0.0

	for _, span := range l.spans {
		if span.swatch > 0 {
			w += span.swatch + devSwatchGap

			continue
		}

		sw, _ := text.Measure(span.text, badgeFace, 0)
		w += sw
	}

	return w
}
