package login

import "context"

func (a *App) onCopy(context.Context) (string, bool, error) {
	field := a.focused()
	if field == nil {
		return "", false, nil
	}

	return *field, true, nil
}

func (a *App) onCut(ctx context.Context) (string, bool, error) {
	text, ok, err := a.onCopy(ctx)
	if err != nil || !ok {
		return text, ok, err
	}

	if field := a.focused(); field != nil {
		a.setField(field, "")
	}

	return text, true, nil
}

func (a *App) onSelectAll(context.Context) error {
	if a.focused() == nil || a.view.Selected {
		return nil
	}

	a.push()
	a.view.Selected = true
	a.page.SetData(a.view)

	return nil
}
