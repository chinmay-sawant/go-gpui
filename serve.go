package gpui

import (
	"context"

	"github.com/chinmay-sawant/go-gpui/internal/web"
)

// Serve listens on addr and blocks.
// The page at / shows the latest picture. Clicks and the type form call the
// same handlers as Run. Serve draws the page first when it has not been drawn yet.
// Developer endpoints stay off; ServeWithOptions opts into them.
func Serve(ctx context.Context, page *Page, addr string) error {
	if err := prepare(ctx, page); err != nil {
		return err
	}

	return web.Serve(page, addr)
}

// ServeOptions tunes ServeWithOptions. It re-exports the web options, so
// callers stay on package gpui.
type ServeOptions = web.Options

// ServeWithOptions listens on addr like Serve, with optional developer
// endpoints. ServeOptions{Perf: true} serves /debug/pprof/*.
func ServeWithOptions(ctx context.Context, page *Page, addr string, opts ServeOptions) error {
	if err := prepare(ctx, page); err != nil {
		return err
	}

	return web.ServeWithOptions(page, addr, opts)
}
