// Package chat is the Chat menu of the Teams example.
package chat

// Data is the chat state the shell prints.
type Data struct {
	Filter         string
	Query          string
	Active         string
	Header         string
	HeaderInitials string
	HeaderColor    string
	HeaderStatus   string
	Chats          []Chat
	Messages       []Message

	// threads is every chat's full message history, keyed by chat id.
	threads map[string][]Message

	// all is every chat; Chats is the filter and search view of it.
	all []Chat
}

// Chat is one conversation in the chat list.
type Chat struct {
	ID, Name, Initials, Color, Preview, Time string
	Unread                                   int
	Pinned, Muted, Group                     bool
}

// Presence returns the presence dot class for a direct chat, or "" for a group.
func (c Chat) Presence() string {
	if c.Group {
		return ""
	}

	return chatPresence[c.ID]
}

// Message is one message in the open thread.
type Message struct {
	ID, Author, Initials, Color, Time, Text string
	Own                                     bool
}

// AllChats returns every chat in list order with the visible list's changes.
func (d Data) AllChats() []Chat {
	out := cloneChats(d.all)

	for i := range out {
		for _, c := range d.Chats {
			if c.ID == out[i].ID {
				out[i] = c

				break
			}
		}
	}

	return out
}

// cloneChats copies a chat list so two callers never share state.
func cloneChats(src []Chat) []Chat {
	out := make([]Chat, len(src))
	copy(out, src)

	return out
}
