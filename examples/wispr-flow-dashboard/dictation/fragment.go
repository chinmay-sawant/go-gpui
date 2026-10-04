package dictation

import "embed"

//go:embed dictation.html dictation.css
var files embed.FS

// HTML returns the page fragment.
func HTML() string { return string(file("dictation.html")) }

// CSS returns the page stylesheet.
func CSS() string { return string(file("dictation.css")) }

// file returns one embedded fragment as text.
func file(name string) string {
	b, err := files.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
