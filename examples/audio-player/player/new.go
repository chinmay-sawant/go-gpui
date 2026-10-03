package player

import (
	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/music"
)

// App is the audio player screen.
type App struct {
	page     *gpui.Page
	view     View
	base     string
	audio    *music.Engine
	audioKey string
}

// New parses the player template, registers its images and handlers.
// New does not fetch; the sample view shows until Load succeeds.
func New() (*App, error) {
	return NewAt(defaultBase)
}

// NewAt is New with a different API base, so tests can point at httptest.
func NewAt(base string) (*App, error) {
	return NewWith(base, music.NewEngine(&music.Library{Primary: music.NewOpenverse()}))
}

// NewWith is NewAt with a caller-supplied audio engine. Tests replace Open
// with a fake voice so no audio device is needed.
func NewWith(base string, engine *music.Engine) (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:     "Aurora — Audio player",
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
	registerImages(page)
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetTick(app.Tick)
	page.SetData(app.view)

	return app, nil
}

// Close releases the audio voice.
func (a *App) Close() {
	if a.audio != nil {
		a.audio.Close()
	}
}
