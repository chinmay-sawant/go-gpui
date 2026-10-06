package telegram

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/ownframe"
)

// onLongPress opens the reaction bar on the long-pressed message. Any
// other box does nothing.
func (a *App) onLongPress(_ context.Context, box ownframe.Box) error {
	if !strings.HasPrefix(box.ID, "msg-") {
		return nil
	}

	id := strings.TrimPrefix(box.ID, "msg-")
	if a.findMessage(id) == nil {
		return nil
	}

	a.view.ReactID = id
	a.rebuild()
	a.page.SetData(&a.view)

	return nil
}

// findMessage returns the open thread's message with that id, or nil.
func (a *App) findMessage(id string) *Message {
	thread := a.threads[a.view.Active]

	for i := range thread {
		if thread[i].ID == id {
			return &thread[i]
		}
	}

	return nil
}

// react puts the named reaction on the message the bar was opened for and
// closes the bar.
func (a *App) react(name string) {
	msg := a.findMessage(a.view.ReactID)
	if msg == nil {
		return
	}

	for _, r := range a.view.Reactions {
		if r.Name == name {
			msg.Reaction = r.Emoji
			break
		}
	}

	a.view.ReactID = ""
}
