package store

// Load reads the whole state back. found is false when the database has
// not been seeded yet.
func (s *Store) Load() (Data, bool, error) {
	empty, err := s.Empty()
	if err != nil {
		return Data{}, false, err
	}

	if empty {
		return Data{}, false, nil
	}

	var d Data

	if d.Activity, err = loadActivity(s.db); err != nil {
		return Data{}, false, err
	}

	if d.Chat, err = loadChat(s.db); err != nil {
		return Data{}, false, err
	}

	if d.Channels, err = loadChannels(s.db); err != nil {
		return Data{}, false, err
	}

	if d.Calendar, err = loadCalendar(s.db); err != nil {
		return Data{}, false, err
	}

	if d.Calls, err = loadCalls(s.db); err != nil {
		return Data{}, false, err
	}

	if d.Files, err = loadFiles(s.db); err != nil {
		return Data{}, false, err
	}

	if d.Presence, d.Dark, err = loadProfile(s.db); err != nil {
		return Data{}, false, err
	}

	return d, true, nil
}
