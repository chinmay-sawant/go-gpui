package wire

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// The queue sizes and shutdown budgets. The budgets are the documented
// shutdown contract: the worker stops within one second, and the engine
// gets a six-second context for its own five-second transfer budget.
const (
	commandQueue = 64
	resultQueue  = 512
	seedJobs     = 100
	workerBudget = time.Second
	closeBudget  = 6 * time.Second
)

// Info reports the mode and storage location the header shows.
func (b *Backend) Info() ui.Info {
	dir := b.dir
	if dir == "" {
		dir = "in-memory (temporary)"
	}

	return ui.Info{Dummy: b.dummy, DataDir: dir}
}
