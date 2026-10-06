package window

import (
	"context"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// menuScreen is a fakeScreen with context menu rows.
type menuScreen struct {
	*fakeScreen
	items  []host.MenuItem
	pasted string
}

func (f *menuScreen) ContextMenu() []host.MenuItem { return f.items }

func (f *menuScreen) Paste(_ context.Context, text string) error {
	f.pasted = text

	return nil
}
