package ui

import (
	"sync"
)

// fakeBackend is an in-memory Backend for headless tests.
type fakeBackend struct {
	mu      sync.Mutex
	sheets  []Sheet
	cells   map[string]map[[2]int]Cell
	rev     uint64
	prefs   map[string]string
	applied [][]Edit
	fetches []Area
	failSet bool
	closed  bool
	undos   []map[[2]int]Cell
	redos   []map[[2]int]Cell
	csv     [][]string
	export  string
}

func newFake() *fakeBackend {
	b := &fakeBackend{
		sheets: []Sheet{
			{ID: "s1", Name: "Sheet1", Rows: 200, Cols: 20},
			{ID: "s2", Name: "Numbers", Rows: 120, Cols: 12},
			{ID: "s3", Name: "Stress", Rows: 100000, Cols: 20},
		},
		cells: map[string]map[[2]int]Cell{},
		prefs: map[string]string{},
	}
	b.cells["s1"] = map[[2]int]Cell{
		{0, 0}:    {Raw: "name", Display: "name"},
		{0, 1}:    {Raw: "count", Display: "count"},
		{1, 0}:    {Raw: "alpha", Display: "alpha"},
		{1, 1}:    {Raw: "12", Display: "12", Num: true},
		{2, 0}:    {Raw: "=SUM(B2:B2)", Display: "12", Num: true},
		{2, 1}:    {Raw: "héllo", Display: "héllo"},
		{3, 0}:    {Raw: "=1/0", Display: "#DIV/0!", Err: "#DIV/0!"},
		{150, 15}: {Raw: "far", Display: "far"},
	}
	b.cells["s3"] = map[[2]int]Cell{
		{0, 0}:      {Raw: "stress", Display: "stress"},
		{99999, 19}: {Raw: "end", Display: "end"},
	}

	return b
}

func (b *fakeBackend) Sheets() []Sheet {
	b.mu.Lock()
	defer b.mu.Unlock()

	return append([]Sheet(nil), b.sheets...)
}

func (b *fakeBackend) Info(id string) (Sheet, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, sh := range b.sheets {
		if sh.ID == id {
			return sh, true
		}
	}

	return Sheet{}, false
}

func (b *fakeBackend) Revision() uint64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.rev
}
