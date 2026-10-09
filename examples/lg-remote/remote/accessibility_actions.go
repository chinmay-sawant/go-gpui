package remote

import (
	"context"
	"strconv"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"
)

func (a *App) accessibilityActions(ctx context.Context) error {
	for action := bridge.TakeAction(); action != ""; action = bridge.TakeAction() {
		kind, value, _ := strings.Cut(action, ":")
		switch kind {
		case "sensitivity":
			v, err := strconv.Atoi(value)
			if err == nil {
				a.setSensitivity(float64(v))
				a.setData()
				if err := a.page.Redraw(ctx); err != nil {
					return err
				}
			}
		case "blur":
			if err := a.page.Focus(ctx, ""); err != nil {
				return err
			}
		case "panel":
			if value == "remote" {
				a.show("remote", "Volume on the left, channels on the right.")
				a.setData()
				if err := a.page.Redraw(ctx); err != nil {
					return err
				}
			}
		case "text":
			a.page.SetFormValue("host", value)
			if err := a.page.Redraw(ctx); err != nil {
				return err
			}
		case "scroll":
			if a.page.ViewLocked() || !a.page.AllowScroll() {
				continue
			}
			_, h := a.page.Size()
			if value == "back" {
				h = -h
			}
			a.page.ScrollBy(0, h*3/4)
		case "click":
			for _, box := range a.page.Boxes() {
				if box.ID == value {
					if err := a.page.Click(ctx, box.X+box.W/2, box.Y+box.H/2); err != nil {
						return err
					}
					break
				}
			}
		}
	}
	return nil
}
