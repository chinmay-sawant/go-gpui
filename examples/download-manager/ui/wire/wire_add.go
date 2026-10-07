package wire

import (
	"path/filepath"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/scheduler"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// runAdd stores the job and refreshes the active set.
func (b *Backend) runAdd(req ui.AddRequest) {
	dir := req.Destination
	if dir == "" {
		dir = b.defaultDir()
	}

	_, err := b.eng.Add(b.ctx, scheduler.AddRequest{URL: req.URL, Dir: dir})
	if err != nil {
		b.notice("Add failed: " + err.Error())

		return
	}

	b.pushActive()
}

// defaultDir is the destination used when the folder field is empty.
func (b *Backend) defaultDir() string {
	if b.dummy || b.dir == "" {
		return transfer.DummyDir()
	}

	return filepath.Join(b.dir, "downloads")
}
