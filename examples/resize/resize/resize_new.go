package resize

import "github.com/chinmay-sawant/go-gpui"

// New parses the embedded resize template.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "Resize",
		HTML:   resizeHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	return &App{page: page}, nil
}
