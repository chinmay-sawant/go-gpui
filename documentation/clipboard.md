# Clipboard

`clipboard.Write` stores the text in memory, then tries the desktop clipboard. `clipboard.Read` returns the desktop clipboard when that read works, and otherwise returns the last `Write`.

On Linux, `Write` and `Read` look for a Wayland or X11 helper:

- `WAYLAND_DISPLAY` set and `wl-copy` or `wl-paste` on `PATH`.
- Otherwise `DISPLAY` set and `xclip` on `PATH`.

Each helper is given 200 milliseconds. Windows, macOS, Android, iOS, and wasm have no OS clipboard path in this build. They keep the in-memory copy only.

The window calls `Write` and `Read` for the text chords. Ctrl or Cmd with C, X, and V, plus Shift-Insert, Ctrl-Insert, and Shift-Delete. The login example copies the real password, not the mask.

Tests call the in-memory copy. They do not write the desktop clipboard. `~/.Xauthority` is not created or changed by this package.

A native clipboard, without `wl-copy` or `xclip`, is not on `feature/v0.0.1` yet.
