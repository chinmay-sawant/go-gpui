package player

import (
	"strconv"
	"strings"

	"github.com/chinmay-sawant/ownframe"
)

// registerImages hands every embedded SVG to the page by base name, then
// points the fixed cover slots at the six placeholder covers.
func registerImages(page *ownframe.Page) {
	entries, err := files.ReadDir("assets")
	if err != nil {
		return
	}

	for _, entry := range entries {
		name := strings.TrimSuffix(entry.Name(), ".svg")
		if name == entry.Name() {
			continue
		}

		page.SetImage(name, []byte(file("assets/"+entry.Name())))
	}

	for i := 0; i < 6; i++ {
		cover := "cover-" + strconv.Itoa(i+1)
		body := []byte(file("assets/" + cover + ".svg"))
		page.SetImage("track-"+strconv.Itoa(i), body)
		page.SetImage("pick-"+strconv.Itoa(i), body)
		page.SetImage("card-"+strconv.Itoa(i), body)
	}
}
