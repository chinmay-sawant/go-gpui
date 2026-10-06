package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// ensureSources returns the dummy fixture sources, or adds the -file paths
// to a file session. A source keeps its committed checkpoint across runs.
func ensureSources(ctx context.Context, st *store.Store, list files, fromEnd bool) ([]entry.Source, error) {
	if len(list) == 0 {
		setup, err := st.EnsureDummy(ctx, store.DummyOptions{})
		if err != nil {
			return nil, err
		}

		log.Printf("live-log-viewer: dummy mode, %d sources, database %s",
			len(setup.Sources), st.Path())

		return setup.Sources, nil
	}

	sess, err := st.EnsureSession(ctx, "Local logs", "file")
	if err != nil {
		return nil, err
	}

	out := make([]entry.Source, 0, len(list))

	for _, path := range list {
		src, err := st.AddSource(ctx, store.SourceSpec{
			Session: sess.ID, Path: path, Label: filepath.Base(path), FromEnd: fromEnd,
		})
		if err != nil {
			return nil, err
		}

		out = append(out, src)
	}

	log.Printf("live-log-viewer: following %d files, database %s", len(out), st.Path())

	return out, nil
}
