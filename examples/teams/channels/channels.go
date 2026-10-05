// Package channels is the Channels menu of the Teams example: the team list
// on the left and posts, files, and notes on the right.
package channels

// Data is the channels state the shell prints.
type Data struct {
	TeamName      string
	ChannelName   string
	ChannelDesc   string
	Tab           string
	ActiveTeam    string
	ActiveChannel string
	Teams         []Team
	Posts         []Post
	Files         []FileItem
	Emojis        []string

	// allPosts keeps every channel's posts, so switching channels does not
	// drop a change. FromDB points Posts at the active channel's entry.
	allPosts map[string][]Post
}

// AllPosts returns every channel's posts keyed by channel id.
func (d Data) AllPosts() map[string][]Post { return d.allPosts }

// Team is one team in the team list.
type Team struct {
	ID, Name, Initials, Color string
	Expanded                  bool
	Channels                  []Channel
}

// Channel is one channel under a team.
type Channel struct {
	ID, Name string
	Unread   int
}

// Post is one post on the posts tab.
type Post struct {
	ID, Author, Initials, Color, Time, Subject, Text string
	Likes                                            int
	Liked, Pinned, Expanded, Picker                  bool
	Replies                                          []Reply
	Reactions                                        []Reaction
}

// Reply is one reply in a post thread.
type Reply struct {
	ID, Author, Initials, Color, Time, Text string
	Own                                     bool
}

// Reaction is one emoji chip under a post.
type Reaction struct {
	Emoji string
	Count int
	Mine  bool
}

// FileItem is one file on the files tab.
type FileItem struct {
	ID, Name, Badge, Kind, Modified, ModifiedBy, Size string
	Starred                                           bool
}
