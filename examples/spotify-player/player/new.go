package player

import (
	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/music"
)

// App is the player screen.
type App struct {
	page       *gpui.Page
	view       View
	base       string
	seconds    int
	lastVolume int
	audio      *music.Engine
	audioPick  int
	audioKey   string
}

// New parses the player template, registers its images, and starts on the
// sample view. New does not fetch; main calls Load for live data.
func New() (*App, error) {
	return NewAt(itunesBase)
}

// NewAt is New with a different Search API base URL. Tests pass a local
// server so they never reach iTunes.
func NewAt(base string) (*App, error) {
	return NewWith(base, music.NewEngine(&music.Library{Primary: music.NewOpenverse()}))
}

// NewWith is NewAt with an injected audio engine, so tests never open the
// OS audio device.
func NewWith(base string, engine *music.Engine) (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:     "Spotify — Player",
		HTML:      buildHTML(),
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

	app := &App{page: page, view: DefaultView(), base: base, audio: engine}
	app.seconds = lengthSeconds(app.view.Now.Length)
	registerImages(page)
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetTick(app.Tick)
	page.SetData(app.view)

	return app, nil
}
