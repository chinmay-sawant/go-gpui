package player

// sampleStatus is the status line while the sample data is showing.
const sampleStatus = "Sample data (offline)"

// DefaultView returns the sample data the player draws offline.
func DefaultView() View {
	tracks := sampleTracks()
	now := tracks[0]

	return View{
		Nav:        "home",
		Liked:      true,
		Volume:     70,
		Progress:   32,
		Elapsed:    "1:08",
		Remaining:  "2:24",
		Now:        now,
		Tracks:     tracks,
		Picks:      samplePicks(),
		Shelf:      sampleShelf(),
		Playlists:  samplePlaylists(),
		Status:     sampleStatus,
		ShelfTitle: "Recently played",
	}
}
