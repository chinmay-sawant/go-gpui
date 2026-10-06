// Package workbook is the spreadsheet example's in-memory model: sparse
// sheets, typed cells, edits with bounded undo, formula recalculation, a
// bounded pending-save queue, and CSV import and export. It never touches
// disk or a window; storage persists it and the ui package drives it.
package workbook

import (
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// ID is a stable workbook identifier.
type ID int64

// SheetID is a stable sheet identifier.
type SheetID int64

// ErrNoSheet reports a command for a sheet the workbook does not have.
var ErrNoSheet = errors.New("workbook: no such sheet")

// DefaultHistory is the number of edit commands undo keeps.
const DefaultHistory = 200

// Workbook is one document: stable IDs, one or more sheets, and a content
// revision that advances with every edit.
type Workbook struct {
	id      ID
	name    string
	rev     int64
	saved   int64
	sheets  []*Sheet
	bySheet map[SheetID]*Sheet
	history []histEntry
	future  []histEntry
	limit   int
	lim     formula.Limits
}

// New builds an empty workbook. Storage assigns the ID and sheet IDs.
func New(id ID, name string) *Workbook {
	return &Workbook{
		id:      id,
		name:    name,
		bySheet: map[SheetID]*Sheet{},
		limit:   DefaultHistory,
		lim:     formula.DefaultLimits(),
	}
}
