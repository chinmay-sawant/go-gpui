package fetcher

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// Get sends a GET with gpui.Fetch and records the result in View.Status.
// The returned error is also recorded, so a caller can show or ignore it.
func (a *App) Get(ctx context.Context, url string) error {
	res, err := gpui.Fetch(ctx, url)
	if err != nil {
		a.view.Status = "error: " + err.Error()
		return err
	}

	a.view.Status = summary(res)

	return nil
}

// Post sends a small text body with gpui.XHR POST and records the result.
func (a *App) Post(ctx context.Context, url string) error {
	res, err := gpui.XHR(ctx, "POST", url,
		map[string]string{"Content-Type": "text/plain"},
		[]byte("note=hello"),
	)
	if err != nil {
		a.view.Status = "error: " + err.Error()
		return err
	}

	a.view.Status = summary(res)

	return nil
}
