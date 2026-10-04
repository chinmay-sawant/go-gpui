package drop

import "github.com/chinmay-sawant/go-gpui"

// New parses the embedded template and registers the drop handler.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "Drag and drop",
		HTML:   dropHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	page.Handle(gpui.Handlers{Drop: app.onDrop, Change: app.onChange})
	page.SetData(&app.view)

	return app, nil
}
