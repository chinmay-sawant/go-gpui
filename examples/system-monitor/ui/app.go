package ui

import (
	"context"
	"sync"
	"time"

	"github.com/chinmay-sawant/ownframe"
)

// closeBudget bounds how long Close waits for the poll goroutine to stop.
// The window returns as soon as the polls end after their context is
// cancelled, so a healthy close takes well under this.
const closeBudget = 2 * time.Second

// App is the system monitor screen.
type App struct {
	page  *ownframe.Page
	cfg   Config
	src   Source
	view  View
	state *state
	mail  mailbox

	cancel context.CancelFunc
	wg     sync.WaitGroup

	pollGen   uint64
	lastSum   *Summary
	lastProcs *ProcSnapshot
	lastTrk   *Tracked
	lastProb  string
}

// Page returns the page Run displays.
func (a *App) Page() *ownframe.Page {
	return a.page
}

// Close stops the polls, removes the tick, and closes the source. The join
// is bounded by closeBudget; a source that overruns it is not waited on.
func (a *App) Close() error {
	if a.cancel != nil {
		a.cancel()
	}

	done := make(chan struct{})

	go func() {
		a.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(closeBudget):
	}

	a.page.SetTick(nil)

	return a.src.Close()
}

// sync rebuilds the template data and hands it to the page. Handlers call it
// before their automatic Redraw; the tick never rebuilds the template.
func (a *App) sync() {
	a.view = a.buildView()
	a.page.SetData(&a.view)
	a.state.dirtyText = false
}
