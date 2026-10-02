package history

import "strings"

// toolbar is the button strip every page carries, including the routed
// pages, so Back and Forward stay reachable after a route loads.
const toolbar = `<div id="red" class="bar" data-action="red">Red route</div>` +
	`<div id="green" class="bar" data-action="green">Green route</div>` +
	`<div id="native" class="bar">Native load</div>` +
	`<div id="back" class="bar">Back</div>` +
	`<div id="forward" class="bar">Forward</div>`

// nativeHTML is the page the #native box loads through Load.
var nativeHTML = routePage("BLUE NATIVE", "#1a56db")

// routePage builds a full page with a colored marker and the toolbar.
func routePage(marker, color string) string {
	return `<!DOCTYPE html><html><head><meta charset="utf-8"><title>History</title><style>` +
		`html,body{height:100%;margin:0}body{background:#f4f1ea;font-family:sans-serif;color:#1c1915}` +
		`.card{width:380px;background:#ffffff;padding:24px;margin:24px}` +
		`h1{font-size:22px;margin:0 0 8px}` +
		`#page{font-size:18px;color:` + color + `;margin:0 0 8px}` +
		`.bar{display:block;background:#e5e7eb;border:1px solid #c8c2b4;padding:8px;margin-top:8px;text-align:center}` +
		`</style></head><body><div id="card" class="card"><h1>History</h1>` +
		`<p id="page">` + marker + `</p>` + toolbar +
		`</div></body></html>`
}

// markerIn names the page a source shows, for the status line.
func markerIn(source string) string {
	for _, marker := range []string{"RED", "GREEN", "BLUE NATIVE"} {
		if strings.Contains(source, marker) {
			return marker
		}
	}

	return "initial"
}
