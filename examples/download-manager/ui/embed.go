package ui

import "embed"

//go:embed templates/*
var shellFS embed.FS

// file returns one embedded template or stylesheet as text.
func file(name string) string {
	b, err := shellFS.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
