package cat

import (
	"context"
	"time"
)

func (c *Companion) receive() bool {
	n, v := c.inbox.Latest()
	if v == c.version {
		return false
	}
	c.version = v
	c.shown = true
	c.view = bubbleView{Shown: true, Message: n.Message, Source: n.Source, Time: time.Now().Format("3:04 pm")}
	c.Page.SetData(c.view)
	if v > 1 {
		index, _ := expressionIndex(n.Expression)
		_ = c.animation.load(index)
		c.animation.cycle = false
	}
	return true
}

func (c *Companion) tick(ctx context.Context, seconds float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if c.receive() {
		if err := c.Page.Redraw(ctx); err != nil {
			return err
		}
	}
	c.seconds = seconds
	c.animation.random = c.RandomBehavior
	if c.ReducedMotion {
		return nil
	}
	return c.animation.paint(ctx, seconds)
}

// Interactive reserves opaque cat pixels and the visible bubble.
// Coordinates are window-local CSS pixels; empty margins pass through.
func (c *Companion) Interactive(x, y int) bool {
	return c.Draggable(x, y) ||
		c.shown && x >= 12 && x < 308 && y >= 12 && y < 132
}
