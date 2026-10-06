# Download manager checklist

Recorded 2026-10-07. Status: planned. All implementation items remain unchecked.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A working local download queue with bounded concurrent transfers, progress and speed, cancellation, supported resume, and durable job history.

## Theme

Minimal utility theme: a plain queue table, compact progress bars, and simple job details. Provide light and dark modes with a persistent toggle. Status remains readable without relying on color.

## Phase 1: Dummy data and runnable foundation

- [ ] Create `examples/download-manager/` with transfer, scheduler, store, and UI packages. Define explicit queued/running/paused/completed/failed/cancelled states and valid transitions.
- [ ] Default to dummy jobs and a deterministic fake transport: 100 history entries, queued jobs, unknown lengths, failures, and variable progress. No external download should start on first launch.
- [ ] Provide a local HTTP fixture service for successful, slow, chunked, range-capable, changing-content, redirect, and interrupted transfers. Give dummy files a dedicated temporary directory.

## Phase 2: Core behavior and edge cases

- [ ] Implement add job, bounded worker scheduling, progress, pause/resume where supported, cancel, retry, destination selection, and completed-job history.
- [ ] Use cancellable requests with connection/read limits and streamed writes. Keep progress events coalesced and nonblocking; durable completion and failure events must not be dropped.
- [ ] Resume only after validating HTTP 206, Content-Range, and saved validators with If-Range. Handle servers ignoring Range with 200, 416, changed ETag/Last-Modified, content encoding, and missing Content-Length without appending corrupt data.
- [ ] Handle redirects, HTTP errors, stalls, counter overflow, negative/unknown totals, checksum mismatch when a checksum is supplied, and elapsed-time gaps after suspend.
- [ ] Use safe destination names independent of untrusted URL paths. Handle Windows reserved names, separators, Unicode, case-insensitive collisions, long paths, disk full, and existing destinations.
- [ ] Write a partial file in the destination directory and explicitly finalize after successful validation. On Windows close handles before rename, handle sharing violations and antivirus locks with bounded retries, and avoid silently overwriting existing files.

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

- [ ] Store job IDs, URL, destination, state, observed progress, expected length, content validators, attempts, and timestamps. Redact credentials from UI/log output and define what secrets, if any, may be persisted.
- [ ] Persist job creation before starting transfer, batch progress checkpoints, and save state transitions durably. Progress saved in SQLite must not be treated as proof that those bytes exist in the partial file.
- [ ] On restart reconcile database state with actual files and validators. Interrupted running jobs become recoverable, not completed. Handle a missing/shorter/larger partial file, deleted final file, and reused destination.
- [ ] Define recovery for a crash between file finalization and database completion using stable job IDs and validation. No atomic transaction spans SQLite and the filesystem; reconciliation must close that gap.
- [ ] Serialize destination ownership to prevent two jobs writing the same file. Prevent a second instance from scheduling the same active jobs through an explicit instance lock or tested ownership protocol.

## Phase 4: Pagination and bounded rendering

- [ ] Keep active jobs in a bounded in-memory view and page durable history with indexed timestamp plus stable ID ordering, initially 50 jobs.
- [ ] Separate active progress sorting from stable history paging. Preserve selection by job ID, reset cursors for changed filters, and discard old page responses.
- [ ] Handle jobs completing or being removed while details are open, empty histories, shrinking final pages, and first/last boundaries. Render only the selected page or visible rows.
- [ ] Query summary counts separately and at a bounded refresh rate; do not scan the full history for each progress event.

## Phase 5: Windows and Ebiten integration

- [ ] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling.
- [ ] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [ ] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback.
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input.
- [ ] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database.
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing.
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges.
- [ ] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview.

References: [ownframe frames](../../../../documentation/frames.md), [features](../../../../documentation/features.md), and [Ebitengine lifecycle](https://ebitengine.org/en/documents/cheatsheet.html).

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
