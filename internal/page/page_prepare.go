package page

import "context"

// Prepare draws the page when it has not been drawn yet.
func Prepare(ctx context.Context, page *Page) error {
	if ctx == nil {
		return errNilContext
	}

	if page == nil {
		return ErrNilPage
	}

	if page.Image() == nil && page.Display() == nil {
		return page.Redraw(ctx)
	}

	return nil
}
