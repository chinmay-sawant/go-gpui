package page

import "github.com/chinmay-sawant/go-gpui/internal/host"

// DevTools reports whether the window devtools overlay starts on.
func (p *Page) DevTools() bool {
	return p.devtools
}

// SetDevTools turns the window devtools overlay on or off.
func (p *Page) SetDevTools(on bool) {
	p.devtools = on
}

var _ host.Inspector = (*Page)(nil)
