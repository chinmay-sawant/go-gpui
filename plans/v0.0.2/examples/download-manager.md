# Download manager checklist

Recorded 2026-10-07. Status: implemented and verified headlessly; Windows-host,
soak, and profiling items remain pending.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A working local download queue with bounded concurrent transfers, progress and speed, cancellation, supported resume, and durable job history.

## Theme

Minimal utility theme: a plain queue table, compact progress bars, and simple job details. Provide light and dark modes with a persistent toggle. Status remains readable without relying on color.

## Phase 1: Dummy data and runnable foundation

- [x] Create `examples/download-manager/` with transfer, scheduler, store, and UI packages. Define explicit queued/running/paused/completed/failed/cancelled states and valid transitions.
- [x] Default to dummy jobs and a deterministic fake transport: 100 history entries, queued jobs, unknown lengths, failures, and variable progress. No external download should start on first launch.
- [x] Provide a local HTTP fixture service for successful, slow, chunked, range-capable, changing-content, redirect, and interrupted transfers. Give dummy files a dedicated temporary directory.

Evidence (2026-10-07, rev 7a3a4b6): `go test -p 2 -timeout 300s ./download-manager/...` passes. `SeedDummy(ctx, 100)` builds 70 completed, 10 failed, 5 cancelled, 8 queued, 7 paused rows; the fake transport opens no socket and writes under `transfer.DummyDir()`; the fixture serves ok, slow, chunked, range, changing, redirect, interrupt, and status routes. The UI integration test `TestEndToEndDummy` reads the seeded history through the page.

## Phase 2: Core behavior and edge cases

- [x] Implement add job, bounded worker scheduling, progress, pause/resume where supported, cancel, retry, destination selection, and completed-job history.
- [x] Use cancellable requests with connection/read limits and streamed writes. Keep progress events coalesced and nonblocking; durable completion and failure events must not be dropped.
- [x] Resume only after validating HTTP 206, Content-Range, and saved validators with If-Range. Handle servers ignoring Range with 200, 416, changed ETag/Last-Modified, content encoding, and missing Content-Length without appending corrupt data.
- [ ] Handle redirects, HTTP errors, stalls, counter overflow, negative/unknown totals, checksum mismatch when a checksum is supplied, and elapsed-time gaps after suspend. (pending: the byte-counter overflow guard in `transfer/http_stream.go` is code-reviewed but has no dedicated test; redirects, 503, stalls, unknown totals, checksum mismatch, and suspend gaps are tested)
- [ ] Use safe destination names independent of untrusted URL paths. Handle Windows reserved names, separators, Unicode, case-insensitive collisions, long paths, disk full, and existing destinations. (pending: real disk-full `ENOSPC` is not reproduced; reserved names, separators, Unicode, collisions, long names, and existing destinations are tested)
- [ ] Write a partial file in the destination directory and explicitly finalize after successful validation. On Windows close handles before rename, handle sharing violations and antivirus locks with bounded retries, and avoid silently overwriting existing files. (pending: native Windows close-before-rename and lock retries; the partial-to-final finalize, refused overwrite, and unwritable-partial paths are tested)

Evidence (2026-10-07, rev 7a3a4b6): 26 transfer tests plus 5 fixture tests pass, including 206 resume with Content-Range and If-Range, 200 restart, 416 completion and refusal, changed ETag restart, redirect, 503, interrupt, stall, chunked unknown length, content-encoding refusal, checksum mismatch, and safe-name/destination cases. The UI test `TestControlFlowFixture` drives pause, resume, and cancel against fixture `/slow` through the page. Not covered: real `ENOSPC` and native Windows rename behavior.

## Phase 3: SQLite storage and failure handling

