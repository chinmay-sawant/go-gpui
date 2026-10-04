package print

import (
	"os"
	"path/filepath"

	"github.com/chinmay-sawant/go-gpui"
)

// New parses the embedded report and registers the click handler. webMode
// adds the /pdf hint to the page.
func New(webMode bool) (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "Quarterly report",
		HTML:   reportHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page, SavePath: defaultSavePath()}
	app.view = View{Title: "Quarterly report", Web: webMode, Status: "Ready."}

	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(&app.view)

	return app, nil
}

// defaultSavePath is gpui-print-report.pdf under the user home directory.
func defaultSavePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "gpui-print-report.pdf"
	}

	return filepath.Join(home, "gpui-print-report.pdf")
}
