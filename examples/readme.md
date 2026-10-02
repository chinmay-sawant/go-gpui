# Examples

Every example opens one HTML template in a window. `go run ./examples/<name>`
opens a desktop window. `go run ./examples/<name> -web` serves the same picture
in a browser on the printed address. Each example package also has a test that
runs headless.

The feature list these examples follow is
[../documentation/features-examples.md](../documentation/features-examples.md).

## Original examples

| Folder | Shows |
|---|---|
| [login](login) | Sign-in screen with real inputs, history, clipboard, undo. |
| [forms](forms) | Every form control: text, password, textarea, checkbox, radio, select, file. |
| [bind](bind) | `data-bind` two-way binding of controls to a struct. |

## Feature examples

| Folder | Feature | `-web` port |
|---|---|---|
| [layout](layout) | HTML -> CSS -> layout pipeline with hit-test boxes. | 8100 |
| [replay](replay) | Display-list replay and the bitmap fallback. | 8101 |
| [shapes](shapes) | Rounded and elliptical fills, masked strokes, letter-spaced text, caches. | 8102 |
| [png](png) | `Page.PNG` on-demand rasterization and caching. | 8103 |
| [controls](controls) | Click to focus, toggle, check, and cycle controls; typing edits; a file control opens the OS dialog. | 8104 |
| [bind-hooks](bind-hooks) | Two-way binding plus `BeforeEdit` veto and `Change` callback. | 8105 |
| [editing](editing) | Select all, undo, redo, and edits that survive a redraw. | 8106 |
| [states](states) | Host `:focus`, `:focus-visible`, `:hover`, `:active`, `:checked` styles. | 8107 |
| [scrolling](scrolling) | Oversized pages scroll on the wheel with scrollbar thumbs. | 8108 |
| [history](history) | `Load`, `Back`, `Forward`, `HTML`, and `data-action` routes. | 8109 |
| [web](web) | Web mode routes: `/`, `/frame.png`, `/click`, `/type`, `/backspace`. | 8110 |
| [ipc](ipc) | In-process `Send`, `Listen`, `Handle`, `Request`, with registered state and cancel. | 8111 |
| [fetch](fetch) | `Fetch` GET and `XHR` with method, headers, and body; a fetched PNG or JPEG becomes the page background. | 8112 |
| [clipboard](clipboard) | Copy, cut, and paste through the page clipboard API. | 8113 |
| [crash](crash) | `Report` files and panic recovery by `Run`. | 8114 |
| [platform](platform) | One page on desktop, WebAssembly, and mobile. | 8115 |

## Run

```sh
go run ./examples/layout          # desktop window
go run ./examples/layout -web     # browser on 127.0.0.1:8100
```

The platform example also builds for the browser canvas:

```sh
GOOS=js GOARCH=wasm go build ./examples/platform
```

Cursor keys, `Ctrl+C`/`Ctrl+X`/`Ctrl+V`, `Ctrl+Z`, and `Ctrl+Y` are handled by
the window, not by the examples.

## Tests

```sh
make test
```

`make test` runs `go test -p 1 ./...` by default so the packages do not all
build and run at once. Set `TEST_P` for more parallelism: `make test TEST_P=4`.
