package window

import (
	"context"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestDropPassSkipsPlainScreen(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}

	if err := s.dropPass(fstest.MapFS{"a.txt": {Data: []byte("x")}}); err != nil {
		t.Fatalf("dropPass: %v", err)
	}
}

func TestDropPassEmptyFSCallsNothing(t *testing.T) {
	t.Parallel()

	screen := &dropScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: screen, ctx: context.Background()}

	if err := s.dropPass(fstest.MapFS{}); err != nil {
		t.Fatalf("dropPass: %v", err)
	}

	if screen.calls != 0 {
		t.Fatalf("calls = %d, want 0", screen.calls)
	}
}

func TestDropPassReturnsScreenError(t *testing.T) {
	t.Parallel()

	want := errors.New("drop failed")
	screen := &dropScreen{fakeScreen: &fakeScreen{}, err: want}
	s := &shell{app: screen, ctx: context.Background()}

	if err := s.dropPass(fstest.MapFS{"a.txt": {Data: []byte("x")}}); !errors.Is(err, want) {
		t.Fatalf("err = %v", err)
	}
}

// absFS adds AbsPath to the entries of a MapFS, the way Ebiten's desktop
// dropped-files system does.
type absFS struct{ fstest.MapFS }

func (f absFS) ReadDir(name string) ([]fs.DirEntry, error) {
	entries, err := f.MapFS.ReadDir(name)
	if err != nil {
		return nil, err
	}

	for i, entry := range entries {
		entries[i] = absEntry{entry}
	}

	return entries, nil
}

type absEntry struct{ fs.DirEntry }

func (e absEntry) AbsPath() string { return "/real/" + e.Name() }

func TestDropPassCarriesAbsPath(t *testing.T) {
	t.Parallel()

	screen := &dropScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: screen, ctx: context.Background()}

	fsys := absFS{MapFS: fstest.MapFS{"a.txt": {Data: []byte("x")}}}

	if err := s.dropPass(fsys); err != nil {
		t.Fatalf("dropPass: %v", err)
	}

	if got := screen.files[0].Path; got != "/real/a.txt" {
		t.Fatalf("path = %q", got)
	}
}