- [x] Use the examples module's existing `modernc.org/sqlite` dependency through `database/sql`; keep storage code out of the root public library.
- [ ] Give this example a separate database under `os.UserConfigDir()/ownframe/<example>/`. Support an explicit data-directory override. Use `filepath` and correctly encoded driver paths for Windows drive letters, spaces, Unicode, and URI punctuation. (pending: the `file:` URI encoding in `store/dsn.go` is code-reviewed but has no dedicated test; the explicit override is exercised by every integration test)
- [x] Implement versioned, transactional migrations and idempotent dummy seeding with a seed-version marker. Reopening must not duplicate fixtures or overwrite user changes. Reject a newer unsupported schema without modifying it.
- [x] Start with one serialized database worker and one connection; apply foreign keys and a finite busy timeout to every physical connection if the pool grows. Keep transactions short and close rows before issuing more queries. Test single-connection self-deadlocks and separate `:memory:` databases across connections.
- [x] Use parameterized statements, explicit constraints, query deadlines, and bounded operation queues. Retry only appropriate lock failures within a deadline; never blindly replay non-idempotent operations.
- [x] For local writable storage, verify WAL activation and manage checkpoint work off the UI loop. Support a documented rollback-journal policy for unsupported storage. WAL requires local host access and still allows only one writer. See [SQLite WAL](https://sqlite.org/wal.html).
- [x] Choose and document durability per table: durable user edits and job transitions, with explicitly disposable telemetry. Never silently claim a failed write was saved.
- [ ] Exercise locked database, second application instance, read-only directory, disk full, failed commit, interrupted migration, corruption, deleted storage, and permission changes. Preserve existing files; offer retry, explicit recovery, or clearly labeled temporary mode without silently resetting data. (pending: disk full, failed commit, interrupted migration, deleted storage under a live handle, and permission changes mid-run; locked database, second instance, unwritable directory, corrupt file, and newer schema are tested)
- [ ] Back up with SQLite-aware operations rather than copying an open database alone. Test restore compatibility. Close database, rows, and file handles before Windows rename/delete; never delete active WAL or SHM files. (pending: native Windows close-before-rename; `VACUUM INTO` backup and restore are tested)
- [ ] Define retention and cleanup in bounded batches; do not run large vacuum or maintenance jobs during an interaction. Test cancellation and shutdown with writes in flight. (pending: checkpoint cancellation during an in-flight write has no dedicated test; bounded cleanup batches and scheduler shutdown with a persisted checkpoint are tested)

Reference: [modernc SQLite driver](https://pkg.go.dev/modernc.org/sqlite) and [SQLite PRAGMAs](https://sqlite.org/pragma.html).

- [x] Store job IDs, URL, destination, state, observed progress, expected length, content validators, attempts, and timestamps. Redact credentials from UI/log output and define what secrets, if any, may be persisted.
- [x] Persist job creation before starting transfer, batch progress checkpoints, and save state transitions durably. Progress saved in SQLite must not be treated as proof that those bytes exist in the partial file.
- [x] On restart reconcile database state with actual files and validators. Interrupted running jobs become recoverable, not completed. Handle a missing/shorter/larger partial file, deleted final file, and reused destination.
- [x] Define recovery for a crash between file finalization and database completion using stable job IDs and validation. No atomic transaction spans SQLite and the filesystem; reconciliation must close that gap.
- [x] Serialize destination ownership to prevent two jobs writing the same file. Prevent a second instance from scheduling the same active jobs through an explicit instance lock or tested ownership protocol.

Evidence (2026-10-07, rev 7a3a4b6): 26 store tests pass, covering open/migrate/reopen, WAL activation, `OpenMemory` isolation, idempotent seeding with user-edit preservation, keyset paging of 120 rows as 50/50/20, summary counts, bounded cleanup, every reconcile case, locked database deadline, newer-schema rejection, corrupt-file and unwritable-directory refusal, `VACUUM INTO` backup plus restore, instance lock and stale-lock reclaim, and 20 goroutines on the one worker. `TestRestartKeepsData` (UI/wire) reopens the directory and finds the completed row and the saved theme. Storage policy is in the `store` package comment.

## Phase 4: Pagination and bounded rendering

- [x] Keep active jobs in a bounded in-memory view and page durable history with indexed timestamp plus stable ID ordering, initially 50 jobs.
- [x] Separate active progress sorting from stable history paging. Preserve selection by job ID, reset cursors for changed filters, and discard old page responses.
- [x] Handle jobs completing or being removed while details are open, empty histories, shrinking final pages, and first/last boundaries. Render only the selected page or visible rows.
- [x] Query summary counts separately and at a bounded refresh rate; do not scan the full history for each progress event.

Evidence (2026-10-07, rev 7a3a4b6): `ui` tests cover stale page discard, filter reset with selection kept, walk-back from an empty shrinking page, and first/last boundaries. `ui/wire` pages history through `store.History` with an opaque `(UpdatedAt, ID)` cursor, 50 rows per page; the filtered views scan keyset pages and keep the cursor real. Summary counts are requested every 2 s and never per progress event. The details pane survives page changes and updates from both active and history rows.

## Phase 5: Windows and Ebiten integration

- [x] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling.
- [x] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [x] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback.
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input. (pending: native Windows host; the 780x520 minimum and page scrolling are in place)
- [x] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database.
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing. (pending: no Windows host on this machine)
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges. (pending: native Windows host; the name and path adapters are unit-tested)
- [x] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview. (not added: the example is desktop-only by scope, and `main.go` has no `-web` flag)

Evidence (2026-10-07, rev 7a3a4b6): `main.go` calls `ownframe.Run`; `ui` holds the view, pager, bindings, and window state; `wire` runs one worker goroutine and a 64-command queue, answers through a 512-result channel, and the tick drains at most 64 results (`TestTickDrainBudget`). Request generations discard stale pages and summaries (`TestTickDiscardsStalePage`, `TestTickDiscardsStaleSummary`); rejected requests are retried (`TestPumpRetriesDroppedRequests`). Retained operations rebind on generation change (`TestPaintRebindsAfterRedraw`), and the nil display list is handled (`TestPaintWithoutDisplay`). Benchmarks: paint-only tick 3,965 ns, relayout tick 15.8 ms, full Redraw 21.9 ms (20 queue rows, 50 history rows, i7-13700HX). Shutdown is one second for the worker plus six for the engine (five inside), documented in the README; the integration tests close cleanly and `TestShutdownCancelsAndPersists` covers the engine. The full Ebiten draw cost and frame pacing were not measured (headless).

## Phase 6: Verification and completion

- [x] Write focused tests for domain behavior, migration/restart recovery, persistence failures, stale asynchronous results, and shutdown. Run concurrency tests with fake collectors, files, or HTTP services; keep clipboard tests on `clipboard.UseMemory`.
- [x] Run the finished example from a fresh data directory using dummy data, then with its real source. Exercise the full user workflow, restart, and verify saved state in the actual database.
- [ ] Profile representative and stress workloads. Record CPU, heap, RSS, allocations, goroutines, queue depth, dropped/coalesced updates, and frame/input timings. After warmup, a 30-minute mixed-workload soak must show bounded history and no sustained memory or goroutine growth. (pending: 30-minute soak and runtime sampling)
- [ ] Target p99 UI frame work below 16.67 ms at 60 Hz and input feedback within 100 ms on a named reference machine. These are acceptance targets, not measured claims. Report failures and renderer limitations explicitly. (pending: measured relayout tick is 15.8 ms and full Redraw 21.9 ms, so the target is at risk; GPU frame timing and input latency were not measured headless)
- [x] Run `gofmt` on changed Go files, keep each Go file at most 2000 characters, run `make test` and `make build` from the repository root, and run targeted race tests where supported. Store generated evidence under ignored `temp/`.
- [x] Add usage, dummy/real mode, storage location, recovery, platform limits, and measured performance to the example README and examples index. (2026-10-07: `examples/download-manager/README.md` has all sections; the index row landed in `examples/readme.md`)
- [ ] Mark complete only when the example works end to end, persistence survives restart, required Windows scenarios pass, and all earlier phase evidence is recorded. Record unavailable checks as pending. (pending: Windows scenarios, soak, and profiling; the headless workflow is complete)

## Evidence log

- 2026-10-07, rev 7a3a4b6, Linux WSL2 amd64, headless. `cd examples && go vet -p 2 ./download-manager/...` clean; `go test -p 2 -timeout 300s ./download-manager/...` all packages pass; `go test -race -p 1 ./download-manager/ui/...` clean. Root `make build BUILD_P=4` and `make test TEST_P=4` pass (63 packages ok).
- Phase 1: core 37 tests (domain, transfer, fixture) plus UI `TestEndToEndDummy`; 100-row seed and fixture routes; pass.
- Phase 2: core 26 transfer + 5 fixture tests plus UI `TestControlFlowFixture`; pass; `ENOSPC` and native Windows rename pending.
- Phase 3: core 26 store tests plus UI `TestRestartKeepsData`; pass; disk-full, failed commit, interrupted migration, deleted storage, permission changes, and native Windows close-before-rename pending.
- Phase 4: UI pager and detail tests (`pager_test.go`, `pager_walk_test.go`, `detail_test.go`); 50-row keyset pages; pass.
- Phase 5: UI integration tests `TestEndToEndDummy`, `TestEndToEndFixture`, `TestControlFlowFixture`, `TestRestartKeepsData`; benchmarks 3,965 ns / 15.8 ms / 21.9 ms; pass; native Windows scenarios and frame pacing pending.
- Phase 6: `make test`/`make build` clean; `gofmt` and the 2000-character limit hold on every file; soak, profiling, and the frame budget remain pending. Evidence files: `temp/download-manager-ui-evidence.md`, `temp/download-manager-core-evidence.md`, `temp/download-manager-integration.md`.

For each phase, add date, revision, command or manual workflow, dataset size, platform/renderer, result, and log or screenshot path here. Leave unsupported or unrun scenarios unchecked.
