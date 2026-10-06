package scene

import "github.com/chinmay-sawant/ownframe"

// paintTexts refreshes the score, level, lines, status, and theme caption
// from the frame.
func (s *Scene) paintTexts(f Frame) {
	setText(s.ops.score, padScore(f.Score))
	setText(s.ops.level, padLevel(f.Level))
	setText(s.ops.lines, padLines(f.Lines))
	setText(s.ops.status, s.status)
	setText(s.ops.theme, themeLabel(s.dark))
}

// setText changes a text operation only when the string changed.
func setText(op *ownframe.DisplayOp, text string) {
	if op != nil && op.Text != text {
		op.Text = text
	}
}
