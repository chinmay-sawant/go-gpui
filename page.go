// Package gpui shows an HTML page in a window.
//
// A program passes a template to New, stores the template data with SetData,
// and registers Handlers for clicks and keys. Run opens the window. Serve
// shows the same picture in a browser. BindMobile registers the page for an
// Android or iOS bind. Redraw parses the HTML, applies the CSS, and lays
// the page out through gowkhtmltopdf. This package has no layout of its own.
package gpui

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"image"
	"image/png"
	"strings"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/html"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// Box is one element a click can land on.
// X, Y, W, H are CSS pixels, origin at the top left of the picture.
type Box = layout.Box

const defaultMax = 2560

var (
	// ErrEmptyHTML means New was given a blank template.
	ErrEmptyHTML = errors.New("gpui: empty html")

	// ErrBadSize means a width or a height is unusable.
	ErrBadSize = errors.New("gpui: bad size")

	// ErrNilPage means Run, Serve, or BindMobile was called without a page.
	ErrNilPage = errors.New("gpui: nil page")

	errNilContext = errors.New("gpui: nil context")
)

// Config is the template and the frame size for a new page.
// Width and Height are the first frame, in CSS pixels.
// MinWidth and MinHeight are the smallest frame. Zero means 1.
// MaxWidth and MaxHeight cap the picture so a large monitor does not
// allocate a matching PNG. Zero means 2560.
type Config struct {
	Title     string
	HTML      string
	Width     int
	Height    int
	MinWidth  int
	MinHeight int
	MaxWidth  int
	MaxHeight int
}

// Handlers are the Go functions for input.
// A nil function ignores that input. After a function returns, the page
// is drawn again.
type Handlers struct {
	Click     func(ctx context.Context, box Box) error
	Type      func(ctx context.Context, text string) error
	Backspace func(ctx context.Context) error
	Submit    func(ctx context.Context) error
}

// Page is one HTML template and the last picture it produced.
type Page struct {
	title      string
	tpl        *template.Template
	data       any
	handlers   Handlers
	img        image.Image
	boxes      []layout.Box
	width      int
	height     int
	minWidth   int
	minHeight  int
	maxWidth   int
	maxHeight  int
	generation uint64
}

// New parses html as an html/template page.
func New(cfg Config) (*Page, error) {
	if strings.TrimSpace(cfg.HTML) == "" {
		return nil, ErrEmptyHTML
	}

	minWidth := cfg.MinWidth
	if minWidth <= 0 {
		minWidth = 1
	}

	minHeight := cfg.MinHeight
	if minHeight <= 0 {
		minHeight = 1
	}

	maxWidth := cfg.MaxWidth
	if maxWidth <= 0 {
		maxWidth = defaultMax
	}

	maxHeight := cfg.MaxHeight
	if maxHeight <= 0 {
		maxHeight = defaultMax
	}

	if cfg.Width <= 0 || cfg.Height <= 0 || minWidth > maxWidth || minHeight > maxHeight {
		return nil, ErrBadSize
	}

	tpl, err := template.New("page").Parse(cfg.HTML)
	if err != nil {
		return nil, err
	}

	title := cfg.Title
	if title == "" {
		title = "go-gpui"
	}

	page := &Page{
		title:     title,
		tpl:       tpl,
		minWidth:  minWidth,
		minHeight: minHeight,
		maxWidth:  maxWidth,
		maxHeight: maxHeight,
	}
	page.width, page.height = page.Clamp(cfg.Width, cfg.Height)

	return page, nil
}

// Handle registers the input functions. A later call replaces them.
func (p *Page) Handle(h Handlers) {
	p.handlers = h
}

// SetData stores the value the template prints on the next Redraw.
func (p *Page) SetData(data any) {
	p.data = data
}

// Title returns the window title.
func (p *Page) Title() string {
	return p.title
}

// Size returns the frame size in CSS pixels.
func (p *Page) Size() (int, int) {
	return p.width, p.height
}

