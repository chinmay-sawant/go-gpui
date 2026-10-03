package login

import "github.com/chinmay-sawant/go-gpui"

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
		Click:      app.onClick,
		Submit:     app.onSubmit,
		BeforeEdit: app.onBeforeEdit,
		Change:     app.onChange,
		Undo:       app.onUndo,
		Redo:       app.onRedo,
	})
	page.SetData(app.view)

	return app, nil
}
