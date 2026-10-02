package fetcher

import (
	"bytes"
	"image"
	_ "image/jpeg" // registers JPEG for DecodeConfig
	_ "image/png"  // registers PNG for DecodeConfig
	"strconv"

	"github.com/chinmay-sawant/go-gpui"
)

// backgroundSrc is the source name the template's background points at.
const backgroundSrc = "fetched"

// applyImage makes a PNG or JPEG body the page background and rewrites the
// status line with the pixel size. Any other body leaves both alone.
func (a *App) applyImage(res gpui.FetchResponse) {
	cfg, format, err := image.DecodeConfig(bytes.NewReader(res.Body))
	if err != nil || (format != "png" && format != "jpeg") {
		return
	}

	a.page.SetImage(backgroundSrc, res.Body)
	a.view.Status = "status=" + strconv.Itoa(res.Status) +
		" bytes=" + strconv.Itoa(len(res.Body)) +
		" image=" + strconv.Itoa(cfg.Width) + "x" + strconv.Itoa(cfg.Height)
}