// MinSize returns the smallest frame this page will draw.
func (p *Page) MinSize() (int, int) {
	return p.minWidth, p.minHeight
}

// Clamp pulls a size into the min and max this page will draw.
func (p *Page) Clamp(width, height int) (int, int) {
	return clamp(width, p.minWidth, p.maxWidth), clamp(height, p.minHeight, p.maxHeight)
}

// SetSize stores the frame size used by the next Redraw.
// Values outside the min and max are pulled back inside that range.
func (p *Page) SetSize(width, height int) {
	p.width, p.height = p.Clamp(width, height)
}

// Generation increases by one on every successful Redraw.
func (p *Page) Generation() uint64 {
	return p.generation
}

// Redraw fills the template and renders the current size.
func (p *Page) Redraw(ctx context.Context) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	var body strings.Builder
	if err := p.tpl.Execute(&body, p.data); err != nil {
		return err
	}

	doc, err := html.Parse([]byte(body.String()))
	if err != nil {
		return err
	}

	styled, err := css.Apply(ctx, doc, css.Options{
		WidthPx:  p.width,
		HeightPx: p.height,
		Media:    "screen",
		Extra:    nil,
	})
	if err != nil {
		return err
	}

	placed, err := layout.Lay(ctx, styled)
	if err != nil {
		return err
	}

	p.img = placed.Image()
	p.boxes = placed.Boxes()
	p.generation++

	return nil
}

// Image returns the last picture, or nil when nothing has been drawn.
func (p *Page) Image() image.Image {
	return p.img
}

// PNG encodes Image to PNG bytes.
// It returns nil when Image is nil.
func (p *Page) PNG() []byte {
	img := p.Image()
	if img == nil {
		return nil
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil
	}

	return buf.Bytes()
}

// Boxes returns the last hit-test boxes, or nil when nothing has been drawn.
func (p *Page) Boxes() []Box {
	return p.boxes
}

// Click hit-tests the last picture and calls the click handler.
// The innermost box is the last one in document order that contains the point.
// Click then draws the page again.
func (p *Page) Click(ctx context.Context, x, y float64) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if box, ok := hit(p.Boxes(), x, y); ok && p.handlers.Click != nil {
		if err := p.handlers.Click(ctx, box); err != nil {
			return err
		}
	}

	return p.Redraw(ctx)
}

// Type calls the type handler and draws the page again.
// It does nothing when no type handler is registered.
func (p *Page) Type(ctx context.Context, text string) error {
	return p.after(ctx, p.handlers.Type == nil, func() error {
		return p.handlers.Type(ctx, text)
	})
}

// Backspace calls the backspace handler and draws the page again.
// It does nothing when no backspace handler is registered.
func (p *Page) Backspace(ctx context.Context) error {
	return p.after(ctx, p.handlers.Backspace == nil, func() error {
		return p.handlers.Backspace(ctx)
	})
}

// Submit calls the submit handler and draws the page again.
// It does nothing when no submit handler is registered.
func (p *Page) Submit(ctx context.Context) error {
	return p.after(ctx, p.handlers.Submit == nil, func() error {
		return p.handlers.Submit(ctx)
	})
}

func (p *Page) after(ctx context.Context, skip bool, fn func() error) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if skip {
		return nil
	}

	if err := fn(); err != nil {
		return err
	}

	return p.Redraw(ctx)
}

func useContext(ctx context.Context) error {
	if ctx == nil {
		return errNilContext
	}

	return ctx.Err()
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

// hit returns the last box that contains x, y.
// Boxes are in document order, so that box is the inner element.
func hit(boxes []Box, x, y float64) (Box, bool) {
	var found Box
	ok := false

	for _, b := range boxes {
		if x < b.X || y < b.Y || x > b.X+b.W || y > b.Y+b.H {
			continue
		}

		found = b
		ok = true
	}

	return found, ok
}

var _ host.Screen = (*Page)(nil)
