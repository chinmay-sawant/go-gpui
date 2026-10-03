package page

// SetImage registers encoded image bytes (PNG, JPEG, or SVG) for one source
// name. A template image source that spells that name, such as
// background-image: url("logo") or <img src="logo">, resolves to data on the
// next Redraw. Passing nil data removes the entry.
func (p *Page) SetImage(src string, data []byte) {
	if p.images == nil {
		p.images = map[string][]byte{}
	}

	if data == nil {
		delete(p.images, src)

		return
	}

	p.images[src] = data
}

// imageBytes resolves one image source for the layout engine.
func (p *Page) imageBytes(src string) ([]byte, error) {
	if data, ok := p.images[src]; ok {
		return data, nil
	}

	return nil, errNoImage
}
