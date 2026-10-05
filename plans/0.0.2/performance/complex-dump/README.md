# Complex layout dump

This command executes a dashboard template, prepares CSS, and creates the page's retained display list. It opens no window and writes no raster image. The fixture uses the same `internal/page.Page` implementation as the public API.

Run from the repository root:

```sh
sh plans/0.0.2/performance/complex-dump/run.sh
```

The script compiles one named command into a temporary file outside the checkout, runs five scenarios in three separate processes each, captures profiles, and runs geometry checks. Measurements run sequentially. Go files obey the repository's 2000-character limit.

For a single scenario:

```sh
go run ./plans/0.0.2/performance/complex-dump -mode cached -out /tmp/gpui-complex-sample
```

Modes are `initial`, `cached`, `data`, `resize`, and `windowed`. The initial mode creates a fresh page for each redraw; the other modes warm one page first. Data changes only the revision label through `SetData`. Resize alternates 1440 and 1520 pixels. Windowed retains 48 of 480 fixed-height rows with a bottom spacer. It demonstrates the top window's cost; it is not an interactive scrolling application.

Each process records seven timed redraws, then at least two seconds of repeated work for profiling. Times include `Redraw`, its opt-in sampling, and `TakeDirty`. Fresh-page construction and scenario preparation happen outside the timed sample; the CPU profile includes them. Initial sample zero carries process-first font work and is reported separately. Later initial samples are fresh pages with warm process caches.

CPU profiling runs during timing samples. These are diagnostic measurements, with profiler overhead, rather than release benchmarks. Heap profiling uses Go's default sampling rate. `alloc_space` is cumulative sampled traffic across the process; `inuse_space` is sampled live heap after GC. JSON `AllocBytes` and `Allocs` are allocator deltas for one timed redraw. None of these is RSS.

`*-layout.json.gz` contains the final boxes and a projection of display operations with kind, geometry, and text. It deliberately excludes font objects, image data, transforms, and other replay payloads. These files support structural analysis and cannot reconstruct a complete rendered frame. They describe the page after the profile loop, while `*.json` keeps individual timing samples.

The geometry analysis checks page dimensions, the grid height, all 480 fixed row heights, the first 48 row boxes, and ordered operation geometry/text intersecting the initial viewport. Operation coordinates are points at 0.75 per CSS pixel. Pixel parity, font payload identity, interactive scrolling, and offscreen paint equivalence require later checks.

Generated files are stored in `temp/complex-dump`, which is ignored by Git. Artifacts contain raw dumps, CPU and heap profiles, readable pprof tables, machine and source metadata, summary JSON, and repository validation logs. The report and live checklist are one directory above this command.
