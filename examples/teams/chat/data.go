package chat

// FromDB builds the chat state the store loaded.
func FromDB(chats []Chat, threads map[string][]Message, filter, query, active string) Data {
	d := Data{Filter: filter, Query: query, all: cloneChats(chats), threads: threads}
	d.rebuild()
	d.Active = active

	if d.visible() {
		open(&d, active)
	} else {
		d.clearThread()
	}

	return d
}
