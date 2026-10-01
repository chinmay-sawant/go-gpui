// Package login is the sign-in screen.
// The screen is an HTML template. Clicks and keystrokes are Go methods.
// This package does not open a window and does not listen on a port.
package login

import (
	"context"
	_ "embed"
	"html/template"
	"strings"
	"unicode/utf8"

	"github.com/chinmay-sawant/gowkhtmltopdf/screen"
)

//go:embed login.html
var loginHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 480
	DefaultHeight = 640

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	// The desktop window refuses to shrink past these.
	MinWidth  = 320
	MinHeight = 400

	// MaxWidth and MaxHeight cap a resized frame so a huge monitor does not
	// allocate a matching PNG on every drag.
	MaxWidth  = 2560
	MaxHeight = 2560

	// emailField and passwordField match the id attributes in login.html.
	emailField    = "email"
	passwordField = "password"
)

// View is the login screen data.
// The template prints Email, PasswordMask, Error, and Focus.
// Password itself is never written into the HTML.
type View struct {
	Error    string
	Email    string
	Password string
	Focus    string
}

// PasswordMask returns one '*' for each rune in Password.
func (v View) PasswordMask() string {
	return strings.Repeat("*", utf8.RuneCountInString(v.Password))
}

// App holds the login template, the current fields, and the last frame.
type App struct {
	tpl        *template.Template
	view       View
	frame      *screen.Frame
	width      int
	height     int
	generation uint64
}

// New parses the embedded login.html template.
func New() (*App, error) {
	tpl, err := template.New("login.html").Parse(loginHTML)
	if err != nil {
		return nil, err
	}

	return &App{
		tpl:    tpl,
		width:  DefaultWidth,
		height: DefaultHeight,
	}, nil
}

// View returns a copy of the current login fields.
func (a *App) View() View {
	return a.view
}

// Size returns the frame size in CSS pixels.
func (a *App) Size() (int, int) {
	return a.width, a.height
}

// ClampSize limits a window size to the frame this screen will draw.
func ClampSize(width, height int) (int, int) {
	return clamp(width, MinWidth, MaxWidth), clamp(height, MinHeight, MaxHeight)
}

// SetSize stores the frame size used by the next Redraw.
// Values outside the min and max are pulled back inside that range.
func (a *App) SetSize(width, height int) {
	a.width, a.height = ClampSize(width, height)
}

// Generation increases by one on every successful Redraw.
func (a *App) Generation() uint64 {
	return a.generation
}

// Redraw fills the template and renders the current size.
func (a *App) Redraw(ctx context.Context) error {
	var body strings.Builder
	if err := a.tpl.Execute(&body, a.view); err != nil {
		return err
	}

	frame, err := screen.Render(
		ctx,
		[]byte(body.String()),
		a.width,
		a.height,
	)
	if err != nil {
		return err
	}

	a.frame = frame
	a.generation++

	return nil
}

// Submit checks the email and password, then redraws.
func (a *App) Submit(ctx context.Context) error {
	a.login()

	return a.Redraw(ctx)
}

// Click hit-tests the last frame. The innermost box is the last one in
// document order that contains the point. focus stores that box id.
// login checks the email and password. Click then redraws.
func (a *App) Click(ctx context.Context, x, y float64) error {
	if box, ok := hit(a.Boxes(), x, y); ok {
		switch box.Action {
		case "focus":
			a.view.Focus = box.ID
		case "login":
			a.login()
		}
	}

	return a.Redraw(ctx)
}

// Type appends text to the focused field and redraws.
// It does nothing when no field is focused.
func (a *App) Type(ctx context.Context, text string) error {
	if a.view.Focus == "" {
		return nil
	}

	if field := a.focused(); field != nil {
		*field += text
	}

	return a.Redraw(ctx)
}

// Backspace drops the last rune of the focused field and redraws.
// It does nothing when no field is focused.
func (a *App) Backspace(ctx context.Context) error {
	if a.view.Focus == "" {
		return nil
	}

	if field := a.focused(); field != nil {
		*field = dropLastRune(*field)
	}

	return a.Redraw(ctx)
}

// PNG returns the last PNG, or nil when nothing has been drawn.
func (a *App) PNG() []byte {
	if a.frame == nil {
		return nil
	}

	return a.frame.PNG
}

// Boxes returns the last hit-test boxes, or nil when nothing has been drawn.
func (a *App) Boxes() []screen.Box {
	if a.frame == nil {
		return nil
	}

	return a.frame.Boxes
}

func (a *App) login() {
	if a.view.Email == "ada@example.com" && a.view.Password == "secret" {
		a.view.Error = ""

		return
	}

	a.view.Error = "Unknown email or password."
}

// focused returns the Email or Password field selected by Focus.
// The ids are the ones in login.html.
func (a *App) focused() *string {
	switch a.view.Focus {
	case emailField:
		return &a.view.Email
	case passwordField:
		return &a.view.Password
	default:
		return nil
	}
}

func dropLastRune(s string) string {
	if s == "" {
		return ""
	}

	_, size := utf8.DecodeLastRuneInString(s)

	return s[:len(s)-size]
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

// hit returns the last box that contains x, y.
// Boxes are in document order, so that box is the inner element.
func hit(boxes []screen.Box, x, y float64) (screen.Box, bool) {
	var found screen.Box
	ok := false

	for _, b := range boxes {
		if x < b.X || y < b.Y || x > b.X+b.W || y > b.Y+b.H {
			continue
		}

		found = b
		ok = true
	}

	return found, ok
}
