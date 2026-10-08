package window

import (
	"github.com/chinmay-sawant/blinkless/layout"
)

// devElementRows returns the Elements tab: the picked target above its
// properties as pretty JSON. The pinned box wins over the hovered one; with
// neither, the tab explains how to pick.
func (s *shell) devElementRows() []devRow {
	box, ok := s.devBox()
	if !ok {
		return []devRow{
			{line: devPlain("Select an element")},
			{},
			{line: devText("Click a box in the page to pin it.", devDim)},
			{line: devText("Hover to preview.", devDim)},
		}
	}

	label, ink := "hovering", devDim
	if s.dev.havePin {
		label, ink = "pinned", devAccent
	}

	rows := []devRow{{line: devText(label+"  "+devBoxName(box), ink)}}
	rows = append(rows, devRow{}, devRow{line: devText("properties", devDim)}, devRow{})
	rows = append(rows, devJSONRows(s.devElementPayload(box), s.dev.collapsed)...)

	return rows
}

// devBox returns the picked box, pinned first.
func (s *shell) devBox() (layout.Box, bool) {
	if s.dev.havePin {
		return s.dev.pinned, true
	}

	if s.dev.haveHov {
		return s.dev.hovered, true
	}

	return layout.Box{}, false
}

// devElementPayload is the JSON payload for one box.
func (s *shell) devElementPayload(box layout.Box) devElementJSON {
	ops := 0
	if s.display != nil {
		ops = devOpsInBox(s.display, box)
	}

	return devElementJSON{
		Tag:    box.Tag,
		ID:     box.ID,
		Action: box.Action,
		Text:   devClip(box.Text, 120),
		Rect: devElementRect{
			X: box.X, Y: box.Y, Width: box.W, Height: box.H,
		},
		Ops: ops,
	}
}
