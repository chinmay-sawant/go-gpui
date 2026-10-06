# Tetris checklist

Recorded 2026-10-07. Status: implemented and verified headless; native
Windows, soak, and profiling checks pending. Every checked item has an
evidence note under its phase and an entry in the Evidence log.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A playable Go-native falling-block game with rotation, line clearing, increasing difficulty, next-piece preview, score, pause/restart, and durable scores. The goal is correct play and reliable input.

## Phase 1: Dummy data and runnable foundation

- [x] Create `examples/tetris/` with a pure game model, input adapter, storage, and ownframe scene. Define board coordinates, piece rotations, collision rules, scoring, and game states.
- [x] Start with fixed-seed dummy piece sequences, seeded score records, and selectable board fixtures for an empty board, near-top-out, and one-to-four line clears. Default play begins from a valid empty board.
- [x] Separate demo score records from player scores. Record the ruleset and fixture versions so tests and replay inputs remain reproducible.

Evidence, 2026-10-07: `cd examples && go test -p 2 -count=1 ./tetris/...` ok (game 52 tests, input 14, store 43, scene suite green). Six named fixtures (`empty`, `near-top-out`, `clear-1` to `clear-4`), 12 demo records `dummy-0001` to `dummy-0012`, and pinned `Sequence(seed, n)` output are covered in the core tests; the scene end-to-end test loads the demo page and tags every row DEMO. Paths: `temp/tetris-core-evidence.md`, `temp/tetris-ui-evidence.md`.

## Phase 2: Core behavior and edge cases

- [x] Implement spawn, movement, rotation, soft/hard drop, lock timing, simultaneous line removal, scoring, level progression, next preview, pause, restart, and game over.
- [x] Use an explicit fixed simulation step through the ownframe tick. Bound catch-up work after stalls; pause rather than fast-forwarding gameplay after focus loss or system sleep.
- [x] Implement held-key repeat in Go because ownframe key callbacks do not forward OS auto-repeat pulses. Clear held inputs after focus loss and test simultaneous directions, rapid taps, and repeated rotation.
- [x] Specify wall/floor rotation behavior and test every piece near board edges, blocked spawn, lock-delay resets, top-out, multiple clears, empty ghost/drop paths, and high-speed gravity.
- [x] Update retained scene operations for routine frames; rebuild only when needed and rebind after resize. Ensure displayed score and board match the model after both replay and bitmap rendering.
- [x] Bound scores/timing values and ensure restart clears every transient game and input state.

Evidence, 2026-10-07: the game tests cover SRS tables, wall and floor kicks, every piece at every wall and the floor, blocked spawn, lock delay with the reset cap, 1 to 4 clears, score/level bounds, gravity timing, and restart clearing every transient. The scene tests cover fixed steps through `game.Clock`, the 500 ms stall pause, focus-loss clear plus pause, refocus resume, retained-operation repaint, the bitmap fallback path, and rebind on a generation change. `BenchmarkTickPaint` is 2579 ns/op on the i7-13700HX. Native resize itself is pending under Phase 5.

## Phase 3: SQLite storage and failure handling

