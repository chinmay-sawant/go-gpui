package player

const (
	defaultBase   = "https://itunes.apple.com"
	defaultTerm   = "daft punk"
	offlineStatus = "Sample data (offline)"
	searchLimit   = 6
	maxCovers     = 12
)

// itunesSong is the field set the player reads from one search result.
type itunesSong struct {
	Kind             string `json:"kind"`
	TrackName        string `json:"trackName"`
	ArtistName       string `json:"artistName"`
	CollectionName   string `json:"collectionName"`
	PrimaryGenreName string `json:"primaryGenreName"`
	TrackTimeMillis  int    `json:"trackTimeMillis"`
	ArtworkURL100    string `json:"artworkUrl100"`
}

// itunesResponse is the search envelope.
type itunesResponse struct {
	Results []itunesSong `json:"results"`
}
