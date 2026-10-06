package drop

import (
	"context"

	"github.com/chinmay-sawant/ownframe"
)

// onDrop adds the path of every dropped file. The whole window takes a drop,
// so any file can land anywhere in it. Path is absolute on desktop; in a
// browser it is empty and the entry's name stands in.
func (a *App) onDrop(_ context.Context, files []ownframe.Drop) error {
	for _, file := range files {
		a.add(file.Path, file.Name)
	}

	return nil
}
