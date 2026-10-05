package app

import "github.com/chinmay-sawant/go-gpui"

// assetNames are the SVG files the template refers to by name. Each name
// exists in assets/ for the light theme and in assets_dark/ for the dark
// theme.
var assetNames = []string{
	"rail-activity", "rail-chat", "rail-teams", "rail-calendar",
	"rail-calls", "rail-files", "rail-more",
	"icon-search", "icon-filter", "icon-compose", "icon-more",
	"icon-close", "icon-check",
	"icon-chevron-down", "icon-chevron-right", "icon-chevron-left",
	"icon-plus", "icon-meet", "icon-call", "icon-mic", "icon-video",
	"icon-share-screen", "icon-send", "icon-attach", "icon-emoji",
	"icon-reply", "icon-edit", "icon-trash", "icon-pin", "icon-clock",
	"icon-people", "icon-person-add", "icon-doc", "icon-download",
	"icon-upload", "icon-star", "icon-info", "icon-bell", "icon-loop",
	"icon-grid", "icon-list", "icon-camera", "icon-gif", "icon-sticker",
	"icon-folder", "icon-calendar-plus", "icon-location", "icon-link",
	"icon-copy", "icon-voicemail", "icon-keypad", "icon-mail",
	"icon-gear", "icon-share", "icon-bookmark",
}

// registerImages hands the assets to the page as named image sources. The
// dark set strokes the icons light, so they stay visible on dark surfaces.
func registerImages(page *gpui.Page, dark bool) {
	dir := "assets/"

	if dark {
		dir = "assets_dark/"
	}

	for _, name := range assetNames {
		if data := file(dir + name + ".svg"); data != "" {
			page.SetImage(name, []byte(data))
		}
	}
}
