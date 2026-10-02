# Screen paint

`Redraw` in `internal/page/page_draw.go` does two steps.

1. It executes the page template with `Page` data.
2. It calls `render.Paint` with that HTML and the current width and height.

`internal/render/paint.go` then calls `html.Parse`, `css.Apply`, and `layout.Lay`. CSS options are the frame size, media `screen`, and no extra sheet. `Paint` returns `placed.Image()` and `placed.Boxes()`.

`layout.Lay` paints through `imageout.RenderLayout`. It returns an `image.Image`. go-gpui does not call `Document.WritePDF`, `Document.PDF`, or `ImageDocument`.

The library that does the layout is still named gowkhtmltopdf. Its image painter uses `pdf.Font` and `pdf.Registry` as font tables. Those types are not a PDF file, and this window does not rasterize one.

`Page.PNG` runs `png.Encode` on the image already stored. The `-web` host needs those bytes for `GET /frame.png`. The Ebiten window, the wasm canvas, and the phone bind draw `Page.Image` directly.

A blank template never reaches `Paint`. `New` rejects it with `ErrEmptyHTML`.

`internal/render/display.go` is a second entry over that same placement. `DisplayList` takes the same arguments as `Paint`: an HTML string, a width, and a height in CSS pixels. It runs `html.Parse` and `css.Apply` with the same options, then calls `layout.DisplayList` instead of `layout.Lay`. `layout.DisplayList` stops before the paint step, so the result carries no `image.Image`.

`Display` holds `Ops`, the operations in source order, and `Order`, the same operations as indices in paint order. `Order` is the order to iterate, because it applies z-index and the outline paint layer. `Width` and `Height` are the canvas in CSS pixels. `PointsPerPixel` and `PixelPerPoint` convert between the op coordinate space and that canvas. Op coordinates are points with y down, and for `OpText` and `OpBullet` the `Y` field is the baseline. `Width` and `Height` come from the same placement `Paint` builds, but the engine converts them from points rather than reading them off a picture, so the height can be one pixel under the painted image height.

The kinds are `OpFillRect`, `OpStrokeRect`, `OpLine`, `OpText`, `OpImage`, `OpLinkURI`, `OpBullet`, `OpGridRun`, `OpUnknown`, and `OpNoop`. An `OpGridRun` carries one table row's collapsed border grid in `Grid.Segs`, replayed in order. Two kinds paint nothing and must be skipped: `OpNoop`, left behind when overflow clipping deactivates an operation, because the box tree stores operation indices that cannot shift, and `OpUnknown`, which the engine emits only as the boundary marker of a blend or isolation group. Treat any kind outside the list as inert, so a kind added later cannot be mistaken for a fill.

`DisplayOp` is the engine's own operation type under a local name, so callers of `render` do not import the engine to name it. Read a rare payload through its accessor methods, `LinkURI`, `ImageBytes`, `ImageAlt`, `Transform`, `BlendModeName`, `Opacity`, `Outline`, `FontFeatures`, `TextLanguage`, `TextAutospace`, `TextTransformValue`, and `NoFakeBoldValue`. The plain fields such as `Kind`, `X`, `Y`, `W`, `H`, `Text`, and `Font` are always safe. Every operation the engine emits today carries its payload, but that payload sits behind an embedded pointer, and a caller outside the engine cannot build or inspect it. The accessors are nil-safe, so they are the supported read path. A blend group is read through `Group`, `GroupBoundary`, `IsGroupBegin`, and `IsGroupEnd`, which are nil-safe too. `op.Font.Bytes()` returns the raw font face so a caller can shape text with its own shaper.

Nothing in this repository draws those operations. `internal/window` still blits `Page.Image`, `internal/web` still encodes it for `GET /frame.png`, and the phone bind still draws it. `DisplayList` is a read path over the placement, and no paint behavior changed when it landed.
