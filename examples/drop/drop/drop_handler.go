package drop

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg" // registers JPEG for DecodeConfig
	_ "image/png"  // registers PNG for DecodeConfig
	"path/filepath"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// onDrop previews the files of one drop. The last file wins. A read error
// becomes the status line, so the window keeps running.
func (a *App) onDrop(_ context.Context, files []gpui.Drop) error {
	for _, file := range files {
		a.show(file)
	}

	return nil
}

func (a *App) show(file gpui.Drop) {
	if file.IsDir {
		a.view = View{Status: file.Name + ": directory"}

		return
	}

	data, err := file.Read()
	if err != nil {
		a.view = View{Status: file.Name + ": " + err.Error()}

		return
	}

	if cfg, format, err := image.DecodeConfig(bytes.NewReader(data)); err == nil {
		if format == "png" || format == "jpeg" {
			a.page.SetImage(imageSrc, data)
			a.view = View{
				Status: fmt.Sprintf("%s: %dx%d image", file.Name, cfg.Width, cfg.Height),
				Image:  true,
			}

			return
		}
	}

	if strings.EqualFold(filepath.Ext(file.Name), ".txt") {
		a.view = View{Status: file.Name, Lines: firstLines(string(data))}

		return
	}

	a.view = View{Status: file.Name + ": no preview"}
}

// firstLines returns at most maxLines lines and drops the carriage returns of
// a CRLF file.
func firstLines(text string) []string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return lines
}
