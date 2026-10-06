package crash

import (
	"context"

	"github.com/chinmay-sawant/ownframe"
)

// Report writes a report for reason and returns its path.
func (a *App) Report(reason string) (string, error) {
	return ownframe.Report(reportTitle, reason)
}

// onClick writes a manual report, or panics so Run can write one.
// The panic unwinds through ownframe.Run, which recovers, writes the report,
// and returns an error naming the file. This handler does not catch it.
func (a *App) onClick(_ context.Context, box ownframe.Box) error {
	switch box.ID {
	case "report":
		path, err := a.Report("manual report")
		if err != nil {
			return err
		}

		a.view.Status = path
		a.page.SetData(a.view)
	case "panic":
		panic("crash example: panic button")
	}

	return nil
}
