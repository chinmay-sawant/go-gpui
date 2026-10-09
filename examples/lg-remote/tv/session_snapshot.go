package tv

// publishStore runs under the connection lock, except during construction.
// UI readers use the immutable snapshot without waiting for network I/O.
func (s *Session) publishStore() {
	copy := s.store
	s.saved.Store(&copy)
}

func (s *Session) savedStore() Store {
	if saved := s.saved.Load(); saved != nil {
		return *saved
	}
	return Store{}
}
