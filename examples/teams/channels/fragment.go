package channels

import "embed"

//go:embed channels.html channels.css
var files embed.FS

// HTML returns the menu fragment.
func HTML() string { return string(file("channels.html")) }

// CSS returns the menu stylesheet.
func CSS() string { return string(file("channels.css")) }

// file returns one embedded fragment as text.
func file(name string) string {
	b, err := files.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
