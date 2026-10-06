package main

import (
	"context"
	"strconv"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func installHover(p *page.Page) error {
	if err := p.SetTheme(`.row{background:#141820}`); err != nil {
		return err
	}
	cache := rowPaintCache{rows: map[string]*gpui.DisplayOp{}}
	p.Handle(gpui.Handlers{Hover: func(_ context.Context, old, next gpui.Box) (bool, error) {
		allowed := func(id string) bool {
			return id == "" || id == "grid" || id == "top-space" || id == "bottom-space" || strings.HasPrefix(id, "row-")
		}
		if !allowed(old.ID) || !allowed(next.ID) {
			return false, nil
		}
		cache.bind(p)
		if op := cache.rows[old.ID]; op != nil {
			n, _ := strconv.Atoi(strings.TrimPrefix(old.ID, "row-"))
			color := [3]float64{20, 24, 32}
			if n%2 == 0 {
				color = [3]float64{29, 37, 50}
			}
			op.R, op.G, op.B = color[0]/255, color[1]/255, color[2]/255
		}
		if op := cache.rows[next.ID]; op != nil {
			op.R, op.G, op.B = 52.0/255, 65.0/255, 89.0/255
		}
		return true, nil
	}})
	return nil
}
