package ui

import "strings"

func (b *fakeBackend) Undo() (UndoResult, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.undos) == 0 {
		return UndoResult{Rev: b.rev, OK: false, Label: "undo"}, nil
	}

	last := b.undos[len(b.undos)-1]
	b.undos = b.undos[:len(b.undos)-1]
	b.redos = append(b.redos, last)
	b.rev++

	return UndoResult{Rev: b.rev, OK: true, Label: "undo"}, nil
}

func (b *fakeBackend) Redo() (UndoResult, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.redos) == 0 {
		return UndoResult{Rev: b.rev, OK: false, Label: "redo"}, nil
	}

	last := b.redos[len(b.redos)-1]
	b.redos = b.redos[:len(b.redos)-1]
	b.undos = append(b.undos, last)
	b.rev++

	return UndoResult{Rev: b.rev, OK: true, Label: "redo"}, nil
}

func (b *fakeBackend) Pref(key string) (string, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	v, ok := b.prefs[key]

	return v, ok
}

func (b *fakeBackend) SetPref(key, value string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.prefs[key] = value

	return nil
}

func (b *fakeBackend) PreviewCSV(id string, data []byte, replace bool) (Preview, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	rows := make([][]string, 0, len(lines))
	cols := 0
	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) > cols {
			cols = len(fields)
		}

		rows = append(rows, fields)
	}

	return Preview{Rows: rows, Total: len(rows), Cols: cols}, nil
}

func (b *fakeBackend) CommitCSV(id string, data []byte, replace bool) (ImportResult, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.rev++
	b.csv = [][]string{{string(data)}}

	return ImportResult{Rev: b.rev, Rows: 1, Cols: 1}, nil
}
