package remote

const (
	netflixSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<rect width="64" height="64" rx="10" fill="#E50914"/>` +
		`<path fill="#fff" d="M18 12h8l12 28V12h8v40h-8L26 24v28h-8z"/>` +
		`</svg>`
	primeSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#00A8E1"/>` +
		`<path fill="none" stroke="#fff" stroke-width="4" d="M16 36c6 8 26 8 32-2"/>` +
		`<path fill="#fff" d="M44 28l8 8-10 2z"/>` +
		`</svg>`
	youTubeSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<rect width="64" height="64" rx="16" fill="#FF0000"/>` +
		`<polygon points="26,18 50,32 26,46" fill="#fff"/>` +
		`</svg>`
	arrowNSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<polygon points="32,14 50,42 14,42" fill="#f4f4f5"/>` +
		`</svg>`
	arrowSSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<polygon points="14,22 50,22 32,50" fill="#f4f4f5"/>` +
		`</svg>`
	arrowWSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<polygon points="42,14 42,50 14,32" fill="#f4f4f5"/>` +
		`</svg>`
	arrowESVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<polygon points="22,14 50,32 22,50" fill="#f4f4f5"/>` +
		`</svg>`
)

func (a *App) installIcons() {
	if a.page == nil {
		return
	}

	a.page.SetImage("logo-netflix", []byte(netflixSVG))
	a.page.SetImage("logo-prime", []byte(primeSVG))
	a.page.SetImage("logo-youtube", []byte(youTubeSVG))
	a.page.SetImage("arrow-n", []byte(arrowNSVG))
	a.page.SetImage("arrow-s", []byte(arrowSSVG))
	a.page.SetImage("arrow-w", []byte(arrowWSVG))
	a.page.SetImage("arrow-e", []byte(arrowESVG))
}
