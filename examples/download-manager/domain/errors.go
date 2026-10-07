package domain

import "fmt"

// TransitionError reports a refused state change.
type TransitionError struct {
	From State
	To   State
}

func (e *TransitionError) Error() string {
	return fmt.Sprintf("download-manager: cannot move job from %s to %s", e.From, e.To)
}
