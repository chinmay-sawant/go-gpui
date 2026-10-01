# Screen paint

`Redraw` in `internal/page/page_draw.go` does two steps.

1. It executes the page template with `Page` data.
2. It calls `render.Paint` with that HTML and the current width and height.

`internal/render/paint.go` then calls `html.Parse`, `css.Apply`, and `layout.Lay`. CSS options are the frame size, media `screen`, and no extra sheet. `Paint` returns `placed.Image()` and `placed.Boxes()`.

`layout.Lay` paints through `imageout.RenderLayout`. It returns an `image.Image`. go-gpui does not call `Document.WritePDF`, `Document.PDF`, or `ImageDocument`.

The library that does the layout is still named gowkhtmltopdf. Its image painter uses `pdf.Font` and `pdf.Registry` as font tables. Those types are not a PDF file, and this window does not rasterize one.

`Page.PNG` runs `png.Encode` on the image already stored. The `-web` host needs those bytes for `GET /frame.png`. The Ebiten window, the wasm canvas, and the phone bind draw `Page.Image` directly.

A blank template never reaches `Paint`. `New` rejects it with `ErrEmptyHTML`.
