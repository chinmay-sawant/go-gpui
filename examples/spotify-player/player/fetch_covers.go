package player

import (
	"bytes"
	"context"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"sync"

	"github.com/chinmay-sawant/ownframe"
)

// fetchCovers downloads the jobs, at most maxCoverWorkers at a time, and
// returns the decoded bytes by image name. Duplicate URLs download once.
func (a *App) fetchCovers(ctx context.Context, jobs []coverJob) map[string][]byte {
	names := map[string][]string{}
	var urls []string

	for _, job := range jobs {
		if job.url == "" {
			continue
		}

		if _, ok := names[job.url]; !ok {
			urls = append(urls, job.url)
		}

		names[job.url] = append(names[job.url], job.name)
	}

	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		out = map[string][]byte{}
		sem = make(chan struct{}, maxCoverWorkers)
	)

	for _, raw := range urls {
		wg.Add(1)

		go func(raw string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			res, err := ownframe.Fetch(ctx, raw)
			if err != nil || res.Status != 200 || !isImage(res.Body) {
				return
			}

			mu.Lock()
			for _, name := range names[raw] {
				out[name] = res.Body
			}
			mu.Unlock()
		}(raw)
	}

	wg.Wait()

	return out
}

// isImage reports whether b decodes as a png or jpeg.
func isImage(b []byte) bool {
	_, format, err := image.DecodeConfig(bytes.NewReader(b))
	if err != nil {
		return false
	}

	return format == "png" || format == "jpeg"
}
