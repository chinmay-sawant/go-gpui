package player

import (
	"context"
	"strconv"
	"strings"
)

// BrowseHero is one wide card in the featured band.
type BrowseHero struct {
	Title, Sub, Img string
	Active          bool
}

// BrowseTile is one square genre tile.
type BrowseTile struct {
	Title  string
	Tone   int
	Img    string
	Active bool
}

// BrowseData is the data behind the browse screen: the featured band and
// the genre grid.
type BrowseData struct {
	Hero  []BrowseHero
	Tiles []BrowseTile
}

// defaultBrowse returns the offline browse screen.
func defaultBrowse() BrowseData {
	titles := []string{
		"Chill", "Focus", "Workout", "Party", "Sleep", "Chillhop",
		"Indie", "Pop", "Rock", "Jazz", "Classical", "Electronic",
	}

	hero := []BrowseHero{
		{Title: "Midnight Drive", Sub: "Nova Waves · Album", Img: "card-3"},
		{Title: "Neon Skyline", Sub: "Lumen Fields · Album", Img: "card-4"},
		{Title: "Slow Motion", Sub: "Velvet Static · Album", Img: "card-5"},
	}

	tiles := make([]BrowseTile, len(titles))
	for i, title := range titles {
		tiles[i] = BrowseTile{Title: title, Tone: i, Img: "card-" + strconv.Itoa(i%6)}
	}

	return BrowseData{Hero: hero, Tiles: tiles}
}

// BrowseSelectHero highlights one featured card.
func (a *App) BrowseSelectHero(i int) {
	if i < 0 || i >= len(a.view.Browse.Hero) {
		return
	}

	for j := range a.view.Browse.Hero {
		a.view.Browse.Hero[j].Active = j == i
	}
}

// BrowseSelectTile highlights one genre tile.
func (a *App) BrowseSelectTile(i int) {
	if i < 0 || i >= len(a.view.Browse.Tiles) {
		return
	}

	for j := range a.view.Browse.Tiles {
		a.view.Browse.Tiles[j].Active = j == i
	}
}

// clickBrowse applies one action on the browse screen.
func (a *App) clickBrowse(ctx context.Context, action string) error {
	switch {
	case strings.HasPrefix(action, "browse-hero-"):
		a.BrowseSelectHero(slot(action, "browse-hero-"))
	case strings.HasPrefix(action, "browse-tile-"):
		a.BrowseSelectTile(slot(action, "browse-tile-"))
	}

	return nil
}
