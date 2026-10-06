package ui

import (
	"context"
	"strings"
	"time"
)

// pollEvery is how often the pump reads the source's cached state. The
// collector publishes about once a second, so this only needs to beat a
// frame; it is not the sampling rate.
const pollEvery = 400 * time.Millisecond

// startPumps launches the poll goroutine. Every blocking source call lives
// here, never on the UI loop.
func (a *App) startPumps(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	a.cancel = cancel

	a.wg.Add(1)

	go func() {
		defer a.wg.Done()

		a.poll(ctx)

		t := time.NewTicker(pollEvery)
		defer t.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				a.poll(ctx)
			}
		}
	}()
}

// poll copies the source's latest state into the mailbox. Every item is
// tagged with the generation the poll saw, so a mode switch discards results
// already in flight. An unchanged sample is not reposted.
func (a *App) poll(ctx context.Context) {
	gen := a.currentGen()

	if gen != a.pollGen {
		a.pollGen = gen
		a.lastSum, a.lastProcs, a.lastTrk = nil, nil, nil
		a.lastProb = ""
	}

	if s, ok := a.src.Summary(ctx); ok && (a.lastSum == nil || !a.lastSum.At.Equal(s.At)) {
		a.lastSum = &s
		a.mail.putSummary(gen, s)
	}

	if p, ok := a.src.Processes(ctx); ok && (a.lastProcs == nil || !a.lastProcs.At.Equal(p.At)) {
		a.lastProcs = &p
		a.mail.putProcs(gen, p)
	}

	if t := a.src.Tracked(ctx); a.lastTrk == nil || *a.lastTrk != t {
		a.lastTrk = &t
		a.mail.putTracked(gen, t)
	}

	if lines := a.src.Problems(ctx); problemsKey(lines) != a.lastProb {
		a.lastProb = problemsKey(lines)
		a.mail.putProblems(gen, lines)
	}
}

// problemsKey collapses problem lines for the unchanged check.
func problemsKey(lines []string) string {
	return strings.Join(lines, "\n")
}
