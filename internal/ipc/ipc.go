// Package ipc passes strings between callers in this process.
// There is no socket, no second process, and no JavaScript.
package ipc

import (
	"context"
	"errors"
	"sync"
)

// ErrNoHandler means Request found no Handle for that channel.
var ErrNoHandler = errors.New("gpui: no ipc handler")

var errNilContext = errors.New("gpui: nil context")

type ear struct {
	id int64
	fn func(string)
}

type ask struct {
	id int64
	fn func(context.Context, string) (string, error)
}

var (
	mu   sync.Mutex
	seq  int64
	ears = map[string][]ear{}
	asks = map[string]ask{}
)

func noop() {}
