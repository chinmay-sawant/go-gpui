package telegram

// themeSource returns the extra stylesheet for the chosen look. The template
// sheet is the light look, so light needs no theme.
func themeSource(dark bool) string {
	if dark {
		return file("components/dark.css")
	}

	return ""
}
