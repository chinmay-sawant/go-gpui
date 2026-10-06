# Tetris checklist

Recorded 2026-10-07. Status: planned. All implementation items remain unchecked.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A playable Go-native falling-block game with rotation, line clearing, increasing difficulty, next-piece preview, score, pause/restart, and durable scores. The goal is correct play and reliable input.

## Phase 1: Dummy data and runnable foundation

- [ ] Create `examples/tetris/` with a pure game model, input adapter, storage, and ownframe scene. Define board coordinates, piece rotations, collision rules, scoring, and game states.
- [ ] Start with fixed-seed dummy piece sequences, seeded score records, and selectable board fixtures for an empty board, near-top-out, and one-to-four line clears. Default play begins from a valid empty board.
- [ ] Separate demo score records from player scores. Record the ruleset and fixture versions so tests and replay inputs remain reproducible.

## Phase 2: Core behavior and edge cases

- [ ] Implement spawn, movement, rotation, soft/hard drop, lock timing, simultaneous line removal, scoring, level progression, next preview, pause, restart, and game over.
- [ ] Use an explicit fixed simulation step through the ownframe tick. Bound catch-up work after stalls; pause rather than fast-forwarding gameplay after focus loss or system sleep.
- [ ] Implement held-key repeat in Go because ownframe key callbacks do not forward OS auto-repeat pulses. Clear held inputs after focus loss and test simultaneous directions, rapid taps, and repeated rotation.
- [ ] Specify wall/floor rotation behavior and test every piece near board edges, blocked spawn, lock-delay resets, top-out, multiple clears, empty ghost/drop paths, and high-speed gravity.
- [ ] Update retained scene operations for routine frames; rebuild only when needed and rebind after resize. Ensure displayed score and board match the model after both replay and bitmap rendering.
- [ ] Bound scores/timing values and ensure restart clears every transient game and input state.

## Phase 3: SQLite storage and failure handling

- [ ] Use the examples module's existing `modernc.org/sqlite` dependency through `database/sql`; keep storage code out of the root public library.
- [ ] Give this example a separate database under `os.UserConfigDir()/ownframe/<example>/`. Support an explicit data-directory override. Use `filepath` and correctly encoded driver paths for Windows drive letters, spaces, Unicode, and URI punctuation.
- [ ] Implement versioned, transactional migrations and idempotent dummy seeding with a seed-version marker. Reopening must not duplicate fixtures or overwrite user changes. Reject a newer unsupported schema without modifying it.
- [ ] Start with one serialized database worker and one connection; apply foreign keys and a finite busy timeout to every physical connection if the pool grows. Keep transactions short and close rows before issuing more queries. Test single-connection self-deadlocks and separate `:memory:` databases across connections.
- [ ] Use parameterized statements, explicit constraints, query deadlines, and bounded operation queues. Retry only appropriate lock failures within a deadline; never blindly replay non-idempotent operations.
- [ ] For local writable storage, verify WAL activation and manage checkpoint work off the UI loop. Support a documented rollback-journal policy for unsupported storage. WAL requires local host access and still allows only one writer. See [SQLite WAL](https://sqlite.org/wal.html).
- [ ] Choose and document durability per table: durable user edits and job transitions, with explicitly disposable telemetry. Never silently claim a failed write was saved.
- [ ] Exercise locked database, second application instance, read-only directory, disk full, failed commit, interrupted migration, corruption, deleted storage, and permission changes. Preserve existing files; offer retry, explicit recovery, or clearly labeled temporary mode without silently resetting data.
- [ ] Back up with SQLite-aware operations rather than copying an open database alone. Test restore compatibility. Close database, rows, and file handles before Windows rename/delete; never delete active WAL or SHM files.
- [ ] Define retention and cleanup in bounded batches; do not run large vacuum or maintenance jobs during an interaction. Test cancellation and shutdown with writes in flight.

Reference: [modernc SQLite driver](https://pkg.go.dev/modernc.org/sqlite) and [SQLite PRAGMAs](https://sqlite.org/pragma.html).

- [ ] Store scores, ruleset versions, control settings, and optional replay seeds/input events. Write one completed-game transaction off the tick; never save the whole board every frame.
- [ ] Assign each completed game a stable unique ID to prevent duplicate score insertion after retries or restart. Tie-break equal scores by stable ID and record replay/ruleset compatibility.
- [ ] Allow play to continue after score-storage failure with a visible unsaved status and retry. Define whether an active game is resumable; if enabled, persist all required state atomically and reject invalid snapshots.

## Phase 4: Pagination and bounded rendering

- [ ] Keep the fixed-size board fully visible; it does not need pagination. Scale or enforce a usable minimum window size without changing logical coordinates.
- [ ] Page score history with indexed ordering by score and stable ID, initially 20 entries. Keep dummy entries identifiable and separate from live rankings.
- [ ] Handle tied scores, score insertion between pages, no scores, expired replay files, and first/last boundaries. Preserve board state while opening and closing history.

## Phase 5: Windows and Ebiten integration

- [ ] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling.
- [ ] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [ ] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback.
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input.
- [ ] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database.
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing.
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges.
- [ ] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview.

References: [ownframe frames](../../../documentation/frames.md), [features](../../../documentation/features.md), and [Ebitengine lifecycle](https://ebitengine.org/en/documents/cheatsheet.html).

## Phase 6: Verification and completion

- [ ] Write focused tests for domain behavior, migration/restart recovery, persistence failures, stale asynchronous results, and shutdown. Run concurrency tests with fake collectors, files, or HTTP services; keep clipboard tests on `clipboard.UseMemory`.
- [ ] Run the finished example from a fresh data directory using dummy data, then with its real source. Exercise the full user workflow, restart, and verify saved state in the actual database.
- [ ] Profile representative and stress workloads. Record CPU, heap, RSS, allocations, goroutines, queue depth, dropped/coalesced updates, and frame/input timings. After warmup, a 30-minute mixed-workload soak must show bounded history and no sustained memory or goroutine growth.
- [ ] Target p99 UI frame work below 16.67 ms at 60 Hz and input feedback within 100 ms on a named reference machine. These are acceptance targets, not measured claims. Report failures and renderer limitations explicitly.
- [ ] Run `gofmt` on changed Go files, keep each Go file at most 2000 characters, run `make test` and `make build` from the repository root, and run targeted race tests where supported. Store generated evidence under ignored `temp/`.
- [ ] Add usage, dummy/real mode, storage location, recovery, platform limits, and measured performance to the example README and examples index.
- [ ] Mark complete only when the example works end to end, persistence survives restart, required Windows scenarios pass, and all earlier phase evidence is recorded. Record unavailable checks as pending.

## Evidence log

For each phase, add date, revision, command or manual workflow, dataset size, platform/renderer, result, and log or screenshot path here. Leave unsupported or unrun scenarios unchecked.
