package calls

import "strings"

// rebuild refills both lists from the search box text.
func (d *Data) rebuild() {
	query := strings.ToLower(strings.TrimSpace(d.Query))

	d.History = nil

	for _, c := range d.allHistory {
		if query == "" || strings.Contains(strings.ToLower(c.Name), query) {
			d.History = append(d.History, c)
		}
	}

	d.Voicemails = nil

	for _, v := range d.allVoicemails {
		if query == "" || strings.Contains(strings.ToLower(v.Name), query) {
			d.Voicemails = append(d.Voicemails, v)
		}
	}
}
