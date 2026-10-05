package chat

import "strings"

// rebuild refills the visible list from the filter chip and the search box.
func (d *Data) rebuild() {
	query := strings.ToLower(strings.TrimSpace(d.Query))
	d.Chats = nil

	for _, c := range d.all {
		if d.match(c, query) {
			d.Chats = append(d.Chats, c)
		}
	}
}

// match reports whether one chat passes the filter and the query.
func (d Data) match(c Chat, query string) bool {
	switch d.Filter {
	case "unread":
		if c.Unread == 0 {
			return false
		}
	case "groups":
		if !c.Group {
			return false
		}
	}

	if query == "" {
		return true
	}

	name := strings.ToLower(c.Name)
	preview := strings.ToLower(c.Preview)

	return strings.Contains(name, query) || strings.Contains(preview, query)
}

// visible reports whether the open chat is still in the list.
func (d Data) visible() bool {
	for _, c := range d.Chats {
		if c.ID == d.Active {
			return true
		}
	}

	return false
}
