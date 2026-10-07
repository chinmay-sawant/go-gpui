package ui

// AddRequest is one queued job. An empty Destination means the backend
// picks the default directory.
type AddRequest struct {
	URL         string
	Destination string
}

// ControlAction is one job control command.
type ControlAction uint8

// The control actions a queue row offers.
const (
	ControlPause ControlAction = iota
	ControlResume
	ControlCancel
	ControlRetry
	ControlRemove
)

// Label returns the button text for the action.
func (c ControlAction) Label() string {
	switch c {
	case ControlPause:
		return "Pause"
	case ControlResume:
		return "Resume"
	case ControlCancel:
		return "Cancel"
	case ControlRetry:
		return "Retry"
	default:
		return "Remove"
	}
}

// Control is one command aimed at one job.
type Control struct {
	Action ControlAction
	ID     string
}
