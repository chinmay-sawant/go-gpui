package player

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// Load fills the queue from the iTunes search API. On any failure the sample
// view stays and Status says so. New never calls Load.
func (a *App) Load(ctx context.Context, term string) error {
	term = strings.TrimSpace(term)
	if term == "" {
		term = defaultTerm
	}

	songs, err := a.searchSongs(ctx, term)
	if err != nil {
		return a.offline(term, err)
	}

	tracks := make([]Track, 0, len(songs))
	for i, song := range songs {
		tracks = append(tracks, trackFrom(i, song))
	}

	if len(tracks) == 0 {
		return a.offline(term, errNoSongs)
	}

	a.view = liveView(term, tracks)
	a.loadCovers(ctx, songs, tracks)
	a.page.SetData(a.view)

	return nil
}

// searchSongs requests up to searchLimit songs and keeps only kind=="song".
func (a *App) searchSongs(ctx context.Context, term string) ([]itunesSong, error) {
	query := url.Values{}
	query.Set("term", term)
	query.Set("entity", "song")
	query.Set("limit", strconv.Itoa(searchLimit))

	res, err := gpui.Fetch(ctx, a.base+"/search?"+query.Encode())
	if err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, fmt.Errorf("itunes: status %d", res.Status)
	}

	var payload itunesResponse
	if err := json.Unmarshal(res.Body, &payload); err != nil {
		return nil, err
	}

	songs := make([]itunesSong, 0, len(payload.Results))
	for _, song := range payload.Results {
		if song.Kind != "song" {
			continue
		}

		songs = append(songs, song)
	}

	return songs, nil
}

// offline restores the sample view and records why.
func (a *App) offline(term string, cause error) error {
	a.view = DefaultView()
	a.view.Query = term
	a.view.Status = offlineStatus
	a.page.SetData(a.view)

	return cause
}
