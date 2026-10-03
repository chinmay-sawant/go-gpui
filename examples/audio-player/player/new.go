package player

import "github.com/chinmay-sawant/go-gpui"

// App is the audio player screen.
type App struct {
	page *gpui.Page
	view View
	base string
}

// New parses the player template, registers its images and handlers.
// New does not fetch; the sample view shows until Load succeeds.
func New() (*App, error) {
	return NewAt(defaultBase)
}

// NewAt is New with a different API base, so tests can point at httptest.
func NewAt(base string) (*App, error) {
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

	app := &App{page: page, view: DefaultView(), base: base}
	registerImages(page)
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(app.view)

	return app, nil
}
