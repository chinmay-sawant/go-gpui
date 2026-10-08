// Package remote is the LG webOS remote page.
// The desktop build talks to the TV over Wi-Fi. The Android build can also
// send the same keys as a Bluetooth HID keyboard. The theme starts dark.
package remote

import "github.com/chinmay-sawant/ownframe"

const (
	// DefaultWidth and DefaultHeight fit a phone remote in a desktop window.
	DefaultWidth  = 420
	DefaultHeight = 800
)

// Key is one tappable control.
type Key struct {
	ID    string
	Label string
	Class string
}

// View is the data the template prints.
type View struct {
	Title         string
	Status        string
	Mode          string
	ThemeLabel    string
	Hint          string
	ShowBluetooth bool
	Face          []Key
	More          []Key
}

// App is the remote screen.
type App struct {
	page       *ownframe.Page
	view       View
	dark       bool
	host       string
	link       linker
	async      bool
	loaded     bool
	wantSearch bool
	btSeen     string
	notes      chan update
}

type update struct {
	status string
	title  string
	host   string
}

type linker interface {
	Exec(host, spec string) (string, error)
	Scan() (string, error)
	Wake() (string, error)
	SavedHost() string
	SavedModel() string
}

// Page is the ownframe page.
func (a *App) Page() *ownframe.Page { return a.page }

// Status is the line under the title.
func (a *App) Status() string { return a.view.Status }

// Mode is "Wi-Fi" or "Bluetooth".
func (a *App) Mode() string { return a.view.Mode }

// ThemeLabel is the theme button caption.
func (a *App) ThemeLabel() string { return a.view.ThemeLabel }
