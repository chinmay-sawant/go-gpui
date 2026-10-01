// Package login is the sign-in example.
// The screen is an HTML template. Clicks and keystrokes are Go functions.
// gpui opens the window. This package does not.
package login

import (
	"context"
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

// App is the sign-in screen.
type App struct {
	page *gpui.Page
	view View
}

// New parses the embedded login template and registers its handlers.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:     "Sign in",
		HTML:      loginHTML,
		Width:     DefaultWidth,
		Height:    DefaultHeight,
		MinWidth:  MinWidth,
		MinHeight: MinHeight,
		MaxWidth:  MaxWidth,
		MaxHeight: MaxHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	page.Handle(gpui.Handlers{
		Click:     app.onClick,
		Type:      app.onType,
		Backspace: app.onBackspace,
		Submit:    app.onSubmit,
	})
	page.SetData(app.view)

	return app, nil
}

// Page returns the gpui page Run and Serve display.
func (a *App) Page() *gpui.Page {
	return a.page
}

// View returns a copy of the current login fields.
func (a *App) View() View {
	return a.view
}

// Size returns the frame size in CSS pixels.
func (a *App) Size() (int, int) {
	return a.page.Size()
}

// SetSize stores the frame size used by the next Redraw.
func (a *App) SetSize(width, height int) {
	a.page.SetSize(width, height)
}

// Generation increases by one on every successful Redraw.
func (a *App) Generation() uint64 {
	return a.page.Generation()
}

// Redraw fills the template and renders the current size.
func (a *App) Redraw(ctx context.Context) error {
	return a.page.Redraw(ctx)
}

// PNG returns the last PNG, or nil when nothing has been drawn.
func (a *App) PNG() []byte {
	return a.page.PNG()
}

// Boxes returns the last hit-test boxes.
func (a *App) Boxes() []gpui.Box {
	return a.page.Boxes()
}

// Click hit-tests the page and applies the login action.
func (a *App) Click(ctx context.Context, x, y float64) error {
	return a.page.Click(ctx, x, y)
}

// Type appends text to the focused field.
func (a *App) Type(ctx context.Context, text string) error {
	return a.page.Type(ctx, text)
}

// Backspace drops the last rune of the focused field.
func (a *App) Backspace(ctx context.Context) error {
	return a.page.Backspace(ctx)
}

// Submit checks the email and password.
func (a *App) Submit(ctx context.Context) error {
	return a.page.Submit(ctx)
}

func (a *App) onClick(_ context.Context, box gpui.Box) error {
	switch box.Action {
	case "focus":
		a.view.Focus = box.ID
	case "login":
		a.signIn()
	}

	a.page.SetData(a.view)

	return nil
}

func (a *App) onType(_ context.Context, text string) error {
	if a.view.Focus == "" {
		return nil
	}

	if field := a.focused(); field != nil {
		*field += text
	}

	a.page.SetData(a.view)

	return nil
}

func (a *App) onBackspace(_ context.Context) error {
	if a.view.Focus == "" {
		return nil
	}

	if field := a.focused(); field != nil {
		*field = dropLastRune(*field)
	}

	a.page.SetData(a.view)

	return nil
}

func (a *App) onSubmit(_ context.Context) error {
	a.signIn()
	a.page.SetData(a.view)

	return nil
}

func (a *App) signIn() {
	if a.view.Email == "ada@example.com" && a.view.Password == "secret" {
		a.view.Error = ""

		return
	}

	a.view.Error = "Unknown email or password."
}

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
