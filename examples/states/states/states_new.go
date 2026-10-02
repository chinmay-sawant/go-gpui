package states

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// New parses the embedded states template and registers its handlers.
func New() (*App, error) {
	p, err := gpui.New(gpui.Config{
		Title:  "CSS states",
		HTML:   statesHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: p}
	app.view.Status = "hover, press, focus, or check a control"
	p.Handle(gpui.Handlers{
		Click:  app.onClick,
		Change: app.onChange,
	})
	p.SetData(&app.view)

	return app, nil
}

// onClick notes which control a click landed on.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	a.view.Status = "clicked " + box.ID

	return nil
}

// onChange notes which control changed before the page redraws.
func (a *App) onChange(_ context.Context, box gpui.Box) error {
	a.view.Status = "changed " + box.ID

	return nil
}
