package nowplaying

import (
	"context"
	"strings"
	"time"
	"unsafe"
)

func readSession(ctx context.Context, s *object) (Track, error) {
	app, err := s.text(6)
	if err != nil {
		return Track{}, err
	}
	name := strings.ToLower(app)
	if !strings.Contains(name, "chrome") && !strings.Contains(name, "chromium") {
		return Track{}, nil
	}
	info, err := s.get(9)
	if err != nil {
		return Track{}, err
	}
	defer info.release()
	var status int32
	if err := info.call(7, uintptr(unsafe.Pointer(&status))); err != nil {
		return Track{}, err
	}
	if status != 4 {
		return Track{}, nil
	}
	op, err := s.get(7)
	if err != nil {
		return Track{}, err
	}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	properties, err := await(ctx, op)
	if err != nil {
		return Track{}, err
	}
	defer properties.release()
	title, err := properties.text(6)
	if err != nil {
		return Track{}, err
	}
	artist, err := properties.text(9)
	if err != nil {
		return Track{}, err
	}
	return Track{Title: strings.TrimSpace(title), Artist: artist, App: app, Playing: true}, nil
}
