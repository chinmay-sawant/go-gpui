package ui

import "github.com/chinmay-sawant/ownframe"

// paintIDs lists every text run the tick rewrites in place.
func paintIDs() []string {
	ids := []string{"proc-new", "proc-at", "sel-name", "sel-status"}

	for _, p := range panelOrder {
		ids = append(ids, p+"-value", p+"-sub", p+"-peak")
	}

	for _, r := range detailRows {
		ids = append(ids, "df-"+r.key)
	}

	return ids
}

// boxByID finds one hit-test box by id.
func boxByID(boxes []ownframe.Box, id string) (ownframe.Box, bool) {
	for i := range boxes {
		if boxes[i].ID == id {
			return boxes[i], true
		}
	}

	return ownframe.Box{}, false
}
