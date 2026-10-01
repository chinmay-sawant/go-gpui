package page

import (
	"html/template"
	"strings"
)

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
