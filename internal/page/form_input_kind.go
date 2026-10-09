package page

// isFileInput reports whether c is a file picker. It stays editable so a
// typed name works without a desktop dialog, but it never shows a caret.
func isFileInput(c Control) bool {
	return c.Type == "file"
}

// willChange reports whether a click can change the control's value.
func willChange(c Control) bool {
	switch {
	case c.Type == "checkbox":
		return true
	case c.Type == "radio":
		return !c.Checked
	case c.Tag == "select":
		return len(c.Options) > 1
	default:
		return false
	}
}
