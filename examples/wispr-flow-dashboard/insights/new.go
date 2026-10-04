package insights

import "github.com/chinmay-sawant/go-gpui"

// App is the insights screen.
type App struct {
	page   *gpui.Page
	view   View
	streak int
}

// New parses the dashboard template, registers its images and handlers.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:     "Insights",
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

	app := &App{page: page, view: DefaultView()}
	registerImages(page)
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(app.view)

	return app, nil
}
