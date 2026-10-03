package player

import (
	"errors"
	"fmt"
	"strings"
)

// errNoSongs means the search answered with no usable song.
var errNoSongs = errors.New("itunes: no songs")

// liveView builds the playing view from freshly fetched tracks.
func liveView(term string, tracks []Track) View {
	v := DefaultView()
	v.Status = "Live from iTunes · " + term
	v.Query = term
	v.Playing = true
	v.Queue = tracks
	v.Recent = recentFrom(tracks)
	v.Now = tracks[0]
	v.QueueCount = countLabel(len(tracks))
	v.setProgress(0)

	return v
}

// trackFrom converts one song result. Missing fields keep a readable label.
func trackFrom(i int, song itunesSong) Track {
	return Track{
		Index:  i,
		Title:  label(song.TrackName, "Unknown title"),
		Artist: label(song.ArtistName, "Unknown artist"),
		Album:  label(song.CollectionName, "Unknown album"),
		Genre:  label(song.PrimaryGenreName, "Music"),
		Length: clock(song.TrackTimeMillis / 1000),
		Cover:  fmt.Sprintf("cover-%d", i),
		Active: i == 0,
	}
}

// label falls back when a result field is blank.
func label(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}

	return value
}

// recentFrom keeps up to the last four queue entries.
func recentFrom(queue []Track) []Track {
	if len(queue) > 4 {
		queue = queue[len(queue)-4:]
	}

	return append([]Track(nil), queue...)
}

// countLabel is the queue header count.
func countLabel(n int) string {
	if n == 1 {
		return "1 track"
	}

	return fmt.Sprintf("%d tracks", n)
}
