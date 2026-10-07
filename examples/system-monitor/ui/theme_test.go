package ui

import (
	"context"
	"testing"
)

func TestThemeTogglePersists(t *testing.T) {
	store := &fakeStore{}
	app := newTestApp(t, &fakeSource{}, store)

	if err := app.handleClick(context.Background(), actTheme); err != nil {
		t.Fatal(err)
	}

	if !app.state.dark {
		t.Fatal("theme did not flip to dark")
	}

	if store.saved.Dark != true || store.saves != 1 {
		t.Fatalf("theme not persisted: %+v saves=%d", store.saved, store.saves)
	}
}

func TestStoredThemeLoads(t *testing.T) {
	store := &fakeStore{saved: Settings{Dark: true}}
	app := newTestApp(t, &fakeSource{}, store)

	if !app.state.dark {
		t.Fatal("stored dark theme was not loaded")
	}
}
