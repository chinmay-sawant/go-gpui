package ws

import "io"

// WriteText sends one masked text frame.
func (c *Conn) WriteText(p []byte) error {
	c.wmu <- struct{}{}
	defer func() { <-c.wmu }()

	return writeFrame(c.c, 0x1, p, true)
}

// ReadText returns the next text payload. Ping frames get a pong.
func (c *Conn) ReadText() ([]byte, error) {
	var buf []byte

	for {
		fin, op, payload, err := readFrame(c.r)
		if err != nil {
			return nil, err
		}

		switch op {
		case 0x1, 0x0:
			buf = append(buf, payload...)
			if fin {
				return buf, nil
			}
		case 0x8:
			return nil, io.EOF
		case 0x9:
			if err := c.writeCtrl(0xA, payload); err != nil {
				return nil, err
			}
		default:
		}
	}
}

// Close sends a close frame and closes the socket.
func (c *Conn) Close() error {
	_ = c.writeCtrl(0x8, nil)

	return c.c.Close()
}

func (c *Conn) writeCtrl(op byte, p []byte) error {
	c.wmu <- struct{}{}
	defer func() { <-c.wmu }()

	return writeFrame(c.c, op, p, true)
}
