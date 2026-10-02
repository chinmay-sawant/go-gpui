package png

import (
	"context"
	"strconv"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick changes the stamp on Redraw, or writes the current picture on
// Save PNG. The library redraws again after this returns.
func (a *App) onClick(ctx context.Context, box gpui.Box) error {
	switch box.ID {
	case "redraw":
		a.view.Stamp++
		a.view.Status = "stamp " + strconv.Itoa(a.view.Stamp)

		return a.Redraw(ctx)
	case "save":
		path, err := a.SavePNG(".")
		if err != nil {
			return err
		}

		a.view.Status = "wrote " + path
	}

	return nil
}
