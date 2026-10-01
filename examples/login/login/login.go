// Package login is the sign-in example.
// The screen is an HTML template. Clicks and keystrokes are Go functions.
// gpui opens the window. This package does not.
package login

import (
	_ "embed"
	"strings"
	"unicode/utf8"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed login.html
var loginHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 480
	DefaultHeight = 640

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 320
	MinHeight = 400

	// MaxWidth and MaxHeight cap a resized frame so a huge monitor does not
	// allocate a matching PNG on every drag.
	MaxWidth  = 2560
	MaxHeight = 2560

	emailField    = "email"
	passwordField = "password"
	undoLimit     = 64
)

// View is the login screen data.
// The template prints Email, PasswordMask, Error, Status, Focus, and Selected.
// Password itself is never written into the HTML.
type View struct {
	Error    string
	Status   string
	Email    string
	Password string
	Focus    string
	Selected bool
}

// PasswordMask returns one '*' for each rune in Password.
func (v View) PasswordMask() string {
	return strings.Repeat("*", utf8.RuneCountInString(v.Password))
}

// App is the sign-in screen.
type App struct {
	page *gpui.Page
	view View
	undo []View
	redo []View
}
