package png

import "github.com/chinmay-sawant/go-gpui"

// New parses the embedded png template and registers its click handler.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "PNG",
		HTML:   pngHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(&app.view)

	return app, nil
}
