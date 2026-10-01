package page

import "context"

// Copy returns the text the page wants on the clipboard.
// ok is false when there is nothing to copy. Copy does not draw.
func (p *Page) Copy(ctx context.Context) (string, bool, error) {
	if err := useContext(ctx); err != nil {
		return "", false, err
	}

	if p.handlers.Copy == nil {
		return "", false, nil
	}

	return p.handlers.Copy(ctx)
}

// Cut returns the text the page wants on the clipboard and draws again.
// ok is false when there is nothing to cut.
func (p *Page) Cut(ctx context.Context) (string, bool, error) {
	if err := useContext(ctx); err != nil {
		return "", false, err
	}

	if p.handlers.Cut == nil {
		return "", false, nil
	}

	text, ok, err := p.handlers.Cut(ctx)
	if err != nil || !ok {
		return text, ok, err
	}

	if err := p.Redraw(ctx); err != nil {
		return text, ok, err
	}

	return text, ok, nil
}

func useContext(ctx context.Context) error {
	if ctx == nil {
		return errNilContext
	}

	return ctx.Err()
}
