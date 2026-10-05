package chat

// open shows one chat, clears its unread count, and loads its thread.
func open(d *Data, id string) {
	for i := range d.Chats {
		if d.Chats[i].ID != id {
			continue
		}

		c := &d.Chats[i]
		d.Active = id
		d.Header = c.Name
		d.HeaderInitials = c.Initials
		d.HeaderColor = c.Color
		d.HeaderStatus = chatStatus[id]
		d.Messages = cloneThread(id)
		c.Unread = 0

		return
	}
}

// toggle flips the pinned or muted flag of the open chat.
func toggle(d *Data, pinned bool) {
	for i := range d.Chats {
		if d.Chats[i].ID != d.Active {
			continue
		}

		if pinned {
			d.Chats[i].Pinned = !d.Chats[i].Pinned
		} else {
			d.Chats[i].Muted = !d.Chats[i].Muted
		}

		return
	}
}

// applyFilter switches the list filter and closes a thread that no longer
// passes it.
func applyFilter(d *Data, filter string) {
	d.Filter = filter
	d.rebuild()

	if !d.visible() {
		d.clearThread()
	}
}

// clearThread closes the open conversation.
func (d *Data) clearThread() {
	d.Active = ""
	d.Header = ""
	d.HeaderInitials = ""
	d.HeaderColor = ""
	d.HeaderStatus = ""
	d.Messages = nil
}
