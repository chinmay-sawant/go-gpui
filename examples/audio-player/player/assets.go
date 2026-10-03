package player

import (
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
)

// assetNames are the SVG files the template refers to by name.
var assetNames = []string{
	"chevron-down", "clock", "cover-1", "cover-2", "cover-3", "cover-4",
	"devices", "heart-fill", "heart", "home", "home-on", "library",
	"library-on", "logo-wave", "more", "next", "note", "pause", "play",
	"plus", "prev", "queue", "repeat", "search", "search-light", "search-on",
	"shuffle", "shuffle-on", "volume-mute", "volume", "repeat-on",
}

// registerImages hands the assets to the page as named image sources.
// cover-0..cover-5 alias the four placeholders, so a live slot whose
// artwork fetch failed still paints a cover.
func registerImages(page *gpui.Page) {
	for _, name := range assetNames {
		page.SetImage(name, []byte(file("assets/"+name+".svg")))
	}

	for i := 0; i < 6; i++ {
		src := fmt.Sprintf("assets/cover-%d.svg", i%4+1)
		page.SetImage(fmt.Sprintf("cover-%d", i), []byte(file(src)))
	}
}
