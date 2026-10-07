package ui

import "strings"

// buildHTML returns the page with the shared sheet in the document, so
// control colours beat the widget defaults. The theme only sets variables.
func buildHTML() string {
	style := "<style>" + file("templates/base.css") + "</style>"

	return strings.Replace(file("templates/page.html"), "</head>", style+"</head>", 1)
}

// themeSource returns the colour variables for one mode. Geometry stays in
// the document sheet, so a switch paints without moving boxes.
func themeSource(dark bool) string {
	if dark {
		return file("templates/dark.css")
	}

	return file("templates/light.css")
}
