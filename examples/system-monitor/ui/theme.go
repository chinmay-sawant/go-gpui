package ui

// themeSource returns the theme sheet for a mode. The accent color stays
// #2f6fd0 in both themes, so the tick's color match survives a theme switch.
func themeSource(dark bool) string {
	if dark {
		return file("components/dark.css")
	}

	return file("components/light.css")
}
