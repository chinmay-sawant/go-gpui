package app

import "embed"

//go:embed components/*.html components/*.css assets/*.svg assets_dark/*.svg
var shellFS embed.FS

// file returns one embedded component or asset as text.
func file(name string) string {
	b, err := shellFS.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
