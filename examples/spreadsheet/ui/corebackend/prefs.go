package corebackend

import "context"

// Pref reads a stored preference.
func (b *Backend) Pref(key string) (string, bool) {
	v, ok, err := b.st.GetPref(context.Background(), key)
	if err != nil {
		return "", false
	}

	return v, ok
}

// SetPref stores a preference.
func (b *Backend) SetPref(key, value string) error {
	return b.st.SetPref(context.Background(), key, value)
}

// Close stops the store worker and closes the database.
func (b *Backend) Close() error {
	if b.st == nil {
		return nil
	}

	return b.st.Close()
}
