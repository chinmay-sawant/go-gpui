package insights

import "github.com/chinmay-sawant/go-gpui"

// assetNames are the SVG files the template refers to by name.
var assetNames = []string{
	"gauge", "icon-share", "icon-info", "icon-trend",
	"icon-chevron-left", "icon-chevron-right",
	"icon-chevron-left-off", "icon-chevron-right-off", "icon-desktop",
	"icon-browser", "icon-doc", "icon-chat", "icon-mail", "icon-work",
}

// registerImages hands the assets to the page as named image sources.
func registerImages(page *gpui.Page) {
	for _, name := range assetNames {
		page.SetImage(name, []byte(file("assets/"+name+".svg")))
	}

	page.SetImage("share-badge", shareBadgePNG())
}
