package page

import "strings"

func findHead(source string) int {
	const end = "</head>"
	for i := 0; i+len(end) <= len(source); i++ {
		if strings.EqualFold(source[i:i+len(end)], end) {
			return i
		}
	}

	return -1
}

func spanCovers(spans []controlSpan, index int) bool {
	for _, sp := range spans {
		if index >= sp.Start && index < sp.End {
			return true
		}
	}

	return false
}

func pickControl(sp controlSpan, live map[string]Control) Control {
	ctrl := sp.Control
	got, ok := live[ctrl.ID]
	if !ok {
		return ctrl
	}

	ctrl.Value = got.Value
	ctrl.Checked = got.Checked
	ctrl.Disabled = got.Disabled
	ctrl.Options = got.Options
	ctrl.Type = got.Type
	ctrl.Name = got.Name
	ctrl.Tag = got.Tag

	return ctrl
}
