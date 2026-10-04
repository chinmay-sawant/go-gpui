// Package snippets is the snippets page: short triggers that expand to full text.
package snippets

// Snippet is one trigger phrase and the text it expands into.
type Snippet struct {
	Trigger string
	Text    string
	Uses    string
	On      bool
}

// Data is the data the snippets page prints.
type Data struct {
	Rows []Snippet
}

// Default returns the snippets page data.
func Default() Data {
	return Data{Rows: []Snippet{
		{Trigger: ";addr", Uses: "18 times", On: true,
			Text: "12 Marine Drive, Bandra West, Mumbai 400050"},
		{Trigger: ";intro", Uses: "7 times", On: true,
			Text: "Hi, this is Chinmay. Thanks for the quick reply."},
		{Trigger: ";sig", Uses: "23 times", On: true,
			Text: "Best regards, Chinmay Sawant"},
		{Trigger: ";meet", Uses: "4 times", On: true,
			Text: "Can we move this to tomorrow morning?"},
	}}
}
