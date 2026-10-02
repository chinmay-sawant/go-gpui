package scrolling

import "github.com/chinmay-sawant/go-gpui"

// New parses the embedded scrolling template.
func New() (*App, error) {
	p, err := gpui.New(gpui.Config{
		Title:  "Scrolling",
		HTML:   scrollingHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	return &App{page: p}, nil
}
