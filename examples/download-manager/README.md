# Download manager

A local download queue in an ownframe window: a queue table with compact
progress bars, job details, paged history, and a light or dark theme. Status
text carries a symbol as well as a color, so it reads without color.

## Run

```sh
cd examples
go run ./download-manager                  # dummy data, default storage
go run ./download-manager -real            # real HTTP transfers
go run ./download-manager -data /tmp/dm    # explicit storage directory
go run ./download-manager -fixture         # start the local fixture server
```

Dummy mode is the default. It seeds 100 deterministic jobs once, serves them
through an in-memory fake transport, and opens no socket. Dummy files land in
`os.TempDir()/ownframe-download-manager-dummy`.

Real mode uses the HTTP transport. With `-fixture` the program starts a local
HTTP service and logs its URLs (`/ok`, `/slow`, `/chunked`, `/range`,
`/changing`, `/redirect`, `/interrupt`). Paste one into the URL field to
exercise a case; any http or https URL works.

## Storage

The database is SQLite through `modernc.org/sqlite`, by default under
`os.UserConfigDir()/ownframe/download-manager/`. `-data DIR` overrides it, and
the theme toggle is saved in `ui.json` beside the database.

A second instance prints `another instance is using ...` and exits. When the
directory will not open, the example falls back to an in-memory database and
labels the header `in-memory (temporary)`; nothing persists in that mode.

## Recovery

The store reconciles rows against the files on disk at startup. A running job
becomes paused and stays resumable, a complete file is adopted as completed,
and a missing file marks the job failed. Partials use the `.ownframe-part`
suffix and are finalized only after validation.

## Shutdown

Closing the window stops the tick, cancels the worker, and waits up to one
second for it, then up to six seconds for the engine, whose own transfer
budget is five seconds. The total shutdown budget is seven seconds.

## Platform limits

Desktop only, through `ownframe.Run`. There is no web or mobile mode, so
`Serve`, `BindMobile`, and the wasm build do not apply. The minimum window is
780x520 CSS pixels; a larger page scrolls. Windows display scaling, monitor
moves, and minimize/restore are not verified on this machine (Linux, headless
tests only) and stay pending.

## Performance notes

Measured 2026-10-07 on an i7-13700HX, Linux, with
`go test -bench . ./download-manager/ui/`, 20 queue rows and 50 history rows:

| Work | Cost |
|------|------|
| Tick that only repaints retained operations, one update | 4.0 µs |
| Tick that relayouts first | 15.8 ms |
| Full `Redraw` | 21.9 ms |

A registered tick makes the window replay the full display list every frame;
these numbers cover the callback and the relayout, not the Ebiten draw. A
progress update stays on the retained operation path and does not relayout,
so a running queue costs the 4 µs tick. A state change, a filter change, or a
resize pays the 16 ms to 22 ms relayout, which is above one 60 Hz frame.
Profiling and a 30-minute soak are pending.

## Layout

- `main.go` parses the flags and opens the window.
- `ui/` owns the page, the keyset pager, the retained operations, and the theme.
- `ui/wire/` adapts `store`, `scheduler`, and `transfer`: one worker runs commands off the UI loop, and the tick polls engine events within a bounded budget.
- `domain/`, `transfer/`, `scheduler/`, `store/`, and `fixture/` hold the core.
