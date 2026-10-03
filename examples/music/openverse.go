package music

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	openverseBase = "https://api.openverse.org"
	maxClipBytes  = 12 << 20
	searchPage    = 20 // the anonymous page_size limit
)

// Openverse resolves free music through the Openverse audio API. No key is
// needed. Downloads go to CacheDir when it is set, so a replay is offline.
type Openverse struct {
	Base     string // default https://api.openverse.org
	CacheDir string // "" keeps downloads in memory only
}

// NewOpenverse returns an Openverse resolver with the public API and the
// user's cache directory for downloaded audio.
func NewOpenverse() *Openverse {
	dir := ""
	if root, err := os.UserCacheDir(); err == nil {
		dir = filepath.Join(root, "go-gpui", "music")
	}

	return &Openverse{Base: openverseBase, CacheDir: dir}
}

// Resolve searches Openverse and downloads the pick-th playable MP3. A pick
// outside the result count wraps, so every track gets a clip when it can.
func (o *Openverse) Resolve(ctx context.Context, query string, pick int) (Clip, error) {
	results, err := o.search(ctx, query)
	if err != nil {
		return Clip{}, err
	}

	if len(results) == 0 {
		return Clip{}, fmt.Errorf("music: no free track for %q", query)
	}

	chosen := results[((pick%len(results))+len(results))%len(results)]

	data, err := o.download(ctx, chosen.URL)
	if err != nil {
		return Clip{}, err
	}

	return Clip{
		Title:   chosen.Title,
		Creator: chosen.Creator,
		License: licenseName(chosen.License, chosen.LicenseVersion),
		URL:     chosen.URL,
		Data:    data,
	}, nil
}

// base is the API root without a trailing slash.
func (o *Openverse) base() string {
	if o.Base == "" {
		return openverseBase
	}

	return strings.TrimSuffix(o.Base, "/")
}
