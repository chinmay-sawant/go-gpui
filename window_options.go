package gpui

import (
	"context"

	"github.com/chinmay-sawant/go-gpui/internal/filepick"
	pagepkg "github.com/chinmay-sawant/go-gpui/internal/page"
	"github.com/chinmay-sawant/go-gpui/internal/window"
)

// WindowOptions controls desktop window appearance and input.
// The zero value opens the same window as Run.
type WindowOptions = window.Options

// RunWithOptions opens a page with optional desktop window settings.
func RunWithOptions(ctx context.Context, page *Page, options WindowOptions) (err error) {
	defer savePanic(page, &err)

	if err := prepare(ctx, page); err != nil {
		return err
	}

	pagepkg.InstallPicker(page, filepick.Pick)
	ensureAudio()

	return window.RunWithOptions(ctx, page, options)
}
