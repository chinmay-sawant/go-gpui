package player

import (
	"strconv"
	"strings"
)

// tracksFrom converts songs into tracklist rows.
func tracksFrom(songs []searchResult) []Track {
	tracks := make([]Track, 0, len(songs))

	for i, s := range songs {
		album := s.CollectionName
		if album == "" {
			album = s.Genre
		}

		tracks = append(tracks, Track{
			Index: i, Num: strconv.Itoa(i + 1),
			Title: clip(s.TrackName, 44), Artist: clip(s.ArtistName, 40),
			Album: clip(album, 26), Year: yearOf(s.ReleaseDate),
			Genre:  clip(s.Genre, 24),
			Length: formatMillis(s.TrackMillis),
			Cover:  "track-" + strconv.Itoa(i),
			Active: i == 0,
		})
	}

	return tracks
}

// cardsFrom converts albums into six greeting tiles and six shelf cards.
func cardsFrom(albums []searchResult) (picks, shelf []Card) {
	for i, al := range albums {
		if i < 6 {
			picks = append(picks, Card{Index: i, Title: clip(al.CollectionName, 32), Sub: clip(al.ArtistName, 32), Cover: "pick-" + strconv.Itoa(i)})
			continue
		}

		j := i - 6
		shelf = append(shelf, Card{Index: j, Title: clip(al.CollectionName, 19), Sub: clip(al.ArtistName, 19), Cover: "card-" + strconv.Itoa(j)})
	}

	return picks, shelf
}

// playlistsFrom fills the sidebar library from the shelf cards.
func playlistsFrom(shelf []Card) []Playlist {
	out := make([]Playlist, 0, 5)

	for _, card := range shelf {
		if len(out) == 5 {
			break
		}

		out = append(out, Playlist{
			Index: len(out),
			Name:  clip(card.Title, 24),
			Meta:  clip("Playlist · "+card.Sub, 24),
			Cover: card.Cover,
		})
	}

	return out
}

// upscale asks iTunes for a larger copy of the same artwork.
func upscale(raw string) string {
	return strings.Replace(raw, "100x100bb", "300x300bb", 1)
}
