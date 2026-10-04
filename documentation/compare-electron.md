# Compare with Electron

Recorded on 2026-10-03. The Electron column comes from the current Electron tutorial, the process model guide, and the API reference. The last column is what this repo does today. The v0.0.1 gap scan is [../plans/v0.0.1/compare.md](../plans/v0.0.1/compare.md); this file is the same gap list next to the framework it was written against.

## Hello world

| Electron | go-gpui |
|---|---|
| `package.json:`<br>`{ "main": "main.js", "scripts": { "start": "electron ." } }`<br>setup: `npm install electron --save-dev`<br><br>`main.js:`<br>`const { app, BrowserWindow } = require('electron/main')`<br><br>`const createWindow = () => {`<br>`  const win = new BrowserWindow({ width: 800, height: 600 })`<br>`  win.loadFile('index.html')`<br>`}`<br><br>`app.whenReady().then(() => {`<br>`  createWindow()`<br>`  app.on('activate', () => {`<br>`    if (BrowserWindow.getAllWindows().length === 0) createWindow()`<br>`  })`<br>`})`<br><br>`app.on('window-all-closed', () => {`<br>`  if (process.platform !== 'darwin') app.quit()`<br>`})`<br><br>`index.html:`<br>`<h1>Hello from Electron renderer!</h1>` | `page, err := gpui.New(gpui.Config{`<br>`  Title:  "Hello",`<br>`` HTML:   `<h1>{{.Title}}</h1>`, ``<br>`  Width:  480,`<br>`  Height: 640,`<br>`})`<br>`page.SetData(struct{ Title string }{"Hello"})`<br>`gpui.Run(context.Background(), page)`<br><br>`(one main.go, go run .)` |

The extra Electron lines buy a process model, a main process plus one renderer per window, and the window-lifecycle rules for macOS and Windows. This repo has none of that machinery, so none of it shows up in the program.

## What differs

| Aspect | Electron | go-gpui |
|---|---|---|
| What it is | Chromium plus a Node.js runtime with JS APIs | a Go library that lays out HTML and paints it in an Ebiten window |
| UI language | HTML, CSS, JavaScript | HTML, CSS, Go |
| App structure | main process plus one renderer process per window, IPC between them | one process, in-process callbacks |
| Data into the UI | DOM APIs, or IPC over a preload bridge | `page.SetData(any)` |
| Input | any DOM event with JS listeners | one `Handlers` struct: `Click`, `Submit`, `Type`, `Undo`, `Redo`, and the rest |
| Layout | Blink, the Chromium engine | gowkhtmltopdf: `html.Parse`, `css.Apply`, `layout.Lay` ([screen.md](screen.md)) |
| Painting | Chromium GPU compositor, the whole web platform | Ebiten canvas, display-list replay with a bitmap fallback ([screen.md](screen.md#replay)) |
| Targets | Windows, macOS, Linux | desktop, wasm in a browser, Android and iOS bind, plus a `-web` PNG page ([platforms.md](platforms.md), [web.md](web.md)) |
| Windows | multi-window, frameless, tray, menus, notifications | one window per page |
| Already in it | DevTools, IME, accessibility, video, canvas, WebGL, printing, drag and drop, auto-update, packaging | themes, IPC, HTML history, fetch, clipboard, crash reports, forms, a file open dialog ([features.md](features.md)) |
| Not in it | mobile targets, small downloads, a no-JS mode | Chromium fidelity, DevTools, IME, an accessibility tree, video and canvas, WebGL, multi-window, auto-update, an installer |
| Footprint | Chromium and Node bundled, installers run tens to over a hundred MB | one Go binary, a fraction of that |
| Maturity | 10+ years, runs VS Code, Slack, Discord, WhatsApp | v0.0.1, one main example |
| Building and shipping | npm, Forge or electron-builder, code signing, an update server | `go build`, then copy the binary |

## What Electron does better

- Fidelity. It is Chromium. Video, canvas, WebGL, IME, accessibility, printing, fonts, and every CSS feature work out of the box.
- Tooling. DevTools, hot reload, the npm ecosystem, electron-builder, auto-update, crash reporting, Playwright.
- Maturity. A decade of docs, a large community, and desktop apps used by millions. Most problems already have an answer on Stack Overflow.
- The native shell. Multiple windows, menus, tray icons, notifications, frameless and kiosk modes.
- The security model. The renderer sandbox and context isolation exist, once configured.
- Hiring. JavaScript developers are everywhere.

## What this repo does better

- Size and startup. No Chromium, no Node, no `node_modules`. One Go program, one binary, one process.
- One language. Go handlers, `go test`, `go build`. No JS, no preload script, no `contextBridge`, and no IPC serialization. A click handler is a direct function call.
- Reach. The same screen runs on desktop, as wasm, on Android, and on iOS. Electron is desktop only.
- No Chromium treadmill. Electron ships Chromium security patches on a schedule and callers have to keep up. This repo has a small attack surface and no browser to patch.
- A small API. `New`, `SetData`, `Handle`, `Run`, `Serve`, `BindMobile` ([window.md](window.md), [platforms.md](platforms.md)).
- A PNG endpoint for free. `-web` serves the same screen as a picture ([web.md](web.md)).

## What is still missing here

- No accessibility tree, no IME, and no spellcheck.
- No multi-window, no frameless mode, no tray, no native menus, no notifications.
- No video, canvas, or WebGL. Audio plays through Ebiten in the examples, not through an HTML element.
- No DevTools, no auto-update, no installer, no code signing.
- IPC is in-process only. There is no cross-process bridge and no JavaScript.
- Navigation is `Load`, `Back`, `Forward`, and `data-action` routes. There is no `loadURL`, no document `<a href>`, and no custom protocol.
- The display-list path needs a `gowkhtmltopdf` that exports `layout.DisplayList` and `Display.Boxes`, so `go.mod` carries a local `replace` for now.

## Sources

- <https://www.electronjs.org/docs/latest/tutorial/tutorial-first-app>
- <https://www.electronjs.org/docs/latest/tutorial/process-model>
