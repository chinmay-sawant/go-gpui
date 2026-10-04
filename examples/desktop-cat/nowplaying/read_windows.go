package nowplaying

import (
	"context"
	"unsafe"
)

func (r *mediaReader) read(ctx context.Context) (Track, error) {
	sessions, err := r.manager.get(7)
	if err != nil {
		return Track{}, err
	}
	defer sessions.release()
	var count uint32
	if err := sessions.call(7, uintptr(unsafe.Pointer(&count))); err != nil {
		return Track{}, err
	}
	for i := uint32(0); i < count; i++ {
		var session *object
		if err := sessions.call(6, uintptr(i), uintptr(unsafe.Pointer(&session))); err != nil {
			return Track{}, err
		}
		track, err := readSession(ctx, session)
		session.release()
		if err != nil {
			return Track{}, err
		}
		if track.Playing && track.Title != "" {
			return track, nil
		}
	}
	return Track{}, nil
}
