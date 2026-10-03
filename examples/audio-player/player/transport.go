package player

// step moves the active queue index by delta with wrap.
func (a *App) step(delta int) {
	n := len(a.view.Queue)
	if n == 0 {
		return
	}

	a.selectQueue((a.activeIndex() + delta + n) % n)
}

// activeIndex is the queue position marked active, or Now's own index.
func (a *App) activeIndex() int {
	for i := range a.view.Queue {
		if a.view.Queue[i].Active {
			return i
		}
	}

	return clamp(a.view.Now.Index, 0, len(a.view.Queue)-1)
}

// selectQueue makes queue entry i the current track and restarts its clock.
func (a *App) selectQueue(i int) {
	if len(a.view.Queue) == 0 {
		return
	}

	i = clamp(i, 0, len(a.view.Queue)-1)
	for j := range a.view.Queue {
		a.view.Queue[j].Active = j == i
	}

	a.view.Now = a.view.Queue[i]
	a.view.setProgress(0)
}

// seekTo moves the seek bar to the right edge of ten-step segment i.
func (a *App) seekTo(i int) {
	a.view.setProgress((i + 1) * 10)
}

// setVolume moves the volume bar to segment i's edge and mutes at zero.
func (a *App) setVolume(i int) {
	a.view.Volume = clamp((i+1)*10, 0, 100)
	a.view.Muted = a.view.Volume == 0
}

// selectPlaylist highlights one sidebar row.
func (a *App) selectPlaylist(action string) {
	i, ok := actionIndex(action, "list-")
	if !ok {
		return
	}

	for j := range a.view.Playlists {
		a.view.Playlists[j].Active = j == i
	}
}
