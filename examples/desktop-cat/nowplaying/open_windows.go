package nowplaying

import (
	"context"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

type mediaReader struct{ manager *object }

func open(ctx context.Context) (reader, error) {
	code, _, _ := combase.NewProc("RoInitialize").Call(1)
	if err := result(code); err != nil {
		return nil, err
	}
	r := &mediaReader{}
	success := false
	defer func() {
		if !success {
			r.close()
		}
	}()
	class, err := hstring("Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager")
	if err != nil {
		return nil, err
	}
	defer deleteString.Call(class)
	guid, _ := windows.GUIDFromString("{2050C4EE-11A0-57DE-AED7-C97C70338245}")
	var factory *object
	code, _, _ = combase.NewProc("RoGetActivationFactory").Call(class, uintptr(unsafe.Pointer(&guid)), uintptr(unsafe.Pointer(&factory)))
	if err := result(code); err != nil {
		return nil, err
	}
	defer factory.release()
	op, err := factory.get(6)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r.manager, err = await(ctx, op)
	if err != nil {
		return nil, err
	}
	success = true
	return r, nil
}

func (r *mediaReader) close() {
	r.manager.release()
	combase.NewProc("RoUninitialize").Call()
}
