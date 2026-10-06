package telegram

// Chat is one row of the chat list and the open thread's header.
type Chat struct {
	ID       string
	Name     string
	Initials string
	Color    int
	Preview  string
	Time     string
	Unread   int
	Pinned   bool
	Muted    bool
	Group    bool
}

// Message is one bubble in an open conversation.
type Message struct {
	ID       string
	Text     string
	Time     string
	Own      bool
	Read     bool
	Author   string
	Initials string
	Color    int
}

// Contact is one row of the contacts tab.
type Contact struct {
	ID       string
	Name     string
	Initials string
	Color    int
	Status   string
	ChatID   string
}

// View is the data the template prints.
type View struct {
	Tab      string // "chats" | "contacts" | "settings"
	Dark     bool
	Query    string
	Draft    string
	Active   string
	Header   string
	Status   string
	Initials string
	Color    int
	Group    bool
	Chats    []Chat
	Contacts []Contact
	Thread   []Message
	Unread   int
}
