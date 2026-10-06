package cat

import (
	"context"
	"math/rand/v2"

	"github.com/chinmay-sawant/ownframe"
)

func (a *animation) nextIndex(seconds float64) int {
	step := int(max(0, seconds) / 8)
	if !a.random {
		return step % len(a.files)
	}
	if step == a.step {
		return a.current
	}
	a.step = step
	if len(a.files) < 2 {
		return a.current
	}
	return (a.current + 1 + rand.IntN(len(a.files)-1)) % len(a.files)
}

func (c *Companion) click(_ context.Context, box ownframe.Box) error {
	index := c.animation.current
	switch box.Action {
	case "latest":
		c.shown = true
		n, _ := c.inbox.Latest()
		index, _ = expressionIndex(n.Expression)
		c.animation.cycle = false
	case "dismiss":
		c.shown = false
		index = 0
		if c.RandomBehavior {
			index = rand.IntN(2)
		}
		c.animation.cycle = c.RandomBehavior
		c.animation.random = c.RandomBehavior
		c.animation.step = int(c.seconds / 8)
	default:
		return nil
	}
	if err := c.animation.load(index); err != nil {
		return err
	}
	c.view.Shown = c.shown
	c.Page.SetData(c.view)
	return nil
}
