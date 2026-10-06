# Spreadsheet checklist

Recorded 2026-10-07. Status: planned. All implementation items remain unchecked.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A working local spreadsheet with editable cells, basic formulas, range selection, undo/redo, CSV import/export, and durable workbooks.

## Theme

Minimal document theme: a plain grid, clear row/column headers, a small formula bar, and restrained selection highlights. Provide light and dark modes with a persistent toggle.

## Phase 1: Dummy data and runnable foundation

- [ ] Create `examples/spreadsheet/` with workbook, formula, storage, and UI packages. Define sparse cell coordinates, stable workbook/sheet IDs, cell types, and edit commands.
- [ ] Start with an idempotently seeded three-sheet dummy workbook of 200 rows by 20 columns, containing text, numbers, formulas, Unicode, blank cells, and intentional formula errors. Add a 100,000-row sparse stress workbook.
- [ ] Define the initial formula language explicitly: arithmetic, cell references, rectangular ranges, SUM, and AVERAGE. Keep unsupported functions visible as errors.

## Phase 2: Core behavior and edge cases

- [ ] Implement cell editing, formula bar, keyboard navigation, rectangular selection, copy/paste, sheet switching, and bounded undo/redo.
- [ ] Handle Escape/cancel, Enter/commit, pasted multi-cell ranges, blank versus zero, Unicode, negative/large numbers, invalid references, division by zero, and cyclic formulas.
- [ ] Recalculate only affected dependencies with limits on formula length, range size, graph depth, and evaluation work. Detect cycles and discard calculations for superseded workbook revisions.
- [ ] Implement CSV import/export off the UI loop with previews, quoted fields, embedded newlines, BOM, unequal row lengths, and explicit type/formula interpretation. Commit imports atomically or leave the existing workbook intact.
- [ ] Preserve edited-but-unsaved content on save failure and keep selection stable while scrolling or changing themes.

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

- [ ] Store workbooks, sheets, sparse cells, preferences, and workbook revisions. Enforce a unique sheet/row/column key and foreign-key cleanup. Persist source formulas; treat calculated caches as invalidatable.
- [ ] Batch edits in transactions, acknowledge saves only after commit, and bound the unsaved edit queue without dropping user changes. Define restart behavior for undo history; do not imply volatile undo survives restart.
- [ ] Handle two windows editing one workbook with revision conflict detection. Test imports interrupted halfway, workbook deletion during pending save, and crash recovery between edit and save acknowledgement.

## Phase 4: Pagination and bounded rendering

- [ ] Use viewport-based row and column windows with overscan, stable coordinates, and full-content spacers. Fetch bounded rectangular cell ranges from SQLite; never materialize the whole sheet to render one screen.
- [ ] Preserve focus, active editor text, range endpoints, frozen headers, and selection across viewport replacement. Verify horizontal and vertical scrolling and keyboard jumps to offscreen cells.
- [ ] Use indexed keyset pages for workbook lists and import previews, initially 50 records per page. Do not split spreadsheet navigation into arbitrary pages.
- [ ] Cancel stale range loads, cache a bounded number of tiles, handle empty sheets and sheet-size changes, and verify lookup query plans for deep row positions.

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
