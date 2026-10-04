package player

import (
	"context"
	"strings"
)

// LibCard is one library card. Kind 1 playlist, 2 album, 3 artist, 4 podcast.
type LibCard struct {
	Index, Kind int
	Title, Sub  string
	Cover       string
	Active      bool
}

// LibraryData holds the library filter pills and the media grid.
type LibraryData struct {
	Filter int
	All    []LibCard
	Cards  []LibCard
}

// defaultLibrary is the offline library screen.
func defaultLibrary() LibraryData {
	all := []LibCard{
		{0, 1, "Neon Nights", "Playlist · Chinmay", "card-0", false},
		{1, 1, "Focus Flow", "Playlist · Chinmay", "card-1", false},
		{2, 1, "Late Drive", "Playlist · Spotify", "card-2", false},
		{3, 1, "Rainy Day", "Playlist · Spotify", "card-3", false},
		{4, 2, "Neon Horizon", "Album · Nova Waves", "card-4", false},
		{5, 2, "Static Bloom", "Album · The Far Coast", "card-5", false},
		{6, 2, "Signal Fire", "Album · Aster Field", "card-0", false},
		{7, 3, "Nova Waves", "Artist", "card-1", false},
		{8, 3, "Aster Field", "Artist", "card-2", false},
		{9, 4, "Signal Path", "Podcast · Weekly", "card-3", false},
	}

	return LibraryData{All: all, Cards: append([]LibCard(nil), all...)}
}

// LibraryFilter rebuilds Cards for the current Filter.
func (a *App) LibraryFilter() {
	lib := &a.view.Library
	lib.Cards = lib.Cards[:0]
	for _, c := range lib.All {
		if lib.Filter == 0 || c.Kind == lib.Filter {
			c.Index = len(lib.Cards)
			c.Active = false
			lib.Cards = append(lib.Cards, c)
		}
	}
}

// clickLibrary applies one action on the library screen.
func (a *App) clickLibrary(ctx context.Context, action string) error {
	switch {
	case strings.HasPrefix(action, "library-filter-"):
		f := slot(action, "library-filter-")
		if f < 0 || f > 4 {
			return nil
		}

		a.view.Library.Filter = f
		a.LibraryFilter()
	case strings.HasPrefix(action, "library-card-"):
		cards := a.view.Library.Cards
		i := slot(action, "library-card-")
		for j := range cards {
			cards[j].Active = j == i
		}
	}

	return nil
}
