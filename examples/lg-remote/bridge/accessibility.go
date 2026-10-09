package bridge

import "strings"

var (
	accessibility string
	actions       = []string{}
)

// SetAccessibility publishes the current visible controls from the game loop.
func SetAccessibility(value string) {
	mu.Lock()
	accessibility = value
	mu.Unlock()
}

// Accessibility is the latest JSON snapshot for Android's virtual nodes.
func Accessibility() string {
	mu.Lock()
	defer mu.Unlock()
	return accessibility
}

// QueueAction schedules a native touch or accessibility action on the game loop.
func QueueAction(action string) bool {
	mu.Lock()
	defer mu.Unlock()
	limit := 63
	if strings.HasPrefix(action, "up:") || strings.HasPrefix(action, "cancel:") {
		limit = 64
	}
	if len(actions) >= limit {
		return false
	}
	actions = append(actions, action)
	return true
}

// TakeAction removes the next accessibility action.
func TakeAction() string {
	mu.Lock()
	defer mu.Unlock()
	if len(actions) == 0 {
		return ""
	}
	action := actions[0]
	actions = actions[1:]
	return action
}
