package tv

import (
	"errors"
	"fmt"
	"time"
)

func (c *Client) writePointer(body string) error {
	if c == nil {
		return errors.New("the pointer socket is closed")
	}

	if c.pointerStale() {
		c.closePointer()
	}

	if c.in != nil {
		if err := c.in.WriteText([]byte(body)); err == nil {
			c.pointerAt = time.Now()
			return nil
		}

		c.closePointer()
		time.Sleep(time.Second)
	}

	c.openPointer()
	if c.in == nil {
		return errors.New("the pointer socket is closed")
	}

	err := c.in.WriteText([]byte(body))
	if err == nil {
		c.pointerAt = time.Now()
	}

	return err
}

// pointerStale is true when the pointer socket has sat idle for 2 seconds.
// The TV then ignores arrow keys until the socket is opened again.
func (c *Client) pointerStale() bool {
	if c.in == nil || c.pointerAt.IsZero() {
		return false
	}

	return time.Since(c.pointerAt) >= 2*time.Second
}

func (c *Client) closePointer() {
	if c.in == nil {
		return
	}

	_ = c.in.Close()
	c.in = nil
}

func (c *Client) pointerMove(spec string) (string, error) {
	dx, dy := 0, 0
	_, _ = fmt.Sscanf(spec, "%d,%d", &dx, &dy)
	body := fmt.Sprintf("type:move\ndx:%d\ndy:%d\ndown:0\n\n", dx, dy)

	return "Pointer", c.writePointer(body)
}

func (c *Client) pointerScroll(spec string) (string, error) {
	dy := 0
	_, _ = fmt.Sscanf(spec, "%d", &dy)
	body := fmt.Sprintf("type:scroll\ndx:0\ndy:%d\n\n", dy)

	return "Scroll", c.writePointer(body)
}
