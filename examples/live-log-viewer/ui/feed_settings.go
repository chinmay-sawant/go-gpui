package ui

import (
	"context"
	"encoding/json"
	"errors"
	"os"
)

// Settings loads the UI preferences from the JSON file beside the database.
// A missing or unreadable file falls back to the defaults: follow on.
func (f *StoreFeed) Settings(context.Context) (Settings, error) {
	def := Settings{Follow: true}

	data, err := os.ReadFile(f.settings)
	if errors.Is(err, os.ErrNotExist) {
		return def, nil
	}

	if err != nil {
		return def, nil
	}

	var s Settings
	if err := json.Unmarshal(data, &s); err != nil {
		return def, nil
	}

	return s, nil
}

// SaveSettings writes the UI preferences atomically.
func (f *StoreFeed) SaveSettings(_ context.Context, s Settings) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return writeFile(f.settings, data)
}
