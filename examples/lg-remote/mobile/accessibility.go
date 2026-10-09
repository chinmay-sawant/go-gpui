package mobile

import "github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"

// Accessibility returns the visible controls as JSON for Android TalkBack.
func Accessibility() string { return bridge.Accessibility() }

// QueueAction puts a native accessibility action on the game loop.
func QueueAction(action string) bool { return bridge.QueueAction(action) }
