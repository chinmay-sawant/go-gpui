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

`Click` calls that handler, then draws the page again. A click on a control updates it before the handler runs. `FormValue` returns the stored string. `FormChecked` reports a checkbox or a radio. `FormSelected` reports whether a control is focused and its whole value is selected. `FocusedField` is the id of the focused control, or empty when none is focused.

Supported `input` types:

- `text`, and an `input` with no type
- `password`, `email`, `search`, `tel`, `url`, and `number`
- `file`
- `checkbox` and `radio`

Any other `input` type is left alone.

Those text-like types, including `file`, take typed text. Click one to focus it. Click a textarea to focus it too, and typing edits that text. Click outside a control to blur. `Type`, `Backspace`, `DeleteWord`, `Paste`, `SelectAll`, `Copy`, and `Cut` edit the focused text-like control, or the focused textarea, even when those handlers are nil. `SelectAll` selects the whole value, so `FormSelected` is true until the next edit or blur. A password stores plaintext. The picture shows one bullet for each rune. `Copy` returns the plaintext.

Click a checkbox to toggle it. Click a radio to check it and uncheck the other radios with the same name. Click a select to cycle to the next option. Those clicks focus the control. A disabled control does not toggle, focus, or take typing. The click handler still runs.

A file input has no dialog. The user types the file name. An empty file field paints the words "No file".

Before paint, a text-like input or a textarea is rewritten to a `span`. The layout engine does not paint an input value. Checkbox and radio stay `input` elements, and the engine paints the mark. A select stays a `select` and paints only the selected label. The rewrite keeps the author's attributes, such as `class`, `style`, `placeholder`, `name`, `autocomplete`, `maxlength`, `readonly`, `required`, `aria-*`, and `data-*`. The span drops `value` and `type`, because the stored text carries the state, and the rewrite drops any `data-gpui-*` that the library manages. Library state appears as attributes: `data-gpui-field="input"`, `data-gpui-field="textarea"`, or `data-gpui-field="select"`, plus `data-gpui-focus="1"`, `data-gpui-selected="1"`, and `data-gpui-placeholder="1"` when they apply. A focused text field or textarea that is not selected carries an empty caret span, `<span data-gpui-caret="1"></span>`. A password paints one bullet for each rune. When the value is empty and the author set `placeholder`, the field paints the placeholder text. A file input still paints "No file". The default look is a `<style>` block. If an opening `<head ...>` tag is found, it is inserted immediately after it and before author CSS. Otherwise, it goes immediately before the first `</head>` tag if one exists; if neither tag is found, it is prepended to the document. The textarea rule targets its rewritten span with `[data-gpui-field="textarea"]` and sets `white-space: pre-wrap`. An author base rule of equal specificity wins because it comes later, while a state rule has higher specificity than an author base rule. A `<button>` is not a form control. A page with a button also gets `button{display:inline-block;padding:4px 12px;border:1px solid #c8c2b4;background:#f0f0f0;color:#1c1915;text-align:center}` from the same stylesheet, before author CSS, so a bare button lays out with a hit box. An `input` with `type=submit`, `type=button`, or `type=reset` is rewritten to a `button` element with the same attributes, and its `value` becomes the label.

A control may carry `data-bind="Field"` to tie it to a field on the struct passed to `SetData`, which must be a pointer. An edit writes the new value into the field before the page is drawn, and `SetData` followed by `Redraw` pushes the field's value back into the control. `Handlers.Change` runs after a bound control changes and receives that control's box; a nil function ignores the event. The full contract, including the field types, is in [binding.md](binding.md).

`Load`, `Back`, and `Forward` reset form state, because the document changed. `SetFormValue` and `SetFormChecked` do not redraw. Call `Redraw` after them. `SetFormValue` does nothing before that control has been drawn. `SetFormChecked` changes a checkbox or a radio, and it does nothing when that id is missing.

There is still no native file dialog.
