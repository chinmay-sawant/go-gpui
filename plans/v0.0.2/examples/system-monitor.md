# System monitor checklist

Recorded 2026-10-07. Status: implemented and verified headless; native Windows,
soak, and profiling checks pending. Each phase carries a short evidence note,
and the Evidence log records commands, revisions, and results.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A responsive desktop monitor that shows CPU, memory, disk and network activity, plus a searchable process table and process details. Collectors run in goroutines and communicate through bounded channels.

## Theme

Minimal utility theme: plain panels, compact tables, restrained line graphs, one accent color, and readable numeric labels. Provide light and dark modes with a persistent toggle. Use textual values alongside color.

## Phase 1: Dummy data and runnable foundation

Headless evidence 2026-10-07 (rev e9207c5, linux/amd64): `go vet -p 2 ./system-monitor/...` and `go test -p 2 -count=1 ./system-monitor/...` clean in the worktree; dummy fixture tests cover determinism, 10,000-process stress, and 120 historical samples (core evidence); the UI renders the overview to 418 display operations and 432 boxes at 1024x700 and the layout/click tests pass.

- [x] Create `examples/system-monitor/` with separate collector, domain, storage, and UI packages inside the existing examples module. Define collector interfaces and immutable snapshots before wiring workers.
- [x] Default to labeled dummy mode with a fixed seed, 250 processes, 120 historical samples, CPU spikes, unavailable sensors, and process exits. Add a stress fixture with 10,000 processes; fixtures must not access the host or mix with real history.
- [x] Make the first runnable screen useful without permissions or platform collectors. Define sample units, timestamps, process identity, missing values, and cancellation contracts. (headless: page render and clicks verified; native window launch pending)

## Phase 2: Core behavior and edge cases

Headless evidence 2026-10-07 (rev e9207c5): domain, pool, and manager tests pass, including first sample, zero elapsed, counter reset, suspend and clock-step gaps, hot-plug, PID reuse, detail identity, and slow-source skipping; UI tests cover overview, search/sort, detail, and the dummy-to-live reset (`go test -race -p 2 ./system-monitor/ui`).

- [x] Implement overview, process search/sort, and process-detail views. Start with read-only monitoring; process termination is outside this initial scope.
- [x] Sample summary counters about once a second and expensive process details less often. Use a bounded collector pool, deadlines, and latest-snapshot delivery; never start a goroutine per process on every sample.
- [x] Use fixed-size graph buffers, monotonic elapsed time for rates, and explicit gaps after suspend/resume. Handle first samples, zero elapsed time, counter reset/wrap, CPU count changes, and network interface hot-plug.
- [x] Handle processes exiting during reads, PID reuse with start-time identity, access denied, unavailable Windows counters, and slow collectors. Show unavailable values instead of misleading zeroes.
- [x] Separate system measurements from this Go process runtime metrics. Switching from dummy to live must reset baselines and graph state.

## Phase 3: SQLite storage and failure handling

Headless evidence 2026-10-07 (rev e9207c5): 36 storage tests pass, covering migrations, newer-schema refusal, idempotent seed, keyset history, retention, recorder failure and deadline, locked database, read-only directory, corruption, interrupted migration, and backup/restore; `TestFileDSN` covers Windows drive, UNC, space, Unicode, and URI punctuation paths; `TestConcurrentReadWrite` and `TestTempStoresAreSeparate` cover the single connection.

