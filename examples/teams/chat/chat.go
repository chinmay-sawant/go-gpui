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

	// all is every sample chat; Chats is the filter and search view of it.
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
