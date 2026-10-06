package main

import (
	"errors"
	"log"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// openStore opens the database, falling back to a temporary directory when
// the configured one is not writable.
func openStore(dir string, temp bool) *store.Store {
	if temp {
		return tempStore()
	}

	st, err := store.Open(dir)
	if err == nil {
		return st
	}

	if errors.Is(err, store.ErrReadOnly) {
		log.Printf("live-log-viewer: %v; using a temporary database instead", err)

		return tempStore()
	}

	log.Fatal(err)

	return nil
}

// tempStore opens a throwaway database that Close removes.
func tempStore() *store.Store {
	st, err := store.OpenWithOptions(store.Options{Temp: true})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("live-log-viewer: temporary database %s", st.Path())

	return st
}
