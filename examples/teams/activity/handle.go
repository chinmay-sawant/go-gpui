package activity

// open activates one item and clears its unread dot.
func (d *Data) open(id string) {
	d.Active = id

	for i := range d.Items {
		if d.Items[i].ID == id {
			d.Items[i].Unread = false
		}
	}
}

// reply marks the active item replied and adds one reply.
func (d *Data) reply() {
	for i := range d.Items {
		if d.Items[i].ID == d.Active {
			d.Items[i].Replied = true
			d.Items[i].Replies++
		}
	}
}

// markAllRead clears every unread dot.
func (d *Data) markAllRead() {
	for i := range d.Items {
		d.Items[i].Unread = false
	}
}
