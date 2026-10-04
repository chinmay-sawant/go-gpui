package window

import (
	"fmt"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// devOpsLine is the operation view toggle and, when it is on, the counts per
// kind and the total. A fallback page has no display list.
func (s *shell) devOpsLine() string {
	mark := "[ ]"
	if s.dev.ops {
		mark = "[x]"
	}

	line := mark + " ops mode (o)"

	if !s.dev.ops {
		return line
	}

	if s.display == nil {
		return line + "  bitmap fallback"
	}

	c := devCountOps(s.display)

	return fmt.Sprintf("%s  fill %d  stroke %d  line %d  text %d  image %d  grid %d  total %d",
		line, c.Fill, c.Stroke, c.Line, c.Text, c.Image, c.Grid, c.total())
}

// devBoxDetail is the tag, geometry, action, text, and operation count of
// one picked box.
func (s *shell) devBoxDetail(box layout.Box) string {
	name := box.Tag
	if box.ID != "" {
		name += "#" + box.ID
	}

	detail := fmt.Sprintf("%s  x%g y%g %gx%g", name, box.X, box.Y, box.W, box.H)
	if box.Action != "" {
		detail += "  action " + box.Action
	}

	if box.Text != "" {
		detail += fmt.Sprintf("  text %q", box.Text)
	}

	if s.display != nil {
		detail += fmt.Sprintf("  ops %d", devOpsInBox(s.display, box))
	}

	return detail
}
