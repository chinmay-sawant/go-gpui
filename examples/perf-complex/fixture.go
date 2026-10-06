package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

//go:embed layout.html
var source string

type row struct {
	ID, Region int
	Name       string
}
type view struct {
	Revision, Top, Bottom int
	Nav                   []string
	Rows                  []row
}

func fixture(window bool) (*page.Page, view, error) {
	v := view{Nav: []string{"Overview", "Services", "Regions", "Deployments", "Incidents", "Capacity"}}
	count := 480
	if window {
		count = 48
		v.Bottom = (480 - count) * 76
	}
	for i := 0; i < count; i++ {
		v.Rows = append(v.Rows, row{i, i % 8, fmt.Sprintf("service-%04d", i)})
	}
	p, err := page.New(page.Config{HTML: source, Width: 1440, Height: 1000, Perf: true})
	if err == nil {
		p.SetData(v)
	}
	return p, v, err
}

func redraw(p *page.Page) error {
	if err := p.Redraw(context.Background()); err != nil {
		return err
	}
	if p.Display() == nil {
		return fmt.Errorf("fixture triggered raster fallback")
	}
	p.TakeDirty()
	return nil
}
