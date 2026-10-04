package page

import (
	"context"
	"html/template"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// Display is the retained vector list behind a replayable page.
type Display = layout.Display

// DisplayOp is one operation in a Display. A frame callback may change its
// paint fields; see Page.SetTick.
type DisplayOp = layout.DisplayOp

// DisplayOpFillRect and DisplayOpText are the operation kinds the frame
// helpers look for.
const (
	DisplayOpFillRect = layout.DisplayOpFillRect
	DisplayOpText     = layout.DisplayOpText
)

// Page is one HTML template and the last picture it produced.
type Page struct {
	title      string
	tpl        *template.Template
	data       any
	theme      *css.Sheet
	themeSrc   string
	handlers   Handlers
	images     map[string][]byte
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
	tick       func(ctx context.Context) error
	past       []string
	pastAt     int
	routes     map[string]string
	form       *formState
	picker     PickFunc
	hover      string
	active     string
	cache      *render.Cache
	stats      pageStats
	devtools   bool
}
