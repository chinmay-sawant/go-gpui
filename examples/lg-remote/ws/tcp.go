package ws

import (
	"context"
	"crypto/tls"
	"net"
	"time"
)

func dial(ctx context.Context, scheme, name, host string) (net.Conn, error) {
	d := &net.Dialer{Timeout: 5 * time.Second}
	if scheme != "wss" {
		return d.DialContext(ctx, "tcp", host)
	}

	td := tls.Dialer{
		NetDialer: d,
		Config: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         name,
		},
	}

	return td.DialContext(ctx, "tcp", host)
}
