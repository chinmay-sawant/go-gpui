package store

import (
	"context"
	"os"
	"path/filepath"
)

// SaveExport exports and writes the bytes to path atomically: a temporary
// file in the same directory is renamed over the target, so a failure
// midway leaves no partial export behind.
func (s *Store) SaveExport(ctx context.Context, o ExportOptions, path string) error {
	ex, err := s.Export(ctx, o)
	if err != nil {
		return err
	}

	tmp, err := os.CreateTemp(filepath.Dir(path), ".live-log-export-*")
	if err != nil {
		return err
	}

	name := tmp.Name()

	if _, err := tmp.Write(ex.Data); err != nil {
		_ = tmp.Close()
		_ = os.Remove(name)

		return err
	}

	if err := tmp.Close(); err != nil {
		_ = os.Remove(name)

		return err
	}

	if err := os.Rename(name, path); err != nil {
		_ = os.Remove(name)

		return err
	}

	return nil
}
