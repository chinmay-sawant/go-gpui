package ui

// CanPause reports whether the row offers Pause.
func (r Row) CanPause() bool { return r.State == StateRunning }

// CanResume reports whether the row offers Resume.
func (r Row) CanResume() bool { return r.State == StatePaused }

// CanCancel reports whether the row offers Cancel.
func (r Row) CanCancel() bool { return r.State.Active() }

// CanRetry reports whether the row offers Retry.
func (r Row) CanRetry() bool { return r.State.Terminal() }

// CanRemove reports whether the row offers Remove.
func (r Row) CanRemove() bool { return r.State.Terminal() }
