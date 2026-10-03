package controls

import "github.com/chinmay-sawant/go-gpui"

// New parses the embedded controls template and registers its handlers.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "Controls",
		HTML:   controlsHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	page.Handle(gpui.Handlers{Click: app.onClick, Change: app.onChange})
	page.SetData(app.view)

	return app, nil
}
