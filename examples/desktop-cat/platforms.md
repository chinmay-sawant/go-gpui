# Transparency across platforms

`gpui.RunWithOptions` is available to every desktop caller. Transparency
requires both `WindowOptions.Transparent` and a page whose background and
image assets preserve alpha. `MousePassthrough` passes all pointer events,
including events over painted pixels, to other windows. `Interactive` instead reserves opaque cat pixels and bubble regions for input and updates
mouse passthrough from the current cursor position each frame.

| Platform | Transparent background | Click-through | Evidence and limits |
|---|---|---|---|
| Windows | Native window alpha | Native window mouse passthrough | Live native capture and input test: empty margins target the underlying app; the bubble accepts a click to dismiss, and the cat reopens it. |
| macOS | Native window alpha | Cocoa ignores mouse events | Toolkit source supports both. Cross-compilation checks cover the example; no Mac runtime was available. |
| Linux | X11/XWayland alpha, with a compositor | X11 input shape | Toolkit source supports both. This WSLg session opened a 24-bit framebuffer and its outer window intercepted input. Regular Linux has not been verified live here. |
| WSL | Native Windows window via the example launcher | Native Windows mouse passthrough | The workaround is in this example, rather than every caller of the library. Windows interop and a source checkout are required. |
| WASM | Transparent browser canvas | Canvas input | The final Go example uses the generic WASM loader. The custom JavaScript wrapper was removed; desktop click-through and dragging do not apply. |
| `Serve` / `-web` | Picture preview | Browser image/page events | No native overlay and no frame ticks. Rasterized preview backgrounds can differ from native alpha rendering. |

The toolkit's `ScreenTransparent` option applies to desktops and browser
canvases. Its mouse passthrough API applies only to desktop windows. macOS
implements it with Cocoa's `setIgnoresMouseEvents`. The Go WASM example uses the generic browser loader.

Sources: [Ebitengine window API](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#SetWindowMousePassthrough)
and [RunGameOptions](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#RunGameOptions).

The desktop-cat example now supports `Draggable` for user-initiated window
movement. State tests verify its press, movement threshold, deltas, and
release behavior. The new native drag gesture and pixel mask were not tested
with desktop mouse automation after the user asked us to leave their cursor
alone. The earlier headless wrapper was subsequently removed. The
Windows evidence in the table predates those additions.

Chrome media detection is native Go code using Windows WinRT. It is available
on Windows and through the WSL launcher. The monitor reads metadata and
playback status; it never sends playback control commands or moves the mouse.
Regular Linux, macOS, and WASM have no native media monitor in this example.