- [x] Use the examples module's existing `modernc.org/sqlite` dependency through `database/sql`; keep storage code out of the root public library.
- [x] Give this example a separate database under `os.UserConfigDir()/ownframe/<example>/`. Support an explicit data-directory override. Use `filepath` and correctly encoded driver paths for Windows drive letters, spaces, Unicode, and URI punctuation.
- [x] Implement versioned, transactional migrations and idempotent dummy seeding with a seed-version marker. Reopening must not duplicate fixtures or overwrite user changes. Reject a newer unsupported schema without modifying it.
- [x] Start with one serialized database worker and one connection; apply foreign keys and a finite busy timeout to every physical connection if the pool grows. Keep transactions short and close rows before issuing more queries. Test single-connection self-deadlocks and separate `:memory:` databases across connections.
- [x] Use parameterized statements, explicit constraints, query deadlines, and bounded operation queues. Retry only appropriate lock failures within a deadline; never blindly replay non-idempotent operations. (lock waits use the finite `busy_timeout`; no replay path exists)
- [x] For local writable storage, verify WAL activation and manage checkpoint work off the UI loop. Support a documented rollback-journal policy for unsupported storage. WAL requires local host access and still allows only one writer. See [SQLite WAL](https://sqlite.org/wal.html).
- [x] Choose and document durability per table: durable user edits and job transitions, with explicitly disposable telemetry. Never silently claim a failed write was saved.
- [ ] Exercise locked database, second application instance, read-only directory, disk full, failed commit, interrupted migration, corruption, deleted storage, and permission changes. Preserve existing files; offer retry, explicit recovery, or clearly labeled temporary mode without silently resetting data. (pending: disk full, deleted storage while open, and permission changes are not portable to force on this host; the other cases are covered)
- [ ] Back up with SQLite-aware operations rather than copying an open database alone. Test restore compatibility. Close database, rows, and file handles before Windows rename/delete; never delete active WAL or SHM files. (pending: native Windows rename/delete not run; backup and restore are covered)
- [x] Define retention and cleanup in bounded batches; do not run large vacuum or maintenance jobs during an interaction. Test cancellation and shutdown with writes in flight.

Reference: [modernc SQLite driver](https://pkg.go.dev/modernc.org/sqlite) and [SQLite PRAGMAs](https://sqlite.org/pragma.html).

- [x] Store settings, saved views, recording sessions, and downsampled metric history. Live process tables stay in memory unless explicitly recording; do not write every process on every tick.
- [x] Index history by session, metric, and timestamp plus stable ID. Set a default 24-hour raw-history retention and a bounded aggregate policy. Handle clock changes, duplicate sample IDs, retention during history reads, and recording write failure.

## Phase 4: Pagination and bounded rendering

Headless evidence 2026-10-07 (UI commits 8c87388 to 1a2b242): `go test -race -p 2 ./system-monitor/ui` passes; table tests cover frozen snapshots, 50-row pages at 10,000 processes, filter reset, shrinking last page, identity tie-break, and selection kept by identity; graph tests cover the 120-sample ring and peak-preserving downsampling to at most 96 columns.

- [x] Page processes in stable snapshots, initially 50 rows per page. Tie-break sort keys with process identity; retain selection by identity when processes arrive or exit.
- [x] Freeze the displayed snapshot while navigating its pages, show the sample timestamp, and let refresh replace it explicitly. Avoid duplicate/skipped rows caused by continuously changing CPU sorting.
- [x] Reset pagination after filter changes; discard stale results, handle an empty or shrinking last page, and hide Previous/Next at boundaries. Limit rendered rows independently of total process count.
- [x] Read historical recordings with indexed keyset pagination and bounded time-range queries. Downsample graphs to the visible pixel width instead of loading all samples.

## Phase 5: Windows and Ebiten integration

Headless evidence 2026-10-07 (UI commits, rev e9207c5): tick tests cover reacquire after redraw and resize, bitmap fallback, paint-only updates, and stale-generation discard; benchmarks measure the active tick at about 2 to 4 us and a full overview repaint at about 14 to 30 ms; `GOOS=windows go vet ./system-monitor/...` is clean; no GUI was launched.

- [x] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling. (headless: launch code and loop ownership verified; native window launch pending)
- [x] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [x] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback.
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input. (pending: native Windows session for scaling, monitor move, minimize/restore, and focus loss; minimum size, headless resize, keyboard navigation, and wheel scrolling are covered by code and tests)
- [ ] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database. (pending: close during migration not exercised; close during write and active work are covered)
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing. (pending: no Windows host on this machine)
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges. (pending: native Windows user-profile and access-denied cases; adapters cross-compile and locked/read-only tests pass)
- [x] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview. (not added: no `-web` mode in this example; the README documents the limit)

References: [ownframe frames](../../../documentation/frames.md), [features](../../../documentation/features.md), and [Ebitengine lifecycle](https://ebitengine.org/en/documents/cheatsheet.html).

## Phase 6: Verification and completion

Headless evidence 2026-10-07 (rev e9207c5 + UI commits): `make test TEST_P=4` and `make build BUILD_P=4` from the worktree root exit 0; `go test -race -p 2 ./system-monitor/ui` passes; every Go file is at most 2000 characters and `gofmt -l` is clean; no GUI, soak, or profiler ran.

- [x] Write focused tests for domain behavior, migration/restart recovery, persistence failures, stale asynchronous results, and shutdown. Run concurrency tests with fake collectors, files, or HTTP services; keep clipboard tests on `clipboard.UseMemory`.
- [ ] Run the finished example from a fresh data directory using dummy data, then with its real source. Exercise the full user workflow, restart, and verify saved state in the actual database. (pending: no GUI on this host; the storage, collector, and UI flow is verified headless and persistence survives reopen in tests)
- [ ] Profile representative and stress workloads. Record CPU, heap, RSS, allocations, goroutines, queue depth, dropped/coalesced updates, and frame/input timings. After warmup, a 30-minute mixed-workload soak must show bounded history and no sustained memory or goroutine growth. (pending: 30-minute soak and profiling not run)
- [ ] Target p99 UI frame work below 16.67 ms at 60 Hz and input feedback within 100 ms on a named reference machine. These are acceptance targets, not measured claims. Report failures and renderer limitations explicitly. (pending: frame p99 needs a window session; the headless tick cost is about 2 to 4 us)
- [x] Run `gofmt` on changed Go files, keep each Go file at most 2000 characters, run `make test` and `make build` from the repository root, and run targeted race tests where supported. Store generated evidence under ignored `temp/`.
- [x] Add usage, dummy/real mode, storage location, recovery, platform limits, and measured performance to the example README and examples index. (2026-10-07: `examples/system-monitor/README.md` is complete; the index row landed in `examples/readme.md`)
- [ ] Mark complete only when the example works end to end, persistence survives restart, required Windows scenarios pass, and all earlier phase evidence is recorded. Record unavailable checks as pending. (pending: native Windows scenarios and a real window run)

## Evidence log

Platform for every entry: linux/amd64, 13th Gen Intel Core i7-13700HX,
headless (no window launched), Go 1.26.4. Commands run from `examples/`
unless noted. Artifacts: `temp/system-monitor-core-evidence.md`,
`temp/system-monitor-ui-evidence.md`, `temp/system-monitor-core-done.md`,
`temp/system-monitor-ui-done.md`.

- Phase 1, 2026-10-07, core rev e9207c5 + UI commits: whole-example vet and
  tests clean; dummy determinism, 10,000-process stress, and 120-sample
  history tests pass; UI overview renders 418 operations and 432 boxes at
  1024x700. Result: packages, dummy mode, and the first screen verified
  headless; native window launch pending.
- Phase 2, 2026-10-07, core rev e9207c5 + UI commits: domain, pool, and
  manager tests pass over first sample, gaps, resets, hot-plug, and PID
  reuse; UI race tests pass for overview, search/sort, detail, and the
  dummy-to-live reset. Result: behavior and edge cases verified headless.
- Phase 3, 2026-10-07, core rev e9207c5: 36 storage tests pass over
  migrations, seed idempotence, keyset history, retention, recorder
  failure and deadline, locked database, read-only directory, corruption,
  interrupted migration, and backup/restore. Result: storage verified;
  disk full, deleted storage while open, permission changes, and native
  Windows rename/delete remain pending.
- Phase 4, 2026-10-07, UI commits: table tests pass over frozen snapshots,
  50-row pages at 10,000 processes, filter reset, shrinking last page,
  identity tie-break, and selection by identity; graph tests pass over the
  fixed ring and downsampling. Result: pagination and bounded rendering
  verified headless.
- Phase 5, 2026-10-07, UI commits + core rev e9207c5: tick tests pass for
  reacquire after redraw and resize, bitmap fallback, paint-only updates,
  and stale-generation discard; benchmarks: tick about 2 to 4 us, overview
  repaint about 14 to 30 ms, process repaint about 7 to 9 ms, stress page
  slice about 0.4 to 0.8 ms; `GOOS=windows go vet` clean. Result: integration
  verified headless; Windows session, scaling, and frame timing pending.
- Phase 6, 2026-10-07, rev e9207c5 + UI commits: `make test TEST_P=4` and
  `make build BUILD_P=4` exit 0; race tests pass; `gofmt -l` clean; every
  Go file at most 2000 characters. Result: gates pass; window run, soak,
  and profiling remain pending.
