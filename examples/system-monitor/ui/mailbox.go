package ui

import "sync"

// mailbox carries results from the poll goroutines to the tick. Each slot
// keeps the newest item and the generation that produced it, so a result
// from before a mode switch is dropped instead of painted.
type mailbox struct {
	mu       sync.Mutex
	summary  *item[Summary]
	procs    *item[ProcSnapshot]
	tracked  *item[Tracked]
	problems *item[[]string]
}

// item is one mailed result with its generation.
type item[T any] struct {
	gen uint64
	v   T
}

// putSummary stores the newest summary.
func (m *mailbox) putSummary(gen uint64, s Summary) {
	m.mu.Lock()
	m.summary = &item[Summary]{gen: gen, v: s}
	m.mu.Unlock()
}

// putProcs stores the newest process snapshot.
func (m *mailbox) putProcs(gen uint64, s ProcSnapshot) {
	m.mu.Lock()
	m.procs = &item[ProcSnapshot]{gen: gen, v: s}
	m.mu.Unlock()
}

// putTracked stores the newest tracked-process state.
func (m *mailbox) putTracked(gen uint64, t Tracked) {
	m.mu.Lock()
	m.tracked = &item[Tracked]{gen: gen, v: t}
	m.mu.Unlock()
}

// putProblems stores the newest collector problems.
func (m *mailbox) putProblems(gen uint64, lines []string) {
	m.mu.Lock()
	m.problems = &item[[]string]{gen: gen, v: lines}
	m.mu.Unlock()
}

// batch is the bounded result of one drain: at most one item per feed,
// which caps the work one tick can do whatever the source's rate is.
type batch struct {
	summary  *item[Summary]
	procs    *item[ProcSnapshot]
	tracked  *item[Tracked]
	problems *item[[]string]
}

// take returns the newest item of each feed and clears the slots.
func (m *mailbox) take() batch {
	m.mu.Lock()
	defer m.mu.Unlock()

	b := batch{
		summary:  m.summary,
		procs:    m.procs,
		tracked:  m.tracked,
		problems: m.problems,
	}
	m.summary, m.procs, m.tracked, m.problems = nil, nil, nil, nil

	return b
}
