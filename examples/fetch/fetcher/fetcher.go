// Package fetcher is the fetch example.
// The screen runs one GET with ownframe.Fetch or one POST with ownframe.XHR and
// prints the status, byte count, and body head. ownframe opens the window.
// This package does not.
package fetcher

import (
	_ "embed"
	"strconv"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed fetcher.html
var fetcherHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 700
	DefaultHeight = 520

	// bodyHead caps the body text shown in the status line.
	bodyHead = 40
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the fetch screen.
type App struct {
	page *ownframe.Page
	view View
}

// summary formats one response for the status line.
func summary(res ownframe.FetchResponse) string {
	body := string(res.Body)
	if len(body) > bodyHead {
		body = body[:bodyHead]
	}

	return "status=" + strconv.Itoa(res.Status) +
		" bytes=" + strconv.Itoa(len(res.Body)) +
		" body=" + body
}
