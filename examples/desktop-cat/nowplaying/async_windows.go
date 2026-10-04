package nowplaying

import (
	"context"
	"fmt"
	"time"
	"unsafe"
)

func await(ctx context.Context, op *object) (*object, error) {
	defer op.release()
	info, err := op.query("{00000036-0000-0000-C000-000000000046}")
	if err != nil {
		return nil, err
	}
	defer info.release()
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	for {
		var status int32
		if err := info.call(7, uintptr(unsafe.Pointer(&status))); err != nil {
			return nil, err
		}
		switch status {
		case 1:
			return op.get(8)
		case 2:
			return nil, context.Canceled
		case 3:
			var code uint32
			if err := info.call(8, uintptr(unsafe.Pointer(&code))); err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("Windows media async error 0x%08X", code)
		}
		select {
		case <-ctx.Done():
			_ = info.call(9)
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}
