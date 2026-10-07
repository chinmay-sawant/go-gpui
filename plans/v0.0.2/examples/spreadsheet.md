# Spreadsheet checklist

Recorded 2026-10-07. Status: Phases 1, 2, and 4 complete; Phase 3 and the
UI half of Phase 5 complete; Phase 6 partially complete. Windows host,
soak, and real profiling checks stay unchecked with pending notes.

This checklist plans a working Go-native ownframe example in the existing examples module. No JavaScript runtime is required. Complete phases in order; record commands, results, platform, workload, and artifact paths under each phase before checking it off. A compile check does not prove native window behavior.

## Working goal

A working local spreadsheet with editable cells, basic formulas, range selection, undo/redo, CSV import/export, and durable workbooks.

## Theme

Minimal document theme: a plain grid, clear row/column headers, a small formula bar, and restrained selection highlights. Provide light and dark modes with a persistent toggle.

## Phase 1: Dummy data and runnable foundation

- [x] Create `examples/spreadsheet/` with workbook, formula, storage, and UI packages. Define sparse cell coordinates, stable workbook/sheet IDs, cell types, and edit commands.
- [x] Start with an idempotently seeded three-sheet dummy workbook of 200 rows by 20 columns, containing text, numbers, formulas, Unicode, blank cells, and intentional formula errors. Add a 100,000-row sparse stress workbook.
- [x] Define the initial formula language explicitly: arithmetic, cell references, rectangular ranges, SUM, and AVERAGE. Keep unsupported functions visible as errors.

Evidence 2026-10-07: `go test -p 2 ./spreadsheet/...` passes for
`workbook` (`TestSeedDummy`, `TestSeedStress`, `TestBlankVersusZero`,
`TestCSVTypes`) and `formula` (`TestArithmetic`, `TestErrors`,
`TestRanges`). The UI test `TestUIShowsSeededWorkbook` opens the seeded
store and renders `row 1`. Core packages committed at 69d2688.

## Phase 2: Core behavior and edge cases

- [x] Implement cell editing, formula bar, keyboard navigation, rectangular selection, copy/paste, sheet switching, and bounded undo/redo.
- [x] Handle Escape/cancel, Enter/commit, pasted multi-cell ranges, blank versus zero, Unicode, negative/large numbers, invalid references, division by zero, and cyclic formulas.
- [x] Recalculate only affected dependencies with limits on formula length, range size, graph depth, and evaluation work. Detect cycles and discard calculations for superseded workbook revisions.
- [x] Implement CSV import/export off the UI loop with previews, quoted fields, embedded newlines, BOM, unequal row lengths, and explicit type/formula interpretation. Commit imports atomically or leave the existing workbook intact.
- [x] Preserve edited-but-unsaved content on save failure and keep selection stable while scrolling or changing themes.

Evidence 2026-10-07: UI tests `TestTypingEscapeAndCommit`,
`TestFormulaBarStartsEdit`, `TestSheetTabSwitches`,
`TestPasteMultiCellRange`, `TestCopySelectionTSV`,
`TestSaveFailureRestoresEditor`, `TestSelectionSurvivesThemeToggle`,
`TestWindowReplacesAndKeepsSelection`. Core tests
`TestRecalcOnlyAffected`, `TestRecalcLimits`, `TestGraphDepthLimit`,
`TestCycles`, `TestStaleRecalc`, `TestCSVBOM`, `TestCSVUnequalRows`,
`TestImportReplaceCommits`, `TestImportAtomicOnAbort`. `TestTickBudget`
and `TestStaleFetchDiscarded` cover the UI side of superseded work.

## Phase 3: SQLite storage and failure handling

