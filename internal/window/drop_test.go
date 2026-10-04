package window

import (
	"context"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// dropScreen is a fakeScreen that accepts drops.
type dropScreen struct {
	*fakeScreen
	calls int
	files []host.Drop
	err   error
}

func (s *dropScreen) Drop(_ context.Context, files []host.Drop) error {
	s.calls++
	s.files = files

	return s.err
}

func TestDropPassCallsScreenOnce(t *testing.T) {
	t.Parallel()

	screen := &dropScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: screen, ctx: context.Background()}

	fsys := fstest.MapFS{
		"a.txt": {Data: []byte("hello")},
		"b.png": {Data: []byte("bytes")},
	}

	if err := s.dropPass(fsys); err != nil {
		t.Fatalf("dropPass: %v", err)
	}

	if screen.calls != 1 {
		t.Fatalf("calls = %d, want 1", screen.calls)
	}

	if len(screen.files) != 2 {
		t.Fatalf("files = %d, want 2", len(screen.files))
	}

	first := screen.files[0]
	if first.Name != "a.txt" || first.Size != 5 || first.IsDir {
		t.Fatalf("first = %+v", first)
	}

	data, err := first.Read()
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "hello" {
		t.Fatalf("data = %q", data)
	}
}

func TestDropPassKeepsDirectoryAsOneEntry(t *testing.T) {
	t.Parallel()

	screen := &dropScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: screen, ctx: context.Background()}

	fsys := fstest.MapFS{"docs": {Mode: fs.ModeDir | 0o755}}

	if err := s.dropPass(fsys); err != nil {
		t.Fatalf("dropPass: %v", err)
	}

	if len(screen.files) != 1 {
		t.Fatalf("files = %d, want 1", len(screen.files))
	}

	file := screen.files[0]
	if !file.IsDir || file.Name != "docs" {
		t.Fatalf("file = %+v", file)
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
