package ui

import (
	"fmt"
	"sort"
)

func (b *fakeBackend) ExportCSV(id string) (string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	keys := make([][2]int, 0, len(b.cells[id]))
	for k := range b.cells[id] {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i][0] != keys[j][0] {
			return keys[i][0] < keys[j][0]
		}

		return keys[i][1] < keys[j][1]
	})
	b.export = fmt.Sprint(keys)

	return b.export, nil
}

func (b *fakeBackend) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.closed = true

	return nil
}
