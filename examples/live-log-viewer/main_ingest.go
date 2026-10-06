package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// stopBudget bounds how long shutdown waits for the ingestors.
const stopBudget = 2 * time.Second

// startIngest ensures the sources and runs one ingestor per source. The
// returned stop function cancels the readers and joins them inside the
// budget before their file handles close.
func startIngest(st *store.Store, list files, fromEnd bool) func() {
	ctx, cancel := context.WithCancel(context.Background())

	srcs, err := ensureSources(ctx, st, list, fromEnd)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	var (
		wg  sync.WaitGroup
		ing []*store.Ingestor
	)

	for _, src := range srcs {
		in, err := st.Ingest(ctx, src.ID, entry.DefaultPolicy())
		if err != nil {
			cancel()
			log.Fatal(err)
		}

		ing = append(ing, in)
		wg.Add(1)

		go func(in *store.Ingestor) {
			defer wg.Done()

			if err := in.Run(ctx); err != nil && ctx.Err() == nil {
				log.Printf("live-log-viewer: source %d: %v", in.Source().ID, err)
			}
		}(in)
	}

	return func() {
		cancel()
		done := make(chan struct{})

		go func() { wg.Wait(); close(done) }()

		select {
		case <-done:
		case <-time.After(stopBudget):
			log.Printf("live-log-viewer: ingestors did not stop within %s", stopBudget)
		}

		for _, in := range ing {
			_ = in.Close()
		}
	}
}
