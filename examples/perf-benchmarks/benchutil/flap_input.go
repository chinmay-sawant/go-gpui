package benchutil

import (
	"context"

	"github.com/chinmay-sawant/ownframe"
)

// onKey flaps on space and never draws; the tick paints the result.
func (f *Flap) onKey(_ context.Context, key string) error {
	if key == "space" {
		f.flap()
	}

	return nil
}

// onClick flaps from a pointer press.
func (f *Flap) onClick(_ context.Context, _ ownframe.Box) error {
	f.flap()

	return nil
}
