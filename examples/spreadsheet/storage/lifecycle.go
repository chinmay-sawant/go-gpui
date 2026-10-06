package storage

// Close stops the worker and closes the database. It is safe twice, and it
// never closes the request channel under a concurrent sender.
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}

	if s.closed.Swap(true) {
		return nil
	}

	close(s.quit)
	<-s.done

	return s.db.Close()
}

// Path is the database file path, or ":memory:".
func (s *Store) Path() string { return s.path }

// JournalMode is the journal mode SQLite reported at open. A build that
// cannot enable WAL keeps the reported rollback mode and keeps working.
func (s *Store) JournalMode() string {
	s.journalMu.Lock()
	defer s.journalMu.Unlock()

	return s.journal
}

func (s *Store) setJournal(mode string) {
	s.journalMu.Lock()
	defer s.journalMu.Unlock()

	s.journal = mode
}
