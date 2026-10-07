package ui

// buildHTML returns the one page template.
func buildHTML() string {
	return file("templates/page.html")
}

// themeSource returns the shared base stylesheet plus the mode overrides.
// Both modes carry the same geometry, so a switch paints without moving
// boxes; only the colour variables differ.
func themeSource(dark bool) string {
	base := file("templates/base.css")
	if dark {
		return base + "\n" + file("templates/dark.css")
	}

	return base + "\n" + file("templates/light.css")
}
