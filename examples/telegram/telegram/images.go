package telegram

import "github.com/chinmay-sawant/go-gpui"

// iconNames are the SVG files the template refers to by name.
var iconNames = []string{
	"icon-back", "icon-search", "icon-send", "icon-check", "icon-checks",
	"icon-pin", "icon-mute", "icon-attach", "icon-gift",
	"icon-camera", "icon-gallery",
}

// registerImages hands the icons to the page as named image sources.
func registerImages(page *gpui.Page) {
	for _, name := range iconNames {
		if data := file("icons/" + name + ".svg"); data != "" {
			page.SetImage(name, []byte(data))
		}
	}
}
