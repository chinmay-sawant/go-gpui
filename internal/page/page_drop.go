package page

import (
	"context"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// Drop is one file or directory dropped on the window.
type Drop = host.Drop

// Drop calls the drop handler with the files dropped on the window and draws
// the page again. An error returns before the redraw.
func (p *Page) Drop(ctx context.Context, files []Drop) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if p.handlers.Drop != nil {
		if err := p.handlers.Drop(ctx, files); err != nil {
			return err
		}
	}

	return p.Redraw(ctx)
}

var _ host.Dropper = (*Page)(nil)