- [x] Use the examples module's existing `modernc.org/sqlite` dependency through `database/sql`; keep storage code out of the root public library.
- [x] Give this example a separate database under `os.UserConfigDir()/ownframe/<example>/`. Support an explicit data-directory override. Use `filepath` and correctly encoded driver paths for Windows drive letters, spaces, Unicode, and URI punctuation.
- [x] Implement versioned, transactional migrations and idempotent dummy seeding with a seed-version marker. Reopening must not duplicate fixtures or overwrite user changes. Reject a newer unsupported schema without modifying it.
- [x] Start with one serialized database worker and one connection; apply foreign keys and a finite busy timeout to every physical connection if the pool grows. Keep transactions short and close rows before issuing more queries. Test single-connection self-deadlocks and separate `:memory:` databases across connections.
- [x] Use parameterized statements, explicit constraints, query deadlines, and bounded operation queues. Retry only appropriate lock failures within a deadline; never blindly replay non-idempotent operations.
- [x] For local writable storage, verify WAL activation and manage checkpoint work off the UI loop. Support a documented rollback-journal policy for unsupported storage. WAL requires local host access and still allows only one writer. See [SQLite WAL](https://sqlite.org/wal.html).
- [x] Choose and document durability per table: durable user edits and job transitions, with explicitly disposable telemetry. Never silently claim a failed write was saved.
- [x] Exercise locked database, second application instance, read-only directory, disk full, failed commit, interrupted migration, corruption, deleted storage, and permission changes. Preserve existing files; offer retry, explicit recovery, or clearly labeled temporary mode without silently resetting data.
- [x] Back up with SQLite-aware operations rather than copying an open database alone. Test restore compatibility. Close database, rows, and file handles before Windows rename/delete; never delete active WAL or SHM files.
- [x] Define retention and cleanup in bounded batches; do not run large vacuum or maintenance jobs during an interaction. Test cancellation and shutdown with writes in flight.

Reference: [modernc SQLite driver](https://pkg.go.dev/modernc.org/sqlite) and [SQLite PRAGMAs](https://sqlite.org/pragma.html).

- [x] Store workbooks, sheets, sparse cells, preferences, and workbook revisions. Enforce a unique sheet/row/column key and foreign-key cleanup. Persist source formulas; treat calculated caches as invalidatable.
- [x] Batch edits in transactions, acknowledge saves only after commit, and bound the unsaved edit queue without dropping user changes. Define restart behavior for undo history; do not imply volatile undo survives restart.
- [x] Handle two windows editing one workbook with revision conflict detection. Test imports interrupted halfway, workbook deletion during pending save, and crash recovery between edit and save acknowledgement.

Evidence 2026-10-07: the core's committed storage suite passes:
`TestMigrationFreshAndReopen`, `TestInterruptedMigration`,
`TestSchemaNewerRejected`, `TestSeedIdempotent`,
`TestSeedPreservesUserChanges`, `TestConcurrentReadsNoDeadlock`,
`TestMemoryIsolation`, `TestLockedDatabaseSurfacesError`,
`TestReadOnlyDirFailsCleanly`, `TestReadOnlyFileFailsOnWrite`,
`TestCorruptDatabaseUntouched`, `TestDeletedStorageReopensFresh`,
`TestBackupAndRestore`, `TestTrimRevisions`,
`TestCloseWithWriteInFlight`, `TestJournalModes`,
`TestCheckpointTruncatesWAL`, `TestRevisionConflict`,
`TestAlternatingWriters`, `TestImportAtomicOnAbort`,
`TestDeleteDuringPendingSave`, `TestCrashBetweenEditAndAck`,
`TestFailedSaveKeepsMemoryEdit`, `TestRangeQueryPlan`. The adapter keeps
its own bounded undo history and persists each step; the UI README states
that undo is volatile across restarts. The core agent's own temp evidence
files were lost with the wiped worktree, so this note cites its committed
tests, which were re-run here.

## Phase 4: Pagination and bounded rendering

- [x] Use viewport-based row and column windows with overscan, stable coordinates, and full-content spacers. Fetch bounded rectangular cell ranges from SQLite; never materialize the whole sheet to render one screen.
- [x] Preserve focus, active editor text, range endpoints, frozen headers, and selection across viewport replacement. Verify horizontal and vertical scrolling and keyboard jumps to offscreen cells.
- [x] Use indexed keyset pages for workbook lists and import previews, initially 50 records per page. Do not split spreadsheet navigation into arbitrary pages.
- [x] Cancel stale range loads, cache a bounded number of tiles, handle empty sheets and sheet-size changes, and verify lookup query plans for deep row positions.

Evidence 2026-10-07: `TestComputeWindowAtTop`,
`TestComputeWindowScrolled`, `TestComputeWindowClampsToSheet`,
`TestHorizontalWindow`, `TestWindowReplacesAndKeepsSelection`,
`TestKeyboardJumpToOffscreenCell`,
`TestEditorTextSurvivesViewportReplacement`,
`TestFrozenChromeAndScrollingCells`, `TestTileCacheBounds`,
`TestStaleFetchDiscarded`, `TestQueuedFetchCoalesced`,
`TestFetchMissesWindowIsPaintOnly`. Rendering fetches bounded ranges
(cap 50,000 cells) from the in-memory workbook that `LoadWorkbook` built
from SQLite; `storage.LoadRange` and `TestRangeQueryPlan` cover the SQL
path. Keyset pages: `TestListWorkbooksKeyset`, `TestCSVPage`; the import
dialog shows the first 8 rows of the 50-row page.

## Phase 5: Windows and Ebiten integration

- [x] Launch `ownframe.Run` from `main` and keep page data, retained display operations, and window-related state on the UI loop. Workers send immutable results through bounded channels; handlers and ticks never wait on SQL, disk, HTTP, or sampling.
- [x] Apply worker results through `Page.SetTick` with a bounded drain budget. Use request generations to discard results from obsolete filters, pages, or closed screens. Preserve important commands; only replace disposable snapshots or progress updates.
- [x] Reacquire retained operation pointers after redraw or resize. Handle nil display lists with a correct bitmap path; distinguish paint-only updates from geometry changes. Measure the full-repaint cost of an active tick callback.
- [ ] Support small windows through an explicit minimum size or usable scrolling. Test resize while work completes, 100/150/200 percent Windows scaling, monitor moves, minimize/restore, focus loss, keyboard navigation, and wheel input. (pending: native Windows host; the minimum size is 480x320, and keyboard navigation, wheel scroll windows, and the bitmap fallback path are covered headless)
- [x] Stop timers and workers on window close, cancel requests, and join workers within a documented shutdown budget. Test close during migration, write, and active work without sending into closed channels or accessing a closed database.
- [ ] Run an actual Windows desktop session and record OS, architecture, display scaling, and renderer. Cross-compilation and WSL execution are separate evidence and do not substitute for native testing. (pending: no Windows host available)
- [ ] Keep Windows-specific file and process handling in platform adapters. Test paths under a user profile, locked files, and access denied without requiring administrator privileges. (pending: native Windows host; the Linux path, locked-database, and read-only tests pass)
- [x] Treat `-web`, if added, as a separate preview path: `Serve` does not tick and PNG output does not reflect retained operation edits. Provide an explicit update mechanism or label it as a still preview. (no `-web` mode was added; the desktop window is the only mode)

References: [ownframe frames](../../../documentation/frames.md), [features](../../../documentation/features.md), and [Ebitengine lifecycle](https://ebitengine.org/en/documents/cheatsheet.html).

Evidence 2026-10-07: `main.go` runs `ownframe.Run` with `-data`,
`-workbook`, and `-perf`. The worker is one goroutine over a bounded
queue; fetches coalesce and stale generations are dropped
(`TestTickBudget`, `TestStaleFetchDiscarded`). The screen does not hold
retained operation pointers, so there is nothing to reacquire after a
redraw; paint-only versus geometry is `applyFetch` returning false when a
range misses the rendered window (`TestFetchMissesWindowIsPaintOnly`).
`BenchmarkTickApply` measures a full-window fetch applied through the
tick plus redraw at 7.9-8.1 ms. Shutdown: `TestCloseStopsWorker`,
`TestCloseDuringPendingWork`, `TestCloseBudgetWhileWorkerBlocked`, with a
2 s join budget.

## Phase 6: Verification and completion

- [x] Write focused tests for domain behavior, migration/restart recovery, persistence failures, stale asynchronous results, and shutdown. Run concurrency tests with fake collectors, files, or HTTP services; keep clipboard tests on `clipboard.UseMemory`.
- [x] Run the finished example from a fresh data directory using dummy data, then with its real source. Exercise the full user workflow, restart, and verify saved state in the actual database.
- [ ] Profile representative and stress workloads. Record CPU, heap, RSS, allocations, goroutines, queue depth, dropped/coalesced updates, and frame/input timings. After warmup, a 30-minute mixed-workload soak must show bounded history and no sustained memory or goroutine growth. (pending: 30-minute soak and real profiling)
- [ ] Target p99 UI frame work below 16.67 ms at 60 Hz and input feedback within 100 ms on a named reference machine. These are acceptance targets, not measured claims. Report failures and renderer limitations explicitly. (pending: headless redraw numbers are 4.2-8.1 ms on an i7-13700HX, but no desktop GPU p99 run)
- [x] Run `gofmt` on changed Go files, keep each Go file at most 2000 characters, run `make test` and `make build` from the repository root, and run targeted race tests where supported. Store generated evidence under ignored `temp/`.
- [x] Add usage, dummy/real mode, storage location, recovery, platform limits, and measured performance to the example README and examples index. (2026-10-07: `examples/spreadsheet/README.md` is written; the index row landed in `examples/readme.md`)
- [ ] Mark complete only when the example works end to end, persistence survives restart, required Windows scenarios pass, and all earlier phase evidence is recorded. Record unavailable checks as pending. (pending: native Windows host and the 30-minute soak)

## Evidence log

For each phase, add date, revision, command or manual workflow, dataset size, platform/renderer, result, and log or screenshot path here. Leave unsupported or unrun scenarios unchecked.

- 2026-10-07, 2b6835a + e676d6c, `cd examples && go vet -p 2 ./spreadsheet/... && go test -p 2 ./spreadsheet/...`, dummy 200x20 and stress 100,000x20 seeds, Linux amd64 (Intel i7-13700HX), all packages pass. `go test -race -p 2 ./spreadsheet/ui/...` passes. Details in `temp/spreadsheet-ui-evidence.md`.
- 2026-10-07, 2b6835a + e676d6c, `make build` and `make test` from the repository root, both modules, Linux amd64, exit 0.
- 2026-10-07, 2b6835a + e676d6c, `go test -bench 'BenchmarkRedraw|BenchmarkTickApply' -benchtime 300x -count 2 ./spreadsheet/ui/`, 1200x800 window, Linux amd64, headless: scroll step 5.9-6.0 ms, window replace 4.2 ms, dense window 6.1-6.3 ms, tick apply 7.9-8.1 ms. Earlier per-cell border layout measured 17.8 ms.
- 2026-10-07, e676d6c, `TestUIShowsSeededWorkbook`, `TestUITypeCommitUndoOverStore`, `TestUICSVImportThroughPicker`, fresh temp dir with the real SQLite store, Linux amd64, seeded data renders, an edit and its undo persist across restart, and a CSV import commits atomically after preview.
