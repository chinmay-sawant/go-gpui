package music

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/chinmay-sawant/ownframe"
)

// search returns the playable MP3 results for query.
func (o *Openverse) search(ctx context.Context, query string) ([]ovResult, error) {
	q := url.Values{}
	q.Set("q", query)
	q.Set("page_size", strconv.Itoa(searchPage))
	q.Set("license_type", "commercial")
	q.Set("category", "music")
	q.Set("mature", "false")

	res, err := ownframe.Fetch(ctx, o.base()+"/v1/audio/?"+q.Encode())
	if err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, fmt.Errorf("music: openverse status %d", res.Status)
	}

	var payload ovResponse
	if err := json.Unmarshal(res.Body, &payload); err != nil {
		return nil, err
	}

	out := make([]ovResult, 0, len(payload.Results))
	for _, r := range payload.Results {
		if r.playable() {
			out = append(out, r)
		}
	}

	return out, nil
}
