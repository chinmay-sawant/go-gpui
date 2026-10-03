package music

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chinmay-sawant/go-gpui"
)

// download returns the bytes for raw, from the cache when it has them.
func (o *Openverse) download(ctx context.Context, raw string) ([]byte, error) {
	path := o.cachePath(raw)
	if path != "" {
		if data, err := os.ReadFile(path); err == nil && len(data) > 0 {
			return data, nil
		}
	}

	res, err := gpui.Fetch(ctx, raw)
	if err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, fmt.Errorf("music: track status %d", res.Status)
	}

	if len(res.Body) == 0 || len(res.Body) > maxClipBytes {
		return nil, fmt.Errorf("music: track size %d", len(res.Body))
	}

	if path != "" {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err == nil {
			_ = os.WriteFile(path, res.Body, 0o644)
		}
	}

	return res.Body, nil
}

// cachePath names the cache file for one URL, or "" without a cache dir.
func (o *Openverse) cachePath(raw string) string {
	if o.CacheDir == "" {
		return ""
	}

	sum := sha1.Sum([]byte(raw))

	return filepath.Join(o.CacheDir, hex.EncodeToString(sum[:])+".mp3")
}
