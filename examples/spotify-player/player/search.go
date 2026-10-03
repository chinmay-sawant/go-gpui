package player

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/chinmay-sawant/go-gpui"
)

// search GETs one iTunes search page.
func (a *App) search(ctx context.Context, term, entity string, limit int) ([]searchResult, error) {
	q := url.Values{}
	q.Set("term", term)
	q.Set("entity", entity)
	q.Set("limit", strconv.Itoa(limit))

	res, err := gpui.Fetch(ctx, a.base+"/search?"+q.Encode())
	if err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, fmt.Errorf("player: search status %d", res.Status)
	}

	var out searchResponse
	if err := json.Unmarshal(res.Body, &out); err != nil {
		return nil, err
	}

	return out.Results, nil
}
