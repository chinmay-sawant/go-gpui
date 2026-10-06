package store

import "os"

// Close stops the worker and closes the connection. A temporary store also
// removes its directory.
func (s *Store) Close() error {
	if !s.closed.CompareAndSwap(false, true) {
		return nil
	}

	close(s.stop)
	s.wg.Wait()

	err := s.db.Close()
	if s.temp {
		_ = os.RemoveAll(s.dir)
	}

	return err
}

// Dir returns the data directory. Path returns the database file, or Memory.
func (s *Store) Dir() string  { return s.dir }
func (s *Store) Path() string { return s.path }

// Temporary reports whether Close removes the data directory.
func (s *Store) Temporary() bool { return s.temp }
