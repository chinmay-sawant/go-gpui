package render

import "github.com/chinmay-sawant/gowkhtmltopdf/css"

// State carries the runtime pointer and focus state the engine's stateful
// pseudo-classes match: the ids of the focused, hovered, and pressed elements.
// An empty id matches none of them.
type State struct {
	Focus  string
	Hover  string
	Active string
	// Images returns encoded image bytes for one src, such as an <img src>
	// value or a CSS background-image url(...) target. Nil resolves none.
	Images func(src string) ([]byte, error)
}

// options builds the screen cascade options for a viewport.
func (s State) options(width, height int) css.Options {
	return css.Options{
		WidthPx:  width,
		HeightPx: height,
		Media:    "screen",
		Extra:    nil,
		Focus:    s.Focus,
		Hover:    s.Hover,
		Active:   s.Active,
	}
}
