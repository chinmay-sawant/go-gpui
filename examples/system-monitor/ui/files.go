package ui

import "embed"

// components holds the stylesheets and screen fragments.
//
//go:embed components/*.html components/*.css
var components embed.FS

// file returns one embedded component as text.
func file(name string) string {
	b, err := components.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
