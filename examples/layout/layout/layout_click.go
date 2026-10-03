package layout

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick writes the clicked box's tag, id, and geometry into the status
// line. The library redraws after this returns, so the template prints it.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	a.view.Status = fmt.Sprintf("tag=%s id=%s x=%.0f y=%.0f w=%.0f h=%.0f",
		box.Tag, box.ID, box.X, box.Y, box.W, box.H)

	return nil
}
