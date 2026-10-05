// Package files is the Files menu of the Teams example: a file list with
// Recent, Teams, Shared, and Starred filters, a Name table, a detail
// column, and star toggles.
package files

// Data is the files state the shell prints.
type Data struct {
	Filter string // "recent" | "teams" | "shared" | "starred"
	Query  string
	Active string
	Files  []File

	all []File // every file; Files is the filter's view of it
}

// File is one file entry.
type File struct {
	ID, Name                                string
	Badge                                   string
	BadgeText                               string
	Kind, Modified, ModifiedBy, Size, Where string
	Team, Shared, Starred                   bool
}

// FromDB rebuilds the files state from stored rows.
func FromDB(all []File, filter, active string) Data {
	d := Data{Filter: filter, Active: active, all: all}
	d.rebuild()

	return d
}

// AllFiles returns every file row, before the filter.
func (d Data) AllFiles() []File { return d.all }
