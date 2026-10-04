package drop_test

import (
	"io/fs"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

// dropsFrom builds gpui.Drop values from a fake file system the way the
// window does.
func dropsFrom(t *testing.T, fsys fs.FS) []gpui.Drop {
	t.Helper()

	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		t.Fatal(err)
	}

	drops := make([]gpui.Drop, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			t.Fatal(err)
		}

		name := entry.Name()

		drops = append(drops, gpui.Drop{
			Name:  name,
			Size:  info.Size(),
			IsDir: info.IsDir(),
			Read: func() ([]byte, error) {
				return fs.ReadFile(fsys, name)
			},
		})
	}

	return drops
}
