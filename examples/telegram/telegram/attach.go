package telegram

// RequestAttach asks the phone for a photo for the open chat. The kind is
// "camera" or "gallery"; the activity polls TakeAttach for it.
func (a *App) RequestAttach(kind string) {
	a.mu.Lock()
	a.attachID = a.view.Active
	a.attachKind = kind
	a.mu.Unlock()
}

// TakeAttach reports and clears a pending photo request. It returns
// "camera", "gallery", or "" when there is none.
func (a *App) TakeAttach() string {
	a.mu.Lock()
	defer a.mu.Unlock()

	kind := a.attachKind
	a.attachKind = ""

	return kind
}

// QueuePhoto stores captured image bytes for the next tick.
func (a *App) QueuePhoto(data []byte) {
	select {
	case a.photos <- data:
	default:
	}
}

// applyPhoto adds the most recent captured photo to the chat that asked.
func (a *App) applyPhoto() {
	select {
	case data := <-a.photos:
		a.addPhoto(data)
	default:
	}
}

// DarkTheme reports whether the dark theme is showing, so the activity can
// match the system bar icons.
func (a *App) DarkTheme() bool { return a.dark.Load() }
