package login

import "context"

func (a *App) onUndo(context.Context) error {
	if len(a.undo) == 0 {
		return nil
	}

	a.redo = append(a.redo, a.view)
	a.view = a.undo[len(a.undo)-1]
	a.undo = a.undo[:len(a.undo)-1]
	a.page.SetData(a.view)

	return nil
}

func (a *App) onRedo(context.Context) error {
	if len(a.redo) == 0 {
		return nil
	}

	a.undo = append(a.undo, a.view)
	a.view = a.redo[len(a.redo)-1]
	a.redo = a.redo[:len(a.redo)-1]
	a.page.SetData(a.view)

	return nil
}

func (a *App) onSubmit(_ context.Context) error {
	a.signIn()
	a.page.SetData(a.view)

	return nil
}
