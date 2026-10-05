package channels

// defaultFiles returns the five sample files of the active channel.
func defaultFiles() []FileItem {
	return []FileItem{
		{ID: "f1", Name: "Mark LXXXV schematic.pdf", Badge: "pdf", Kind: "PDF", Modified: "Today, 9:12 AM", ModifiedBy: "Robert Downey Jr.", Size: "8.2 MB"},
		{ID: "f2", Name: "Vibranium yield report.xlsx", Badge: "xls", Kind: "X", Modified: "Yesterday", ModifiedBy: "Shuri", Size: "1.1 MB"},
		{ID: "f3", Name: "Mission log Zurich.docx", Badge: "doc", Kind: "W", Modified: "Yesterday", ModifiedBy: "Natasha Romanoff", Size: "640 KB", Starred: true},
		{ID: "f4", Name: "Arc reactor diagnostics.pptx", Badge: "ppt", Kind: "P", Modified: "Oct 1", ModifiedBy: "Bruce Banner", Size: "12 MB"},
		{ID: "f5", Name: "Suit lab test data.zip", Badge: "zip", Kind: "ZIP", Modified: "Sep 28", ModifiedBy: "Peter Parker", Size: "24 MB"},
	}
}
