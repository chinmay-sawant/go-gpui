# Live log viewer checklist

Recorded 2026-10-07. Status: implemented and verified headless on Linux.
Native Windows checks, a 30-minute soak, and real profiling stay unchecked
with pending notes.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A working desktop viewer that follows multiple local log files, searches and filters entries, pauses display, and browses stored history with bounded resource use.

## Theme

Minimal developer-tool theme: a compact source sidebar, a readable monospaced log area, and a small filter bar. Provide light and dark modes with a persistent toggle. Severity has text labels as well as restrained color.

Done: `ui/ui.html` holds the light theme, `ui/theme_dark.css` the dark
overrides, and the toggle persists through the adapter settings file.

## Phase 1: Dummy data and runnable foundation

- [x] Create `examples/live-log-viewer/` with reader, parser, store, and UI packages. Define source/session IDs, monotonically ordered entry IDs, offsets, and ingestion policy.
- [x] Default to a seeded dummy generator producing 10,000 initial entries and a controlled stream. Include mixed severity, Unicode, malformed timestamps, long entries, repeated messages, and multiline records.
- [x] Provide a reproducible burst fixture and explicit rate/size limits. Dummy mode requires no user files and has its own sessions.

Evidence: `go test -p 2 ./live-log-viewer/...` passes for entry, parser,
reader, store, and ui. The 10,000-entry fixture shapes, deterministic seed,
and burst caps are in `temp/live-log-viewer-core-evidence.md`. The UI runs
against the same store through `ui/feed_store.go`; its dummy-flow test uses
a 500-entry temporary database.

## Phase 2: Core behavior and edge cases

- [x] Implement multiple sources, follow/pause, source/severity/text filters, entry details, and export of a bounded selection.
- [x] Use cancellable readers and bounded ingestion batches. Define exactly what happens under overload: pause file ingestion when possible, or report counted loss for a non-replayable source. Never silently discard log entries.
- [x] Separate paused display from ingestion: pausing the view may continue bounded recording, with an unread count. Resume to the chosen position instead of unexpectedly moving the user.
- [x] Handle append, truncate, rotation, rename/delete/recreate, partial UTF-8, CRLF, missing newline, malformed encoding, and multiline boundaries. Limit maximum record bytes and show truncation explicitly.
- [ ] For Windows rotation, test file sharing and open-handle behavior; use a platform reader that permits expected rename/delete operations and can reopen rotated files. (pending: native Windows host; the sharing adapter compiles under `GOOS=windows go vet`, but no native session ran)
- [x] Compile filters off the UI loop and bound expensive searches. Support plain-text search first; any later regex engine must have a documented resource bound.

Evidence: `TestAppFollowsTail` proves paused display keeps counting and
resume does not jump; reader and parser edge tests and the overload tests
(`TestStreamLoss`, `TestIngestOverloadLoss`, `TestFileLagNotLoss`) are in the
core evidence file. Search debounces 250 ms, runs on the worker, and is
bounded by the store page cap and a 5 s query deadline. Rows show `trunc`
and `part` tags; the detail pane labels both.

## Phase 3: SQLite storage and failure handling

