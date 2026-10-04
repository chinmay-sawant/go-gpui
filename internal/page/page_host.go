package page

import "github.com/chinmay-sawant/go-gpui/internal/host"

var (
	_ host.Focuser         = (*Page)(nil)
	_ host.ContextMenu     = (*Page)(nil)
	_ host.CursorShape     = (*Page)(nil)
	_ host.ScrollRequester = (*Page)(nil)
)
