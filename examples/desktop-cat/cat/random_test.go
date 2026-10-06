package cat

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
)

func TestRandomOptionKeepsLatestNotification(t *testing.T) {
	inbox := NewInbox()
	c, err := NewWithInbox(1, inbox)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	_ = inbox.Submit(Notification{Expression: "worried", Message: "Waiting for approval."})
	_ = c.Page.Redraw(ctx)
	if err := c.tick(ctx, 0); err != nil {
		t.Fatal(err)
	}
	_ = c.click(ctx, ownframe.Box{Action: "dismiss"})
	c.SetRandomBehavior(true)
	if !c.animation.cycle {
		t.Fatal("random option did not start idle behavior")
	}
	c.SetRandomBehavior(false)
	if c.animation.cycle {
		t.Fatal("random option did not stop idle behavior")
	}
	if err := c.click(ctx, ownframe.Box{Action: "latest"}); err != nil {
		t.Fatal(err)
	}
	if c.view.Message != "Waiting for approval." || c.animation.current != 22 {
		t.Fatal("option replaced latest notification")
	}
}
