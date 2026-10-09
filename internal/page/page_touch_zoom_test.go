package page

import "testing"

func TestTouchZoomConfigAndOverride(t *testing.T) {
	p, err := New(Config{HTML: "<main>ok</main>", Width: 100, Height: 100})
	if err != nil {
		t.Fatal(err)
	}
	if !p.TouchZoomAllowed() {
		t.Fatal("touch zoom should remain enabled by default")
	}
	p.SetTouchZoomAllowed(false)
	if p.TouchZoomAllowed() {
		t.Fatal("runtime override did not disable touch zoom")
	}
	p, err = New(Config{HTML: "<main>ok</main>", Width: 100, Height: 100, DisableTouchZoom: true})
	if err != nil {
		t.Fatal(err)
	}
	if p.TouchZoomAllowed() {
		t.Fatal("config did not disable touch zoom")
	}
}
