package workbook

import (
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// ID returns the stable workbook identifier.
func (w *Workbook) ID() ID { return w.id }

// SetID gives the workbook its storage identifier.
func (w *Workbook) SetID(id ID) { w.id = id }

// Name returns the workbook name.
func (w *Workbook) Name() string { return w.name }

// SetName renames the workbook.
func (w *Workbook) SetName(name string) { w.name = name }

// Rev returns the content revision; every edit advances it.
func (w *Workbook) Rev() int64 { return w.rev }

// SetRev sets the revision, which storage does on load.
func (w *Workbook) SetRev(rev int64) { w.rev = rev }

// Limits returns the parse and evaluation limits.
func (w *Workbook) Limits() formula.Limits { return w.lim }

// SetLimits replaces the limits for future parses.
func (w *Workbook) SetLimits(lim formula.Limits) { w.lim = lim.Normalized() }

// Sheets returns the sheets in order.
func (w *Workbook) Sheets() []*Sheet {
	return append([]*Sheet(nil), w.sheets...)
}

// Sheet returns the sheet with id, or nil.
func (w *Workbook) Sheet(id SheetID) *Sheet { return w.bySheet[id] }
