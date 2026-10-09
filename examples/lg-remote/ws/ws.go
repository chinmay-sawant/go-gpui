// Package ws is a small WebSocket client for the TV.
// The handshake sends no Origin header. webOS rejects some browser origins,
// and a desktop or phone client is not a page.
package ws

import "errors"

// ErrHandshake means the server did not accept the upgrade.
var ErrHandshake = errors.New("ws: handshake was refused")
