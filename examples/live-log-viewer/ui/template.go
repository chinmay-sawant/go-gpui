package ui

import _ "embed"

//go:embed ui.html
var pageHTML string

// buildHTML returns the page source. It is a function so later fragments
// can be composed without touching callers.
func buildHTML() string { return pageHTML }
