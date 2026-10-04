# Drag and drop

The window reads `ebiten.DroppedFiles()` once per frame, after the pointer
pass. An empty file system means no drop. A non-empty one becomes a slice of
`host.Drop` values, and the window calls `Drop` on a screen that implements
`host.Dropper`. A screen without the method ignores the files.

```go
type Drop struct {
	Name  string
	Size  int64
	IsDir bool
	Path  string
	Read  func() ([]byte, error)
}

type Dropper interface {
	Drop(ctx context.Context, files []Drop) error
}
```

`Path` is absolute on desktop and empty in a browser. `Read` returns the
bytes. A dropped directory arrives as one entry with `IsDir` true, and the
window never walks the tree. `Read` on a directory returns the file system's
error.

## The page handler

`page.Handlers.Drop` takes the same slice, and `gpui.Drop` is the root name
for the type. `Page.Drop` calls the handler and draws the page again. An
error returns before the redraw.

```go
page.Handle(gpui.Handlers{
	Drop: func(ctx context.Context, files []gpui.Drop) error {
		data, err := files[0].Read()
		if err != nil {
			return err
		}

		...
	},
})
```

## Frame scope

`ebiten.DroppedFiles` is scoped to the `Update` frame that reads it. The
window builds the values and calls `Drop` inside that frame, so a handler
has one frame to act:

- A handler that keeps `Path` can open the file later, on desktop.
- A handler that reads the bytes and stores them keeps them.
- A handler that returns without reading gets nothing, and a `Read` saved
  for later is not valid after the call.

The API cannot hide this. Read or copy what you need during the call.

## Web and wasm

`-web` runs the page through `internal/web`, not the Ebiten loop, so
`ebiten.DroppedFiles` is never read and no route carries a drop. The wasm
canvas build runs the same `Update` as the desktop window, so a browser drop
arrives there.

[examples/drop](../examples/drop) shows both previews: a dropped PNG or JPEG
through `SetImage`, and the first lines of a dropped `.txt`. Its test drives a
`testing/fstest.MapFS` through the handler.
