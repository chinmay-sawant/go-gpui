package tv

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/ws"
)

func (c *Client) openPointer() {
	p, err := c.Request("com.webos.service.networkinput/getPointerInputSocket", nil)
	if err != nil {
		return
	}

	path, _ := p["socketPath"].(string)
	if path == "" {
		return
	}

	conn, err := ws.Dial(context.Background(), path)
	if err != nil {
		return
	}

	c.in = conn
}

func (c *Client) readInfo() {
	if c.Model == "" {
		if p, err := c.Request("system/getSystemInfo", nil); err == nil {
			c.Model = textField(p, "modelName")
		}
	}

	p, err := c.Request("com.webos.service.connectionmanager/getinfo", nil)
	if err != nil {
		return
	}

	c.MACs = macsOf(p)
}

func textField(p map[string]any, key string) string {
	s, _ := p[key].(string)

	return s
}
