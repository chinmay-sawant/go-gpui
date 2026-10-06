package page

import (
	"context"
	"fmt"
	"os"

	"github.com/chinmay-sawant/ownframe/internal/print"
)

// Print renders the last template output to a temporary PDF and hands it to
// the OS print path. A failed hand-off removes the temporary file; a
// successful one leaves it for the system temp cleaner, because a viewer
// may read it after this call returns. Without a print helper the error
// wraps ErrNoPrinter and names SavePDF.
func (p *Page) Print(ctx context.Context, opts PDFOptions) error {
	if !print.Available() {
		return fmt.Errorf("ownframe: print: %w; use SavePDF and print the file yourself", ErrNoPrinter)
	}

	file, err := os.CreateTemp("", "ownframe-print-*.pdf")
	if err != nil {
		return err
	}

	path := file.Name()
	handed := false

	defer func() {
		if !handed {
			os.Remove(path)
		}
	}()

	if err := p.WritePDF(ctx, file, opts); err != nil {
		file.Close()

		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := print.Print(ctx, path); err != nil {
		return fmt.Errorf("ownframe: print: %w; use SavePDF and print the file yourself", err)
	}

	handed = true

	return nil
}
