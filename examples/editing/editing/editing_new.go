package editing

import "github.com/chinmay-sawant/go-gpui"

// New parses the embedded template and registers its handlers.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "Editing",
		HTML:   editingHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	page.Handle(gpui.Handlers{
		Click:      app.onClick,
		BeforeEdit: app.onBeforeEdit,
		Change:     app.onChange,
		SelectAll:  app.onSelectAll,
		Undo:       app.onUndo,
		Redo:       app.onRedo,
	})
	page.SetData(app.view)

	return app, nil
}
