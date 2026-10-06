# Tetris core API, the seam

Status: implemented and tested in `examples/tetris/game`, `examples/tetris/input`,
and `examples/tetris/store`. The call shapes below are final. Ruleset
`tetris-classic-1`, fixture version 1, schema version 1.

## Import paths

```go
"github.com/chinmay-sawant/ownframe/examples/tetris/game"
"github.com/chinmay-sawant/ownframe/examples/tetris/input"
"github.com/chinmay-sawant/ownframe/examples/tetris/store"
```

## The fixed-step tick

One window frame runs zero or more fixed simulation steps of
`game.FixedStep`, which is `time.Second / 60`. `game.Clock` does the
accumulator: it returns at most `game.MaxCatchUp` steps per call and
discards a gap at or over `game.StallLimit`, so a sleep or a stall pauses
play instead of fast-forwarding it.

```go
keys := input.NewTracker(input.DefaultKeymap())
clk := &game.Clock{}
g := game.New(seed)

page.SetTick(func(ctx context.Context) error {
	now := time.Now()
	if !ebiten.IsFocused() {
		keys.ReleaseAll() // no key-up arrives while unfocused
		clk.Reset(now)
	}

	var events []game.Event
	for i := 0; i < clk.Advance(now); i++ {
		for _, a := range keys.Step(game.FixedStep) {
			events = append(events, g.Apply(a)...)
		}
		events = append(events, g.Step(game.FixedStep)...)
	}

	// Read g for painting; handle events for status and storage.
	return nil
})

page.Handle(ownframe.Handlers{
	KeyDown: func(_ context.Context, key string) error { keys.KeyDown(key); return nil },
	KeyUp:   func(_ context.Context, key string) error { keys.KeyUp(key); return nil },
})
```

Run `Apply` before `Step` inside each iteration. `Apply` handles pause,
restart, and start immediately; the game ignores gameplay actions unless
the phase is `game.PhaseRunning`.

## Game state the scene reads

Exported fields on `game.Game`: `ID`, `Seed`, `Phase`, `Board`, `Piece`,
`Rot`, `X`, `Y`, `Next`, `Score`, `Lines`, `Level`, `Pieces`, `Elapsed`,
`Steps`.

- `Phase` is `PhaseReady`, `PhaseRunning`, `PhasePaused`, or `PhaseOver`.
- `Board` is `[20][10]Cell`, row 0 at the top. Zero is empty; other values
  are `Piece` numbers 1 to 7. Draw `ActiveCells()` and `GhostCells()` over it.
- `Next` has up to `game.Preview` pieces. `NextPiece()` returns the first.
- `New` starts ready on an empty board. `Start()` spawns and runs.
  `Restart()` clears every transient value and starts a fresh run.
  `TogglePause()` and `SetPaused(bool)` do what they say.

`Step` and `Apply` return `[]game.Event` with `Kind`, `Lines`, `Score`,
`Level`. Kinds: `EventSpawn`, `EventLock`, `EventClear`, `EventLevelUp`,
`EventHardDrop`, `EventTopOut`, `EventPause`, `EventRestart`.

`game.Result()` summarizes the run for storage. `game.Snapshot()` returns
the resumable state and `game.FromSnapshot(s)` rebuilds it or returns
`game.ErrInvalidSnapshot`.

## Input

`input.Tracker` maps lowercase ownframe key names to actions. A press
fires immediately; movement and rotation then repeat after `DAS` and every
`ARR`, and soft drop repeats after `SoftDropDelay`. Hard drop, pause,
restart, and start fire once. Holding two directions: the last pressed
wins, and the other resumes on release. `ReleaseAll()` is for focus loss.

`input.Keymap` is JSON-tagged and rides in `store.Settings`.

## Store

`store.Open("")` uses `store.DefaultDir()`, which is
`os.UserConfigDir()/ownframe/tetris`. `store.OpenMemory()` is for tests.
One worker goroutine owns one SQLite connection. Every method takes a
context and applies a deadline when the caller sets none.

Call store methods from a background goroutine, never from the tick. The
UI sends results back through its own queue and applies them in the tick.

```go
go func() {
	err := st.SaveGame(ctx, res, rep)
	// send err to the UI loop
}()
```

- `SaveGame(ctx, game.Result, *game.Replay)` is one transaction and is
  idempotent by `Result.ID`. A retry after a failure cannot duplicate the
  score. If it fails, play continues and the UI shows an unsaved badge and
  retries with the same result.
- `TopScores(ctx, limit)` lists live scores, highest first, ties broken by
  ID ascending. `DummyScores` lists seeded demo entries, which never mix
  into the live list.
- `SaveSettings` and `Settings` round-trip `store.Settings`. `SaveSnapshot`,
  `LoadSnapshot`, `ClearSnapshot` handle the single resume slot; invalid
  snapshots are rejected with `store.ErrInvalid`.
- `Checkpoint` and `Backup` run off the UI loop. `JournalMode` reports the
  active journal: `wal` on local files, `memory` for `OpenMemory`, and the
  rollback mode when WAL is unavailable.

Snapshot resume is enabled: one slot, written atomically with
`SaveSnapshot`, cleared by the UI on restart and game over.
