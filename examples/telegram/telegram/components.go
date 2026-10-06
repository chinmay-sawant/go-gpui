package telegram

import "embed"

//go:embed components/*.css components/*.html icons/*.svg
var appFS embed.FS

// file returns one embedded component or icon as text.
func file(name string) string {
	b, err := appFS.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
