package player

// sampleTracks is the six-track sample tracklist.
func sampleTracks() []Track {
	return []Track{
		{Index: 0, Num: "1", Title: "Midnight Circuit", Artist: "Nova Waves", Album: "Neon Horizon", Year: "2025", Genre: "Electronic", Length: "3:32", Cover: "track-0", Active: true},
		{Index: 1, Num: "2", Title: "Paper Satellites", Artist: "The Far Coast", Album: "Static Bloom", Year: "2024", Genre: "Indie", Length: "4:05", Cover: "track-1"},
		{Index: 2, Num: "3", Title: "Golden Static", Artist: "Aster Field", Album: "Signal Fire", Year: "2023", Genre: "Ambient", Length: "2:58", Cover: "track-2"},
		{Index: 3, Num: "4", Title: "Velvet Machinery", Artist: "Mono Arcade", Album: "Soft Errors", Year: "2025", Genre: "Synthwave", Length: "3:47", Cover: "track-3"},
		{Index: 4, Num: "5", Title: "Low Orbit Lullaby", Artist: "Night Cartography", Album: "Blue Hour", Year: "2022", Genre: "Lo-Fi", Length: "4:21", Cover: "track-4"},
		{Index: 5, Num: "6", Title: "Afterglow Avenue", Artist: "Nova Waves", Album: "Neon Horizon", Year: "2025", Genre: "Electronic", Length: "3:12", Cover: "track-5"},
	}
}

// samplePicks is the six-tile Good evening grid.
func samplePicks() []Card {
	return []Card{
		{Index: 0, Title: "Daily Mix 1", Sub: "Nova Waves & more", Cover: "pick-0"},
		{Index: 1, Title: "Discover Weekly", Sub: "The Far Coast & more", Cover: "pick-1"},
		{Index: 2, Title: "Release Radar", Sub: "Aster Field & more", Cover: "pick-2"},
		{Index: 3, Title: "Deep Focus", Sub: "Mono Arcade & more", Cover: "pick-3"},
		{Index: 4, Title: "Lo-Fi Beats", Sub: "Night Cartography & more", Cover: "pick-4"},
		{Index: 5, Title: "On Repeat", Sub: "Echo Valley & more", Cover: "pick-5"},
	}
}
