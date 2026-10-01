package gpui

import (
	"context"

	"github.com/chinmay-sawant/go-gpui/internal/web"
)

// Serve listens on addr and blocks.
// The page at / shows the latest picture. Clicks and the type form call the
// same handlers as Run. Serve draws the page first when it has not been drawn yet.
func Serve(ctx context.Context, page *Page, addr string) error {
	if err := prepare(ctx, page); err != nil {
		return err
	}

	return web.Serve(page, addr)
}
