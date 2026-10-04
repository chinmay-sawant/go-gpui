// Package dictionary is the dictionary page: custom words and their sounds.
package dictionary

// WordRow is one row in the dictionary table.
type WordRow struct {
	ID     string
	Word   string
	Sounds string
	Added  string
}

// Data is the data the dictionary page prints.
type Data struct {
	Words []WordRow
}

// Default returns the dictionary page data.
func Default() Data {
	return Data{
		Words: []WordRow{
			{"1", "Kubernetes", "koo-ber-net-eez", "Aug 14"},
			{"2", "Chinmay", "chin-my", "Aug 02"},
			{"3", "GowkHTML", "gawk-html", "Jul 28"},
			{"4", "DeepSeek", "deep-seek", "Jul 21"},
			{"5", "Ebiten", "eh-bih-ten", "Jul 12"},
			{"6", "gowkhtmltopdf", "gawk-html-to-pdf", "Jun 30"},
		},
	}
}
