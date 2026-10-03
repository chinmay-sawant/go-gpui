package player

// sampleShelf is the six-card Recently played shelf.
func sampleShelf() []Card {
	return []Card{
		{Index: 0, Title: "Neon Horizon", Sub: "Nova Waves", Cover: "card-0"},
		{Index: 1, Title: "Static Bloom", Sub: "The Far Coast", Cover: "card-1"},
		{Index: 2, Title: "Signal Fire", Sub: "Aster Field", Cover: "card-2"},
		{Index: 3, Title: "Soft Errors", Sub: "Mono Arcade", Cover: "card-3"},
		{Index: 4, Title: "Blue Hour", Sub: "Night Cartography", Cover: "card-4"},
		{Index: 5, Title: "Paper Trails", Sub: "Echo Valley", Cover: "card-5"},
	}
}

// samplePlaylists is the five-row sidebar library.
func samplePlaylists() []Playlist {
	return []Playlist{
		{Index: 0, Name: "Liked Songs", Meta: "Playlist · 214 songs", Cover: "card-0"},
		{Index: 1, Name: "Neon Nights", Meta: "Playlist · Chinmay", Cover: "card-1"},
		{Index: 2, Name: "Focus Flow", Meta: "Playlist · Chinmay", Cover: "card-2"},
		{Index: 3, Name: "Late Drive", Meta: "Playlist · Spotify", Cover: "card-3"},
		{Index: 4, Name: "Rainy Day", Meta: "Playlist · Spotify", Cover: "card-4"},
	}
}
