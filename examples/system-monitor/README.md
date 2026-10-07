# System monitor

A desktop ownframe example that shows CPU, memory, disk, and network
activity, a searchable process table, and a process detail view. It is
read-only: nothing in it terminates a process.

## Run

From the `examples` directory:

    go run ./system-monitor

Dummy mode is the default. It reads no host state, needs no permissions, and
starts from a fixed seed with 250 processes, 120 historical samples, CPU
spikes, unavailable sensors, and process exits, so the first screen is useful
on any machine.

    go run ./system-monitor -mode live     # read the host instead
    go run ./system-monitor -stress        # 10,000-process dummy fixture
    go run ./system-monitor -data /tmp/sm  # data directory override
    go run ./system-monitor -seed 42       # dummy fixture seed

| Flag | Default | Meaning |
| --- | --- | --- |
| `-mode` | `dummy` | `dummy` or `live` |
| `-data` | user config dir | SQLite directory override |
| `-stress` | `false` | build the 10,000-process fixture |
| `-seed` | `1` | dummy fixture seed |

Live mode reads `/proc` on Linux and kernel32 on Windows. Elsewhere every
value arrives invalid and the cards read `n/a`.

## Screens

The sidebar switches three screens.

Overview draws one card per metric: CPU, memory, disk, and network. Each card
has a numeric value, a secondary line, a peak label, and a line graph drawn
from a fixed-size buffer. The graph uses one accent color, the value is
always printed as text, and an unavailable sensor reads `n/a` rather than
`0`.

Processes shows a frozen snapshot, 50 rows per page. The search box filters
by name or PID and resets to page one. The column headers sort by PID, name,
CPU, or memory; equal keys tie-break on the process identity, so equal rows
never swap between snapshots. Paging and sorting never shift rows under the
cursor. A newer snapshot waits behind the Refresh button and the "new sample
ready" label, and refresh keeps the selected process by identity even when
it arrived or exited in the meantime. An exited process keeps its
last-known row and the detail card says "no longer running". Previous and
Next only appear when a page exists in that direction.

Details follows the selected process. The card refreshes its values in place
from the collector, which reads full details less often than the summary.

The theme button flips light and dark. The window minimum is 900x600, and a
page taller than the window scrolls on the wheel.

## Storage

The default data directory is `os.UserConfigDir()/ownframe/system-monitor`,
and `-data` replaces it. The store is SQLite through
`modernc.org/sqlite`; schema migrations run at open. The UI saves the theme
toggle under the `ui.theme` setting and nothing else. On a fresh dummy run
the example seeds the labeled 120-sample fixture once; reopening does not
duplicate it.

## Recovery

- A database written by a newer schema makes `storage.Open` return
  `storage.ErrSchemaNewer`, and the program exits without modifying the
  file.
- A seeding failure is logged and the window still opens.
- A failed theme read falls back to light; a failed theme write shows a
  notice in the window instead of claiming the setting was saved.
- `ownframe.Run` writes a crash report and returns an error that names the
  report file.

## Platform limits

- Native Windows behavior is not verified here. The 100/150/200 percent
  scaling, monitor move, minimize/restore, and focus-loss checks stay
  pending until someone runs them on a Windows desktop.
- There is no `-web` mode. `Serve` does not tick, so retained-operation
  edits would not appear in a still preview.
- The measured numbers below come from a headless Linux machine and cover
  the page pipeline only. Frame and input timing needs a real window session.

## Measured performance

Headless on a 13th Gen Intel Core i7-13700HX, Linux, no window renderer:

    cd examples
    go test -run '^$' -bench . -benchtime 200x -benchmem ./system-monitor/ui

| Work | Cost |
| --- | --- |
| Active tick, four graphs, retained operations | about 2 to 4 us, 8 allocs |
| Full repaint, overview, 384 graph bars | about 14 to 30 ms |
| Full repaint, process screen, 50 of 10,000 rows | about 7 to 9 ms |
| Refresh plus page slice, 10,000 processes | about 0.4 to 0.8 ms |
| Overview display list at 1024x700 | 418 operations, 432 boxes |

The ranges cover repeated runs on the same idle machine. The tick stays far
below the 16.67 ms frame budget at 60 Hz. A full repaint happens on a click,
a resize, or a refresh, not per frame, and it stays inside the 100 ms input
feedback target. The process screen renders 50 rows whatever the snapshot
size. The graphs downsample the 120-sample ring into at most 96 columns
chosen from the visible width, so paint work never follows the history size.

## Tests

    cd examples
    go test ./system-monitor/...

The UI tests cover paging, filter reset, the frozen snapshot, stale-result
discard, formatting, graph downsampling, and the tick paint path against a
real page, all headless. The collector tests live under `collector/`.
