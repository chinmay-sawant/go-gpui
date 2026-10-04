package window

import (
	"io/fs"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// dropPass offers the files dropped during this frame to a screen that takes
// drops. A nil file system means no drop, which is what DroppedFiles returns
// outside a real frame and on a frame with no drop.
func (s *shell) dropPass(fsys fs.FS) error {
	if fsys == nil {
		return nil
	}

	files, err := dropFiles(fsys)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	screen, ok := s.app.(host.Dropper)
	if !ok {
		return nil
	}

	return screen.Drop(s.ctx, files)
}

// dropFiles describes every root entry of fsys as one host.Drop. A directory
// stays one entry; its tree is not walked.
func dropFiles(fsys fs.FS) ([]host.Drop, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}

	files := make([]host.Drop, 0, len(entries))
	for _, entry := range entries {
		d, err := dropValue(fsys, entry)
		if err != nil {
			return nil, err
		}

		files = append(files, d)
	}

	return files, nil
}

// dropValue describes one entry. Read reads through fsys, so it works only
// while the frame that owns fsys is alive.
func dropValue(fsys fs.FS, entry fs.DirEntry) (host.Drop, error) {
	info, err := entry.Info()
	if err != nil {
		return host.Drop{}, err
	}

	name := entry.Name()

	d := host.Drop{
		Name:  name,
		Size:  info.Size(),
		IsDir: info.IsDir(),
		Read: func() ([]byte, error) {
			return fs.ReadFile(fsys, name)
		},
	}

	if p, ok := entry.(interface{ AbsPath() string }); ok {
		d.Path = p.AbsPath()
	}

	return d, nil
}
