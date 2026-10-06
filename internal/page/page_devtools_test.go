package page_test

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/host"
	"github.com/chinmay-sawant/ownframe/internal/page"
)

// fakeInspector is a screen double that satisfies host.Inspector.
type fakeInspector struct {
	on    bool
	stats host.Stats
}

func (f *fakeInspector) DevTools() bool      { return f.on }
func (f *fakeInspector) SetDevTools(on bool) { f.on = on }
func (f *fakeInspector) Stats() host.Stats   { return f.stats }

var _ host.Inspector = (*fakeInspector)(nil)

func TestDevToolsDefaultsOff(t *testing.T) {
	t.Parallel()

	p := newCachePage(t, `<p>static</p>`)
	if p.DevTools() {
		t.Fatal("devtools on by default")
	}

	p.SetDevTools(true)
	if !p.DevTools() {
		t.Fatal("SetDevTools(true) did not turn the overlay on")
	}

	p.SetDevTools(false)
	if p.DevTools() {
		t.Fatal("SetDevTools(false) did not turn the overlay off")
	}
}

func TestConfigDevToolsStartsOn(t *testing.T) {
	t.Parallel()

	p, err := page.New(page.Config{
		HTML:     `<p>static</p>`,
		Width:    320,
		Height:   200,
		DevTools: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if !p.DevTools() {
		t.Fatal("Config.DevTools did not start the overlay")
	}
}
