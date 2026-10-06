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
	Gift     bool
	Photo    string
	PhotoW   int
	PhotoH   int
	Reaction string
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
	// InsetTop and InsetBottom pad the page for the phone's system bars.
	InsetTop    int
	InsetBottom int
	// BarTop and BottomTop pin the thread bars to the viewport as the
	// page scrolls.
	BarTop    int
	BottomTop int
	// AttachOpen draws the attachment sheet above the composer in a thread.
	AttachOpen bool
	// CanGift allows the gift button in one chat only.
	CanGift bool
	// ReactID is the message whose long-press reaction bar is open.
	ReactID string
	// Reactions is the emoji bar shown while ReactID is set.
	Reactions []Reaction
}
