// Package transforms is the transforms page: one-command rewrites.
package transforms

// Data is the data the transforms page prints.
type Data struct {
	Rows []TransformRow
}

// TransformRow is one saved transform. N is the one-based number the Run
// button id uses, Desc is the muted line under the name, and Prompt is the
// instruction Flow sends.
type TransformRow struct {
	N      int
	Name   string
	Desc   string
	Prompt string
}

// Default returns the transforms page data.
func Default() Data {
	return Data{Rows: []TransformRow{
		{1, "Make it shorter", "Trim the dictation to the essential sentence.",
			"Rewrite the text below in as few words as possible without losing the point."},
		{2, "Fix grammar", "Correct punctuation and spelling but keep my wording.",
			"Fix spelling, punctuation, and grammar in the text below. Change nothing else."},
		{3, "Summarize", "Condense into three bullets.",
			"Summarize the text below in three short bullet points."},
		{4, "Turn into bullets", "Split the text into bullet points.",
			"Split the text below into bullet points, one idea per line."},
		{5, "Reply as email", "Rewrite as a short professional email.",
			"Draft a short, professional email reply from the text below."},
	}}
}
