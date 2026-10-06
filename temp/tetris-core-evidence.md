# Tetris core evidence

Core agent record. Platform: Linux headless, go1.26.4 linux/amd64. No GUI
window was opened. Worktree `/tmp/opencode/gpui-wt-tetris`, branch
`feat/example-tetris`. Ruleset `tetris-classic-1` version 1, fixture version 1,
schema version 1, seed version 1. Revision at writing: `232b499`.

## Phase 1: model, fixtures, dummy data

- Command: `cd examples && go test -count=1 -p 2 ./tetris/game/...`. Result:
  ok, 52 tests.
- Fixtures: 6 named, `empty`, `near-top-out`, `clear-1` to `clear-4`. Each one
  is loaded and run in `fixture_test.go`.
- Dummy scores: 12 fixed records, IDs `dummy-0001` to `dummy-0012`. Reopen
  idempotence and user-edit survival are in `seed_test.go`.
- Fixed-seed sequences: `Sequence(seed, n)` is pinned in `sequence_test.go`;
  every bag is a permutation of the seven pieces.
- Versions: every score, snapshot, and replay row records ruleset and fixture
  version.

## Phase 2: core behavior

- The 52 game tests cover SRS rotation tables for T and I, wall and floor
  kicks, a blocked rotation in a one-wide well, every piece at every wall and
  the floor, blocked spawn and lock-above top-out, the 500 ms lock delay with
  the 15-reset cap, 1 to 4 simultaneous clears, score and level bounds, restart
  clearing every transient, gravity timing (1 s per row at level 1, 60 ms at
  level 20), the clock catch-up bound of 5 steps and the 500 ms stall drop,
  fixtures, snapshot round trip and rejection, and replay reproducibility.
- Command: `cd examples && go test -count=1 -p 2 ./tetris/input/...`. Result:
  ok, 14 tests. Coverage: DAS 170 ms, ARR 50 ms, soft-drop delay 40 ms,
  simultaneous directions with last-press priority, rapid taps, repeated
  rotation, one-shot actions, and ReleaseAll on focus loss.

## Phase 3: SQLite

- Command: `cd examples && go test -count=1 -p 2 ./tetris/store/...`. Result:
  ok, 43 tests.
- Covered: fresh migration, newer-schema rejection without modification, failed
  migration rollback, idempotent seeding, user edits surviving reopen, score
  insert idempotence by ID, tie-break by ID ascending, dummy and live ranking
  separation, batched pruning, settings round trip, snapshot round trip and
  corrupt rejection, busy retry within a deadline, a held lock giving up, a
  nested query timing out on the single connection, separate `:memory:`
  databases staying isolated, WAL activation on files, VACUUM INTO backup and
  restore, Close draining queued and in-flight writes, and canceled contexts.
- Store path: `os.UserConfigDir()/ownframe/tetris/tetris.db`, with `Open(dir)`
  as the explicit override and `OpenMemory()` for tests.
- Journal: WAL on files (asserted), `memory` for `OpenMemory`, rollback
  fallback when WAL fails. `Checkpoint` runs `PRAGMA wal_checkpoint(TRUNCATE)`
  and is a no-op off WAL.
- Durability: WAL plus `synchronous=FULL` on every connection. Scores,
  settings, snapshots, and replays are durable. There is no telemetry table.

## Phase 6, core half

- Command: `go test -race -p 2 ./tetris/game/... ./tetris/input/...
  ./tetris/store/...`. Result: ok, 3.3 s for store.
- Joint gate: `cd examples && go vet -p 2 ./tetris/...` is clean, and
  `go test -count=1 -p 2 ./tetris/...` is ok including `scene`.
- Repo gates from the worktree root: `make test TEST_P=4` exit 0 and
  `make build BUILD_P=4` exit 0. An earlier `make test` run failed in
  `tetris/scene` at `TestEndToEndWithMemoryStore` while the UI agent was
  editing `click.go`; the same test passes after that edit, and a copied
  instrumented build of the scene showed no store-side error. No
  `tetris-integration.md` entry was needed.
- Every Go file in game, input, and store is at most 2000 characters
  (`wc -m`); gofmt is clean.
- Pending, not run here: the native Windows checklist, a real window run
  through `ownframe.Run`, the 30-minute soak, and frame timings. This machine
  is headless, so no display path was exercised.
