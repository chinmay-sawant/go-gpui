package chat

// AllThreads returns a copy of every thread for the database.
func (d Data) AllThreads() map[string][]Message {
	out := make(map[string][]Message, len(d.threads))

	for id, msgs := range d.threads {
		out[id] = append([]Message(nil), msgs...)
	}

	return out
}
