package ui

import (
	"context"
	"sync"
)

// fakeBackend is an in-test Backend that records commands and lets the test
// queue worker updates by hand. Its recording methods live in
// fake_ops_test.go.
type fakeBackend struct {
	mu       sync.Mutex
	queue    []Update
	dark     bool
	started  bool
	closed   bool
	added    []AddRequest
	controls []Control
	pages    []PageRequest
	actives  int
	summary  []uint64
}

func newFake() *fakeBackend { return &fakeBackend{} }

func (f *fakeBackend) Start(_ context.Context) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.started = true
}

func (f *fakeBackend) Info() Info {
	return Info{Dummy: true, DataDir: "/tmp/downloads"}
}

func (f *fakeBackend) Add(req AddRequest) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.added = append(f.added, req)

	return nil
}
