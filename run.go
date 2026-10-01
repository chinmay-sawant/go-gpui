package gpui

import (
	"context"

	ebitenmobile "github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/chinmay-sawant/go-gpui/internal/window"
)

// Run opens a window and blocks until it closes.
// On a WebAssembly build the window is the browser canvas.
// Run draws the page first when it has not been drawn yet.
func Run(ctx context.Context, page *Page) error {
	if err := prepare(ctx, page); err != nil {
		return err
	}

	return window.Run(ctx, page)
}

// BindMobile registers the page with Ebitengine's mobile view.
// Call it from the package that ebitenmobile bind compiles.
// Do not call Run from that package. BindMobile draws the page first
// when it has not been drawn yet.
func BindMobile(ctx context.Context, page *Page) error {
	if err := prepare(ctx, page); err != nil {
		return err
	}

	ebitenmobile.SetGame(window.NewGame(ctx, page))

	return nil
}

func prepare(ctx context.Context, page *Page) error {
	if ctx == nil {
		return errNilContext
	}

	if page == nil {
		return ErrNilPage
	}

	if len(page.PNG()) == 0 {
		return page.Redraw(ctx)
	}

	return nil
}
