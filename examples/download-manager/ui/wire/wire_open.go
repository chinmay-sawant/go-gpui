package wire

import (
	"context"
	"log"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/scheduler"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// New opens the store, builds the transport and engine, and returns the
// backend. Dummy mode seeds deterministic jobs.
func New(ctx context.Context, cfg Config) (*Backend, error) {
	b := &Backend{
		cmds:    make(chan command, commandQueue),
		results: make(chan ui.Update, resultQueue),
		done:    make(chan struct{}),
		speeds:  map[string]speed{},
		dummy:   cfg.Dummy,
	}

	if err := b.open(ctx, cfg); err != nil {
		return nil, err
	}

	if cfg.Fixture {
		b.startFixture()
	}

	b.dark = readDark(b.themePath())

	return b, nil
}

// open picks the storage directory and builds the engine, falling back to
// an in-memory database when the directory will not open.
func (b *Backend) open(ctx context.Context, cfg Config) error {
	if dir := b.pickDir(cfg.DataDir); dir != "" {
		if err := b.openDir(dir); err != nil {
			return err
		}
	}

	if b.store == nil {
		st, err := store.OpenMemory()
		if err != nil {
			return err
		}

		b.store = st
		log.Printf("download-manager: storage unavailable, using memory")
	}

	if _, err := b.store.Reconcile(ctx); err != nil {
		log.Printf("download-manager: reconcile: %v", err)
	}

	if cfg.Dummy {
		if err := b.store.SeedDummy(ctx, seedJobs); err != nil {
			log.Printf("download-manager: seed: %v", err)
		}
	}

	var tr transfer.Transport = transfer.NewFake(transfer.FakeOptions{})
	if !cfg.Dummy {
		tr = transfer.NewHTTP(transfer.HTTPOptions{})
	}

	eng, err := scheduler.New(scheduler.Options{Transport: tr, Store: b.store})
	if err != nil {
		return err
	}

	b.eng = eng
	if err := eng.Recover(ctx); err != nil {
		log.Printf("download-manager: recover: %v", err)
	}

	return nil
}
