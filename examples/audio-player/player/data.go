package player

// DefaultView returns the sample library the player draws offline.
func DefaultView() View {
	queue := sampleQueue()
	v := View{
		Nav:        "home",
		Playing:    true,
		Volume:     72,
		Now:        queue[0],
		Queue:      queue,
		Recent:     recentFrom(queue),
		Playlists:  samplePlaylists(),
		Status:     offlineStatus,
		QueueCount: countLabel(len(queue)),
		Credit:     "Free music via Openverse",
	}
	v.setProgress(38)

	return v
}

// sampleQueue is the six-track fictional library.
func sampleQueue() []Track {
	return []Track{
		{Index: 0, Title: "Midnight Drive", Artist: "Nova Waves", Album: "Neon Horizon",
			Genre: "Synthwave", Length: "3:42", Cover: "cover-1", Active: true},
		{Index: 1, Title: "Paper Lanterns", Artist: "Iris Vale", Album: "Slow Light",
			Genre: "Indie Pop", Length: "4:05", Cover: "cover-2"},
		{Index: 2, Title: "Velvet Static", Artist: "The February Club", Album: "Analog Heart",
			Genre: "Alt Rock", Length: "3:18", Cover: "cover-3"},
		{Index: 3, Title: "Blue Hour", Artist: "Sable & Co.", Album: "City Weather",
			Genre: "Jazz", Length: "5:12", Cover: "cover-4"},
		{Index: 4, Title: "Golden Thread", Artist: "Mira Sol", Album: "Weave",
			Genre: "Folk", Length: "3:51", Cover: "cover-1"},
		{Index: 5, Title: "Echo Chamber", Artist: "Kite Machine", Album: "Signal Lost",
			Genre: "Electronic", Length: "4:27", Cover: "cover-2"},
	}
}

// samplePlaylists is the six-row playlist column.
func samplePlaylists() []Playlist {
	return []Playlist{
		{Index: 0, Name: "Late Night", Count: "24 songs", Active: true, Tone: "tone-violet"},
		{Index: 1, Name: "Focus Flow", Count: "41 songs", Tone: "tone-indigo"},
		{Index: 2, Name: "Morning Run", Count: "18 songs", Tone: "tone-amber"},
		{Index: 3, Name: "Vinyl Crates", Count: "57 songs", Tone: "tone-green"},
		{Index: 4, Name: "Soft Static", Count: "32 songs", Tone: "tone-rose"},
		{Index: 5, Name: "Deep Cuts", Count: "76 songs", Tone: "tone-sky"},
	}
}
