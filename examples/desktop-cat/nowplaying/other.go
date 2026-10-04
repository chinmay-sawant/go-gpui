//go:build !windows

package nowplaying

import "context"

func open(context.Context) (reader, error) { return nil, ErrUnsupported }
