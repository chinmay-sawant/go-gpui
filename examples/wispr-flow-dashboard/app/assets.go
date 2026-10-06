package app

import (
	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/insights"
)

// assetNames are the SVG files the template refers to by name.
var assetNames = []string{
	"gauge", "icon-share", "icon-info", "icon-trend",
	"icon-chevron-left", "icon-chevron-right",
	"icon-chevron-left-off", "icon-chevron-right-off", "icon-desktop",
	"icon-browser", "icon-doc", "icon-chat", "icon-mail", "icon-work",
	"logo-flow", "icon-panel", "icon-history", "icon-mic", "icon-note",
	"icon-chart", "icon-book", "icon-scissors", "icon-text", "icon-wand",
	"icon-scratch", "icon-phone", "icon-invite", "icon-gift", "icon-gear",
	"icon-help", "icon-search", "icon-play", "icon-copy", "icon-pin",
	"icon-more", "icon-plus", "icon-trash", "icon-check", "icon-close",
	"hero-art", "voice-art", "orb", "qr-code",
}

// registerImages hands the assets to the page as named image sources.
func registerImages(page *ownframe.Page) {
	for _, name := range assetNames {
		page.SetImage(name, []byte(file("assets/"+name+".svg")))
	}

	page.SetImage("share-badge", insights.ShareBadgePNG())
}
