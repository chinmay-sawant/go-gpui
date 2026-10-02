package page

import (
	"html/template"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// Page is one HTML template and the last picture it produced.
type Page struct {
	title      string
	tpl        *template.Template
	data       any
	handlers   Handlers
	img        image.Image
	display    *layout.Display
	boxes      []layout.Box
	source     string
	png        []byte
	width      int
	height     int
	minWidth   int
	minHeight  int
	maxWidth   int
	maxHeight  int
	generation uint64
	past       []string
	pastAt     int
	routes     map[string]string
	form       *formState
	picker     PickFunc
	hover      string
	active     string
}
