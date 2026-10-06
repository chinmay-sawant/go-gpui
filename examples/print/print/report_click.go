package print

import (
	"context"

	"github.com/chinmay-sawant/ownframe"
)

// onClick saves or prints the report. A failure becomes the status line
// instead of an error, so the page still redraws.
func (a *App) onClick(ctx context.Context, box ownframe.Box) error {
	switch box.ID {
	case "save":
		a.save(ctx)
	case "print":
		a.print(ctx)
	}

	return nil
}

func (a *App) save(ctx context.Context) {
	if err := a.page.SavePDF(ctx, a.SavePath, ownframe.PDFOptions{}); err != nil {
		a.view.Status = "Save failed: " + err.Error()

		return
	}

	a.view.Status = "Saved to " + a.SavePath
}

func (a *App) print(ctx context.Context) {
	if err := a.page.Print(ctx, ownframe.PDFOptions{}); err != nil {
		a.view.Status = "Print failed: " + err.Error()

		return
	}

	a.view.Status = "Sent to the printer."
}
