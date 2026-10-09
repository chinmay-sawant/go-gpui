package remote

import "context"

func (a *App) sensitivityKey(ctx context.Context, key string) error {
	value := a.view.Sensitivity
	switch key {
	case "arrowleft", "arrowdown":
		value -= 5
	case "arrowright", "arrowup":
		value += 5
	case "home":
		value = 25
	case "end":
		value = 300
	default:
		return nil
	}
	a.setSensitivity(float64(value))
	a.flash("sensitivity")
	a.setData()
	return a.page.Redraw(ctx)
}
