# Forms

The template may contain `input`, `textarea`, and `select`. Each control needs an id. A control with no id is ignored. The library stores the value. `examples/forms` reads the stored values when the send control is clicked.

```go
page.Handle(gpui.Handlers{
    Click: func(_ context.Context, box gpui.Box) error {
        if box.Action != "send" {
            return nil
        }
        email := page.FormValue("email")
        page.SetData(struct{ Status string }{Status: email})
        return nil
    },
})
```

`Click` calls that handler, then draws the page again. A click on a control updates it before the handler runs. `FormValue` returns the stored string. `FormChecked` reports a checkbox or a radio. `FocusedField` is the id of the focused control, or empty when none is focused.

Supported `input` types:

- `text`, and an `input` with no type
- `password`, `email`, `search`, `tel`, `url`, and `number`
- `file`
- `checkbox` and `radio`

Any other `input` type is left alone.

Those text-like types, including `file`, take typed text. Click one to focus it. Click a textarea to focus it too, and typing edits that text. Click outside a control to blur. `Type`, `Backspace`, `DeleteWord`, `Paste`, `SelectAll`, `Copy`, and `Cut` edit the focused text-like control, or the focused textarea, even when those handlers are nil. A password stores plaintext. The picture shows one bullet for each rune. `Copy` returns the plaintext.

Click a checkbox to toggle it. Click a radio to check it and uncheck the other radios with the same name. Click a select to cycle to the next option. Those clicks focus the control. A disabled control does not toggle, focus, or take typing. The click handler still runs.

A file input has no dialog. The user types the file name. An empty file field paints the words "No file".

Before paint, a text-like input is rewritten to a `span`. That includes `text`, `password`, and `file`. The layout engine does not paint an input value. Checkbox and radio stay `input` elements, and the engine paints the mark. A textarea stays a `textarea`. A select stays a `select` and paints only the selected label.

`Load`, `Back`, and `Forward` reset form state, because the document changed. `SetFormValue` and `SetFormChecked` do not redraw. Call `Redraw` after them. `SetFormValue` does nothing before that control has been drawn. `SetFormChecked` changes a checkbox or a radio, and it does nothing when that id is missing.

There is still no native file dialog. The login example still uses its own div fields.
