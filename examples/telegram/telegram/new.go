package telegram

import "github.com/chinmay-sawant/go-gpui"

// App is the Telegram demo screen.
type App struct {
	page     *gpui.Page
	view     View
	chats    []Chat
	threads  map[string][]Message
	contacts []Contact
}

// New parses the embedded template and registers its images and handlers.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:     "Telegram",
		HTML:      buildHTML(),
		Width:     DefaultWidth,
		Height:    DefaultHeight,
		MinWidth:  MinWidth,
		MinHeight: MinHeight,
	})
	if err != nil {
		return nil, err
	}

	app := seed()
	app.page = page
	app.view.Tab = "chats"
	app.rebuild()
	registerImages(page)
	page.Handle(gpui.Handlers{
		Click:  app.onClick,
		Change: app.onChange,
		Submit: app.onSubmit,
	})
	page.SetData(&app.view)

	return app, nil
}
