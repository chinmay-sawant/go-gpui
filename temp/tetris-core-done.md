# Tetris core done

Core agent handoff. Revision `232b499` plus the docs commit. All core packages
compile, are gofmt clean, and stay under the 2000-character file cap.

## What works

- `examples/tetris/game`: pure model. SRS rotations with wall and floor kicks,
  7-bag pieces on a splitmix64 generator with a saveable state, spawn at box
  row -1, lock delay 500 ms with at most 15 move resets, simultaneous 1 to 4
  line clears, score table 100/300/500/800 times level, level 1 plus lines/10
  capped at 20, soft drop 1 per cell, hard drop 2 per cell, pause, restart,
  top-out, next preview of 3, fixtures, snapshots, and replays.
- `examples/tetris/input`: held-key repeat in Go. DAS 170 ms, ARR 50 ms, soft
  drop 40 ms. Movement, rotation, and soft drop repeat. Hard drop, pause,
  restart, and start fire once. The last pressed direction wins, and the other
  resumes on release. `ReleaseAll` clears everything for focus loss.
- `examples/tetris/store`: one worker goroutine owns one SQLite connection.
  Migrations are versioned and transactional. Demo scores seed once behind a
  seed-version marker. Live and dummy rankings never mix. `SaveGame` is one
  transaction and idempotent by `Result.ID`, so a retry cannot duplicate a
  score. Equal scores order by ID ascending. Snapshots are one validated slot.
  WAL runs on files with `synchronous=FULL`; `Checkpoint` and `Backup` are for
  background calls.

## Test results

- `go test -count=1 -p 2 ./tetris/game/...` ok, 52 tests.
- `go test -count=1 -p 2 ./tetris/input/...` ok, 14 tests.
- `go test -count=1 -p 2 ./tetris/store/...` ok, 43 tests.
- `go test -race -p 2 ./tetris/game/... ./tetris/input/... ./tetris/store/...`
  ok.
- `go vet -p 2 ./tetris/...` clean; `go test -p 2 ./tetris/...` ok including
  `scene`.

## What the scene must know

- One tick iteration: `clock.Advance(now)` gives a step count, then for each
  step run `keys.Step(game.FixedStep)` and `g.Apply(action)`, then
  `g.Step(game.FixedStep)`. Apply before Step.
- On focus loss call `keys.ReleaseAll()` and `clock.Reset(now)`. A gap of
  500 ms or more is dropped by the clock, so play pauses instead of catching
  up. The scene already checks `ebiten.IsFocused()` in its tick.
- `Start` spawns on ready, resumes on paused, and restarts after over.
  `Restart` mints a new game ID. `FromSnapshot` returns a paused game; call
  `Start` to resume.
- Store methods take a context and must run off the tick. A failed
  `SaveGame` is retried with the same `game.Result`, including the same ID.
  A full queue returns `store.ErrBusy`; the scene treats it as unsaved.
- `Checkpoint` and `PruneDummy` are background jobs. Never call them from the
  tick.
- Replays record `InputEvent{Step, Action}` with the game's `Steps` counter,
  so `Replay.Play` reproduces the score for the same seed, ruleset, and
  fixture version.

## Decisions and bounds

- Spawn: box origin `(3, -1)`, O at `(4, -1)`, I cells land in row 0. Cells
  above the board are legal; a piece that locks entirely above the board tops
  out.
- Lock delay resets count toward the cap only while the piece stays grounded.
  A successful descent resets the counter.
- Score caps at 999,999,999, lines at 9999, level at 20, elapsed at 24 h.
- Store: `QueryTimeout` 2 s when the caller sets no deadline, queue 64,
  dummy IDs `dummy-NNNN`, WAL plus `synchronous=FULL`. There is no
  rollback-journal durability difference: the same synchronous setting
  applies.
- Resume is enabled: one snapshot slot, written atomically, cleared by the
  scene on restart and on game over.

## Still pending

- Native Windows desktop checklist items and any window run. This machine is
  headless.
- 30-minute soak and frame timings, which need the window path.
