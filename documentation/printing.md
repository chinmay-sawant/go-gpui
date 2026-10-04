# Printing and PDF export

`Page.PDF` renders the page's template output to PDF bytes. The render starts
from the last template execution, the same string `Redraw` produced, with the
current theme injected as a `<style>` element in the head. It is a fresh
render through the engine, so the paper pagination can differ from the
window's layout. The live display list is not printed.

## The calls

| Call | What it does |
|------|--------------|
| `Page.PDF(ctx, opts) ([]byte, error)` | Returns the PDF bytes. |
| `Page.WritePDF(ctx, w, opts) error` | Writes the PDF to an `io.Writer`. |
| `Page.SavePDF(ctx, path, opts) error` | Writes the PDF to a file. |
| `gpui.SavePDF(ctx, page, path, opts) error` | The same, for a caller that holds the page separately. |
| `Page.Print(ctx, opts) error` | Renders to a temporary PDF and hands it to the OS print path. |

`PDFOptions` carries three fields. `PageSize` is an engine name such as
`"A4"` or `"Letter"`. `Margin` is one width in millimetres applied to all
four sides. `Profile` is a PDF conformance profile such as `"PDF/A-4"` or
`"PDF/UA-1"`. The zero value uses the engine defaults: A4, 10 mm margins, and
no profile. A profile the engine does not know is an error.

```go
data, err := page.PDF(ctx, gpui.PDFOptions{PageSize: "Letter", Margin: 20})
err = page.SavePDF(ctx, "report.pdf", gpui.PDFOptions{Profile: "PDF/A-4"})
```

## Printing

`Page.Print` writes the PDF to a temporary file and runs the platform print
helper. The helpers, tried in order:

| System | Commands |
|--------|----------|
| Linux | `lp`, then `xdg-open` |
| macOS | `osascript` asking Preview to print, then `open` |
| Windows | `powershell.exe`, then `pwsh.exe`, with `Start-Process -Verb Print` |
| wasm, Android, iOS | none, so `Print` returns `ErrNoPrinter` |

`lp` prints directly. `xdg-open` and `open` open the PDF in the default
viewer, where the print dialog is one step away. When no helper exists, the
error wraps `ErrNoPrinter` and names `SavePDF`, which is the fallback: write
the file and let the person print it. A failed hand-off removes the
temporary PDF; a successful one leaves it for the system temp cleaner,
because a viewer can read it after `Print` returns.

`GPUI_PRINT_DEBUG=1` logs each lookup, the command that ran, and each
fallback to stderr.

## Web mode

`Serve` registers `GET /pdf` when the screen implements

```go
PDF(ctx context.Context, opts gpui.PDFOptions) ([]byte, error)
```

which `*Page` does. The route returns the bytes with
`Content-Type: application/pdf` and `Cache-Control: no-store`. A screen
without that method keeps the old routes, and `/` and `/frame.png` do not
change. `examples/print` serves it on 127.0.0.1:8125.

## Limits

- The PDF re-renders from source. Pagination, page breaks, and anything the
  print path does not implement can differ from the window.
- No header or footer option in `PDFOptions` yet.
- The display list is not printed. A page drawn only from replay still
  renders its source again for the PDF.
- Relative resources resolve against nothing, because the source goes to the
  engine as in-memory HTML. An image registered with `SetImage` does not
  appear in the PDF.
