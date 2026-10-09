package remote

import "context"

func (a *App) buttonTouch(ctx context.Context, kind, id string) error {
	if kind != "down" {
		a.stopRepeat()
		a.nativeHeld = ""
		if kind == "cancel" {
			a.view.PressedID = ""
		}
		a.touchDirty = true
		return nil
	}
	key, enabled := a.controlNames()[id]
	if !enabled || key.Disabled {
		return nil
	}
	for _, box := range a.page.Boxes() {
		if box.ID != id {
			continue
		}
		a.nativeHeld = id
		a.armRepeat(id)
		a.touchDirty = true
		return a.onClick(ctx, box)
	}
	return nil
}
