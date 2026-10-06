package cat

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
)

func TestSilhouetteAndDismissBehavior(t *testing.T) {
	inbox := NewInbox()
	c, err := NewWithInbox(1, inbox)
	if err != nil {
		t.Fatal(err)
	}
	c.RandomBehavior = true
	ctx := context.Background()
	if err := c.Page.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	if err := c.tick(ctx, 1); err != nil {
		t.Fatal(err)
	}
	for _, p := range [][2]int{{100, 142}, {101, 143}, {299, 341}, {110, 320}} {
		if c.Draggable(p[0], p[1]) {
			t.Fatalf("transparent image corner reserved input: %v", p)
		}
	}
	if !c.Draggable(200, 240) {
		t.Fatal("cat body not clickable")
	}

	if err := inbox.Submit(Notification{Expression: "sad", Message: "Waiting for your input."}); err != nil {
		t.Fatal(err)
	}
	if err := c.tick(ctx, 2); err != nil {
		t.Fatal(err)
	}
	if err := c.click(ctx, ownframe.Box{Action: "dismiss"}); err != nil {
		t.Fatal(err)
	}
	if c.animation.current > 1 || c.shown || !c.animation.cycle {
		t.Fatal("dismiss did not choose happy idle behavior")
	}
	if err := c.Page.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	happy := c.animation.current
	if err := c.tick(ctx, 3); err != nil {
		t.Fatal(err)
	}
	if c.animation.current != happy {
		t.Fatal("happy face disappeared immediately")
	}
	if err := c.tick(ctx, 8); err != nil {
		t.Fatal(err)
	}
	if c.animation.current == happy {
		t.Fatal("random idle expression did not change")
	}
	if err := c.click(ctx, ownframe.Box{Action: "latest"}); err != nil {
		t.Fatal(err)
	}
	if err := c.Page.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	if err := c.tick(ctx, 9); err != nil {
		t.Fatal(err)
	}
	if !c.shown || c.animation.current != 18 || c.animation.cycle {
		t.Fatal("latest click did not restore notification")
	}
}
