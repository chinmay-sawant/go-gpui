package dino

// view maps the fixed 900x300 game scene onto the page. A touch page
// scales the scene to its width, never up, and drops it to the bottom, so
// a taller phone screen becomes sky above the ground.
type view struct {
	scale float64
	oy    float64
}

// fitView returns the view for a page of pageW x pageH CSS pixels.
func fitView(pageW, pageH int) view {
	if pageW <= 0 || pageH <= 0 {
		return view{scale: 1}
	}

	scale := min(float64(pageW)/sceneW, float64(pageH)/sceneH, 1)

	return view{scale: scale, oy: float64(pageH) - sceneH*scale}
}

// syncView refits the scene when the page size changed. Desktop pages keep
// the identity view.
func (a *App) syncView() {
	if !a.touch {
		return
	}

	w, h := a.page.Size()
	if w == a.viewW && h == a.viewH {
		return
	}

	a.viewW, a.viewH = w, h
	a.view = fitView(w, h)
}