- [x] Use the examples module's existing `modernc.org/sqlite` dependency through `database/sql`; keep storage code out of the root public library.
- [x] Give this example a separate database under `os.UserConfigDir()/ownframe/<example>/`. Support an explicit data-directory override. Use `filepath` and correctly encoded driver paths for Windows drive letters, spaces, Unicode, and URI punctuation.
- [x] Implement versioned, transactional migrations and idempotent dummy seeding with a seed-version marker. Reopening must not duplicate fixtures or overwrite user changes. Reject a newer unsupported schema without modifying it.
- [ ] Start with one serialized database worker and one connection; apply foreign keys and a finite busy timeout to every physical connection if the pool grows. Keep transactions short and close rows before issuing more queries. Test single-connection self-deadlocks and separate `:memory:` databases across connections. (pending: one worker, one connection, `busy_timeout(5000)`, foreign keys, and `TestMemoryStoresAreIndependent` are verified; no self-deadlock test was found)
- [x] Use parameterized statements, explicit constraints, query deadlines, and bounded operation queues. Retry only appropriate lock failures within a deadline; never blindly replay non-idempotent operations.
- [x] For local writable storage, verify WAL activation and manage checkpoint work off the UI loop. Support a documented rollback-journal policy for unsupported storage. WAL requires local host access and still allows only one writer. See [SQLite WAL](https://sqlite.org/wal.html).
- [x] Choose and document durability per table: durable user edits and job transitions, with explicitly disposable telemetry. Never silently claim a failed write was saved.
- [x] Exercise locked database, second application instance, read-only directory, disk full, failed commit, interrupted migration, corruption, deleted storage, and permission changes. Preserve existing files; offer retry, explicit recovery, or clearly labeled temporary mode without silently resetting data.
- [ ] Back up with SQLite-aware operations rather than copying an open database alone. Test restore compatibility. Close database, rows, and file handles before Windows rename/delete; never delete active WAL or SHM files. (pending: `VACUUM INTO` backup and restore are verified; closing handles before a Windows rename/delete needs a native host)
- [x] Define retention and cleanup in bounded batches; do not run large vacuum or maintenance jobs during an interaction. Test cancellation and shutdown with writes in flight.

Reference: [modernc SQLite driver](https://pkg.go.dev/modernc.org/sqlite) and [SQLite PRAGMAs](https://sqlite.org/pragma.html).

- [x] Store sources, sessions, ordered entries, and committed ingestion checkpoints. Insert entries and advance a source checkpoint in one transaction to avoid gaps after restart.
- [x] Identify sources by file identity and generation as well as path and offset. Define restart replay/deduplication for rotation and truncation; a reused path is not automatically the same file.
- [x] Index session/source/severity and ordered entry IDs. Establish default row/byte/age retention limits and prune in batches. Keep UI cursors valid or explain when retained history expired.
- [x] Test storage falling behind readers, full disk, shutdown with partial lines, and exported files failing midway. Preserve retrievable file offsets and report any unrecoverable loss.

Evidence: `go test -p 2 ./live-log-viewer/store/...` and
`go test -race -p 2 ./live-log-viewer/store/...` pass (29.7 s race run).
Failure cases, migrations, seeding, backup, retention, export atomicity, and
restart dedup are listed with test names in
`temp/live-log-viewer-core-evidence.md`. Durability is documented in the
`store` package comment. The UI maps `store.ErrReadOnly` to a labeled
temporary database in `main.go` and shows `Expired` as a footer note.

## Phase 4: Pagination and bounded rendering

- [x] Use indexed keyset pagination with a stable entry-ID tie-breaker, initially 200 entries per page. Freeze a high-water mark for browsing history while new entries arrive.
- [x] Render only visible fixed-height summary rows with overscan; open multiline content in an entry-detail view. Do not assume arbitrary variable-height virtualization already works in ownframe.
- [x] Debounce searches, cancel old queries, and attach filter generations to results. Handle retention removing cursor rows, no matches, deleted sources, and first/last boundaries.
- [x] Follow the newest page only when follow mode is enabled. Preserve the reading anchor during new inserts, page changes, and window resize.

Evidence: `ui/feed_page.go` maps `Query` to `store.Page` with
`MaxID = HighWater` while browsing and `MaxID = 0` live. UI tests cover
keyset paging, stale page and detail discard, filter generations, retention
expiry, no matches, deleted-source fallback, follow/pause, and resize
anchoring (`TestResizeKeepsAnchor`, `TestResizeAtBottomStaysBottom`).
`TestRenderLongRowStaysFixedHeight` asserts every rendered row box is
exactly 22 px with long Unicode, tabs, and unbroken text.

## Phase 5: Windows and Ebiten integration

- [x] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling.
- [x] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [ ] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback. (pending: the live dot rebinds after a redraw, the nil-display path is a no-op, and paint-only versus geometry is split, all verified headless; the per-frame replay measurement needs a GPU window session)
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input. (pending: the 1000x400 minimum and resize anchoring are verified headless; the Windows scenarios need a native host)
- [ ] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database. (pending: `TestCloseStopsWorker` cancels in-flight work inside the 2 s budget and core shutdown-with-partial tests pass; close during migration was not exercised)
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing. (pending: native Windows host)
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges. (pending: the adapter compiles under `GOOS=windows go vet`; user-profile paths, locked files, and access denied were not executed)
- [x] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview.

Evidence: `main.go` and `main_ingest.go` open the store, ensure sources,
run one ingestor per source, and pass the adapter to `ui.New`. The tick
drains at most 8 results or 2 ms, discards stale generations, and coalesces
snapshots when the request channel is full (`-perf` shows the count). No
`-web` mode exists; the README states why.

References: [ownframe frames](../../../documentation/frames.md), [features](../../../documentation/features.md), and [Ebitengine lifecycle](https://ebitengine.org/en/documents/cheatsheet.html).

## Phase 6: Verification and completion

- [x] Write focused tests for domain behavior, migration/restart recovery, persistence failures, stale asynchronous results, and shutdown. Run concurrency tests with fake collectors, files, or HTTP services; keep clipboard tests on `clipboard.UseMemory`.
- [ ] Run the finished example from a fresh data directory using dummy data, then with its real source. Exercise the full user workflow, restart, and verify saved state in the actual database. (pending: the headless real-store workflow and the core fresh-directory and restart tests pass; running the window binary needs a display)
- [ ] Profile representative and stress workloads. Record CPU, heap, RSS, allocations, goroutines, queue depth, dropped/coalesced updates, and frame/input timings. After warmup, a 30-minute mixed-workload soak must show bounded history and no sustained memory or goroutine growth. (pending: 30-minute soak not run)
- [ ] Target p99 UI frame work below 16.67 ms at 60 Hz and input feedback within 100 ms on a named reference machine. These are acceptance targets, not measured claims. Report failures and renderer limitations explicitly. (pending: real profiling not run)
- [x] Run `gofmt` on changed Go files, keep each Go file at most 2000 characters, run `make test` and `make build` from the repository root, and run targeted race tests where supported. Store generated evidence under ignored `temp/`.
- [ ] Add usage, dummy/real mode, storage location, recovery, platform limits, and measured performance to the example README and examples index. (pending: `examples/live-log-viewer/README.md` is complete; the `examples/readme.md` index line was not added because this agent does not own that file)
- [ ] Mark complete only when the example works end to end, persistence survives restart, required Windows scenarios pass, and all earlier phase evidence is recorded. Record unavailable checks as pending. (pending: native Windows scenarios, the 30-minute soak, and real profiling)

Evidence: `make build BUILD_P=2` (vet, no linking) is clean and
`make test TEST_P=4` exits 0 across 61 packages. Race tests pass on the four
core packages. The UI suite is headless: paging, stale results, filter
generations, anchor preservation, follow/pause, export bounds, fixed row
height, the retained-operation pulse, and close behavior. One redraw of 200
rows at 1100x720 measured about 9.7 ms headless
(`TestMeasureRedrawCost`).

## Evidence log

- 2026-10-07, revision aabacd7, `go test -p 2 ./live-log-viewer/ui/` on the
  UI shell: window math, debounce, follow/pause, headless render tests
  pass. Artifact: `temp/live-log-viewer-ui-evidence.md`.
- 2026-10-07, revision 55039d1 plus ff243d6, core suite: 10,000-entry dummy
  fixture, reader/parser edges, migrations, failure cases, retention,
  export, restart dedup, race run. Platform: Linux headless, Go 1.26.4.
  Artifact: `temp/live-log-viewer-core-evidence.md`.
- 2026-10-07, revision 009e2bc, joint gate:
  `go vet -p 2 ./live-log-viewer/...` clean; `go test -p 2
  ./live-log-viewer/...` passes for all five packages, including the store
  adapter tests (`TestStoreFeedDummyPaging`, `TestStoreFeedSeverityFilter`,
  `TestStoreFeedNoMatches`, `TestStoreFeedDetailAndExport`,
  `TestStoreFeedSettingsPersist`).
- 2026-10-07, revision bc5860f, `TestMeasureRedrawCost`: 3 full redraws of
  200 rows took 29.1 ms total, 9.7 ms for the last. Linux headless layout,
  no GPU window.
- 2026-10-07, revision 009e2bc, `make build BUILD_P=2` clean and
  `make test TEST_P=4` exit 0 (61 packages ok). `wc -m` keeps every Go file
  in `examples/live-log-viewer/` at or below 2000 characters.
- Pending and not logged as passing: native Windows session (scaling,
  monitor moves, minimize/restore, focus loss, rotation under a writer),
  30-minute mixed-workload soak, and p99 frame profiling.
