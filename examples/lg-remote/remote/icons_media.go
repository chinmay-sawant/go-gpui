package remote

const (
	playSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<polygon points="26,18 50,32 26,46" fill="#f4f4f5"/>` +
		`</svg>`
	pauseSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<rect x="22" y="18" width="7" height="28" fill="#f4f4f5"/>` +
		`<rect x="35" y="18" width="7" height="28" fill="#f4f4f5"/>` +
		`</svg>`
	stopSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">` +
		`<circle cx="32" cy="32" r="30" fill="#2a2a30"/>` +
		`<rect x="20" y="20" width="24" height="24" rx="2" fill="#f4f4f5"/>` +
		`</svg>`
)

func (a *App) installMedia() {
	if a.page == nil {
		return
	}

	a.page.SetImage("icon-play", []byte(playSVG))
	a.page.SetImage("icon-pause", []byte(pauseSVG))
	a.page.SetImage("icon-stop", []byte(stopSVG))
}