- [x] Use the examples module's existing `modernc.org/sqlite` dependency through `database/sql`; keep storage code out of the root public library.
- [x] Give this example a separate database under `os.UserConfigDir()/ownframe/<example>/`. Support an explicit data-directory override. Use `filepath` and correctly encoded driver paths for Windows drive letters, spaces, Unicode, and URI punctuation.
- [x] Implement versioned, transactional migrations and idempotent dummy seeding with a seed-version marker. Reopening must not duplicate fixtures or overwrite user changes. Reject a newer unsupported schema without modifying it.
- [x] Start with one serialized database worker and one connection; apply foreign keys and a finite busy timeout to every physical connection if the pool grows. Keep transactions short and close rows before issuing more queries. Test single-connection self-deadlocks and separate `:memory:` databases across connections.
- [x] Use parameterized statements, explicit constraints, query deadlines, and bounded operation queues. Retry only appropriate lock failures within a deadline; never blindly replay non-idempotent operations.
- [x] For local writable storage, verify WAL activation and manage checkpoint work off the UI loop. Support a documented rollback-journal policy for unsupported storage. WAL requires local host access and still allows only one writer. See [SQLite WAL](https://sqlite.org/wal.html).
- [x] Choose and document durability per table: durable user edits and job transitions, with explicitly disposable telemetry. Never silently claim a failed write was saved.
- [ ] Exercise locked database, second application instance, read-only directory, disk full, failed commit, interrupted migration, corruption, deleted storage, and permission changes. Preserve existing files; offer retry, explicit recovery, or clearly labeled temporary mode without silently resetting data. (pending: locked database, failed migration rollback, failed commit, and corrupt snapshots are covered; second instance, read-only directory, disk full, deleted storage, and permission changes are not exercised)
- [ ] Back up with SQLite-aware operations rather than copying an open database alone. Test restore compatibility. Close database, rows, and file handles before Windows rename/delete; never delete active WAL or SHM files. (pending: VACUUM INTO backup and restore pass; Windows rename/delete behavior needs a native host)
- [x] Define retention and cleanup in bounded batches; do not run large vacuum or maintenance jobs during an interaction. Test cancellation and shutdown with writes in flight.
- [x] Store scores, ruleset versions, control settings, and optional replay seeds/input events. Write one completed-game transaction off the tick; never save the whole board every frame.
- [x] Assign each completed game a stable unique ID to prevent duplicate score insertion after retries or restart. Tie-break equal scores by stable ID and record replay/ruleset compatibility.
- [x] Allow play to continue after score-storage failure with a visible unsaved status and retry. Define whether an active game is resumable; if enabled, persist all required state atomically and reject invalid snapshots.

Reference: [modernc SQLite driver](https://pkg.go.dev/modernc.org/sqlite) and [SQLite PRAGMAs](https://sqlite.org/pragma.html).

Evidence, 2026-10-07: store suite ok (43 tests) and `go test -race -p 2 ./tetris/game/... ./tetris/input/... ./tetris/store/...` ok. WAL activation is asserted on files, seeding is idempotent behind a version marker, `SaveGame` is idempotent by `Result.ID`, and backups use `VACUUM INTO`. The scene worker runs every call off the tick, drains at most four results per frame, shows `SCORE NOT SAVED - RETRYING`, and retries every two seconds; the scene tests cover a refused queue and a failed result.

## Phase 4: Pagination and bounded rendering

- [x] Keep the fixed-size board fully visible; it does not need pagination. Scale or enforce a usable minimum window size without changing logical coordinates.
- [x] Page score history with indexed ordering by score and stable ID, initially 20 entries. Keep dummy entries identifiable and separate from live rankings.
- [x] Handle tied scores, score insertion between pages, no scores, expired replay files, and first/last boundaries. Preserve board state while opening and closing history.

Evidence, 2026-10-07: the window minimum is 680x740 CSS pixels and the board keeps its 10x20 logical coordinates. The scene history tests cover 20-row pages, padding, first/last boundaries, `PAGE n+` when more may exist, the no-scores message, live/demo separation, a refetch after a new top score lands between pages, and board preservation across open/close. The core tests cover tie-break by ID ascending and replay validation (ruleset and fixture compatibility); the store lists live and dummy rankings separately.

## Phase 5: Windows and Ebiten integration

- [x] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling.
- [x] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [x] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback.
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input. (pending: native Windows host; the minimum size and focus-loss pause are implemented and headless-tested)
- [ ] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database. (pending: close during in-flight writes is covered, close during migration is not exercised)
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing. (pending: this machine is headless Linux)
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges. (pending: a held lock is covered; access-denied and read-only paths are not exercised)
- [x] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview.

Evidence, 2026-10-07: `main.go` calls `ownframe.Run` and opens the store worker; the worker has an 8-slot request queue and a 32-slot result queue, and the tick drains four results per frame with per-kind request generations. Retained operation pointers rebind on every generation change, the nil display list path rebuilds with a 50 ms throttle and has tests, and a phase change is the only rebuild a running game triggers. Headless costs: `BenchmarkTickPaint` 2579 ns/op, `BenchmarkRedraw` 6102239 ns/op, `BenchmarkTickBitmap` 1373298 ns/op. `-web` was not added; the example is desktop only.

## Phase 6: Verification and completion

- [x] Write focused tests for domain behavior, migration/restart recovery, persistence failures, stale asynchronous results, and shutdown. Run concurrency tests with fake collectors, files, or HTTP services; keep clipboard tests on `clipboard.UseMemory`.
- [ ] Run the finished example from a fresh data directory using dummy data, then with its real source. Exercise the full user workflow, restart, and verify saved state in the actual database. (pending: a headless file-store end-to-end test drives a fresh directory, plays, saves, reopens, and reads the score back; the native window run and process restart are pending)
- [ ] Profile representative and stress workloads. Record CPU, heap, RSS, allocations, goroutines, queue depth, dropped/coalesced updates, and frame/input timings. After warmup, a 30-minute mixed-workload soak must show bounded history and no sustained memory or goroutine growth. (pending: 30-minute soak and real profiling need a window session)
- [ ] Target p99 UI frame work below 16.67 ms at 60 Hz and input feedback within 100 ms on a named reference machine. These are acceptance targets, not measured claims. Report failures and renderer limitations explicitly. (pending: native renderer frame timings; headless scene numbers are recorded in the README)
- [x] Run `gofmt` on changed Go files, keep each Go file at most 2000 characters, run `make test` and `make build` from the repository root, and run targeted race tests where supported. Store generated evidence under ignored `temp/`.
- [ ] Add usage, dummy/real mode, storage location, recovery, platform limits, and measured performance to the example README and examples index. (pending: `examples/tetris/README.md` covers all of it; the `examples/readme.md` index entry is owned by the orchestrator)
- [ ] Mark complete only when the example works end to end, persistence survives restart, required Windows scenarios pass, and all earlier phase evidence is recorded. Record unavailable checks as pending. (pending: native Windows scenarios, soak, and profiling)

Evidence, 2026-10-07: `gofmt -l` clean, every tetris Go file at most 2000 characters (`wc -m`), `go vet -p 2 ./tetris/...` clean, `go test -p 2 -count=1 ./tetris/...` ok, `go test -race -p 2 ./tetris/scene/` ok, and from the worktree root `make build BUILD_P=4` exit 0 and `make test TEST_P=4` exit 0. No window was opened on this headless machine.

## Evidence log

- 2026-10-07, branch `feat/example-tetris`, UI revision `adae4d1`, core revision `232b499`, Linux headless amd64, Go 1.26.4, 13th Gen Intel Core i7-13700HX. Commands: `go vet -p 2 ./tetris/...` clean; `go test -p 2 -count=1 ./tetris/...` ok (game 52, input 14, store 43, scene suite); `go test -race -p 2 ./tetris/scene/` ok; `make build BUILD_P=4` exit 0; `make test TEST_P=4` exit 0; benchmarks `BenchmarkTickPaint` 2579 ns/op, `BenchmarkRedraw` 6102239 ns/op, `BenchmarkTickBitmap` 1373298 ns/op. No window was opened. Evidence files: `temp/tetris-core-evidence.md`, `temp/tetris-ui-evidence.md`, `temp/tetris-integration.md`. Pending: native Windows session, 30-minute soak, real profiling, and the examples index entry.

For each phase, add date, revision, command or manual workflow, dataset size, platform/renderer, result, and log or screenshot path here. Leave unsupported or unrun scenarios unchecked.
