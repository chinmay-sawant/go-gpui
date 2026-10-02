package page

// Control is one form field taken from the executed HTML.
// A control with an empty id is ignored.
type Control struct {
	ID       string
	Tag      string
	Type     string
	Name     string
	Value    string
	Checked  bool
	Disabled bool
	Options  []Option
}

// Option is one choice inside a select.
type Option struct {
	Value    string
	Label    string
	Selected bool
}

// controlSpan is one control and where it sits in the executed HTML.
// Start is the index of '<'. End is the index after the element.
type controlSpan struct {
	Control Control
	Start   int
	End     int
}

// formState is the value of each control on the current document.
type formState struct {
	byID     map[string]Control
	order    []string
	focusID  string
	selected bool
	doc      int
}

// textLike reports whether typing edits this control.
// An empty type is a text input.
func textLike(kind string) bool {
	switch kind {
	case "", "text", "password", "email", "search", "tel", "url", "number", "file":
		return true
	default:
		return false
	}
}
