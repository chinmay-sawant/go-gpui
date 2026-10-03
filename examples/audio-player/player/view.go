package player

// Track is one song in the queue or the recent list. Index is its position
// in the queue, so a recent tile can select the same queue entry.
type Track struct {
	Index         int
	Title, Artist string
	Album, Genre  string
	Length        string
	Cover         string
	Active        bool
}

// Playlist is one row in the sidebar.
type Playlist struct {
	Index  int
	Name   string
	Count  string
	Active bool
	Tone   string
}

// View is the data the player template prints.
type View struct {
	Nav        string // "home" | "search" | "library"
	Playing    bool
	Liked      bool
	Shuffle    bool
	Repeat     bool
	Muted      bool
	Volume     int // 0..100
	Progress   int // 0..100 of the current track
	Elapsed    string
	Remaining  string
	Now        Track
	Queue      []Track
	Recent     []Track
	Playlists  []Playlist
	Status     string
	Query      string
	QueueCount string
}
