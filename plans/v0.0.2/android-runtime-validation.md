# Android runtime validation

## Scope and implementation

Work is on `feature/android-runtime-blinkless-34d9386`. Both modules require
`v0.0.0-20261010174808-34d93862f9ed`, resolving to the requested Blinkless
commit `34d93862f9ed2eec613b0cdea11d3d6da089b574`. Old engine checksums are removed.
The workspace configuration is unchanged, and the examples keep their local
OwnFrame replacement. No renderer, browser engine, JavaScript engine, or new
module dependency was introduced.

The audit reused existing cached parsing and stylesheets, display lists, dirty
buffers, history/routes, form bindings, text selection, IME composition, long
press, flick scrolling, pinch zoom, EbitenView lifecycle, and AndroidHost.
The engine API comparison identified corrected PointsPerPixel and PixelPerPoint
units; callers now use the factor matching their coordinate conversion.

Implemented corrections:

- Fixed operations use viewport culling. Scroll-buffer composition respects
  engine paint order, and GPU colors have valid premultiplied opacity.
- Bitmap fallback paints rounded geometry, borders, grids, affine images and
  text, opacity, and CSS blend groups. Invalid payloads and unsupported shaping
  options report errors. Canvas and blend-buffer sizes are bounded.
- A cached fallback viewport preserves fixed positioning during committed-size
  scrolling and reuses the GPU image on idle frames.
- One touch tap delivers one press. Pinch and native cancellation release active
  gestures. A shared native adapter clears Android touches because the current
  Ebiten input adapter ignores ACTION_CANCEL. Pause invokes cancellation.
- IME rectangles account for zoom and resize. Late callbacks cannot edit another
  focused field. Telegram queues keyboard blur and preserves composer drafts.
- Telegram reserves lateral system insets in its FrameLayout. GPU image eviction
  disposes resources, and font source/face caches have size limits.

## Automated and build checks

Checks run on Linux amd64 with Go 1.26.4 and an Xvfb display where needed.
Logs and profiles are local ignored artifacts under `temp/android-runtime/`.

| Check | Result and boundary |
| --- | --- |
| Both modules: `xvfb-run -a make test` | Passed on the final tree, with package concurrency 1. Runs root and examples packages. |
| Compilation/vet: `xvfb-run -a make build` | Passed on the final tree. Compiles root and examples without linking every executable. |
| GPU replay: `OWNFRAME_REPLAY_GPU_TEST=1 xvfb-run -a go test ./internal/replay -run TestGPUVisibleParity -count=1` | Passed, including fixed-layer culling and translucent fill regression. Desktop GPU path under Xvfb. |
| Runtime wasm: `GOOS=js GOARCH=wasm go vet -p 1 ./... ./examples/login ./examples/platform` | Passed on the final tree. |
| All examples wasm vet | Blocked by existing SQLite-backed examples: modernc/sqlite lacks the required js build support. Runtime and supported examples are checked separately. |
| `sh scripts/android.sh` | Passed on the final tree for all default Android ABIs. Builds Telegram AAR and debug APK. |
| Other Android hosts | Login, Platform, Dino, and LG Remote arm64 AARs and debug APKs passed. LG Remote passed its serial rerun. |
| Android installation and device runtime | Not run: `adb devices` lists no device or emulator. |

Regression coverage includes real engine unit conversion; bitmap paint, blending,
unknown operations, transformed text budgets, and resource limits; fallback viewport scrolling and GPU
reuse; touch press counts and cancellation; IME bounds and focus targeting;
font-cache bounds; PNG rejection of unsupported retained paint; and Telegram draft preservation on blur.
`TestAndroidLayoutsResizeAndModalInput` covers a responsive Flexbox/Grid dashboard,
long settings content, modal hit ordering, and form binding at 360×640, 640×360,
412×915, then 360×640 again. These tests exercise logical CSS sizes, not native
device density or the Android keyboard.

## Measurements

Measurements use Linux amd64, Intel Core i7-13700HX, 24 logical Go processors,
Xvfb, and the current tree. Other builds were running concurrently. They are
single-environment observations, with no controlled before/after comparison.
They establish neither Android frame pacing nor a performance improvement.

The existing page benchmarks ran three times with a one-second target and CPU
and allocation profiles:

| Benchmark | Time range | Bytes/op | Allocations/op |
| --- | ---: | ---: | ---: |
| RedrawWarm | 0.337–0.429 ms | 328,442 | 126 |
| RedrawHeavyWarm | 0.586–0.868 ms | 369,050–369,063 | 106 |

These are complete warm Page.Redraw operations, including template execution,
engine relayout, and display-list preparation. They do not isolate layout time,
GPU presentation, or input-to-render latency. The allocation profile attributes
83.4% of allocation volume to Blinkless styleStore.append. Its 6.15 GiB aggregate
is allocation volume across benchmark iterations, not live RSS.

The uninstrumented Platform desktop example, after five seconds of warmup,
used 41.97% of one CPU core over a 30-second idle interval. RSS went from
157.88 MiB to 162.56 MiB. A separate 35-second idle CPU profile sampled 36.84%
of one core; 73.59% of samples were in runtime.cgocall and 82.57% cumulatively
in Ebiten OSThread.loop. This points to native graphics/engine calls but does
not identify a specific native function. The runtime still ticks and presents
at Ebiten's normal cadence, even when it reuses the HTML layout. Low idle CPU
and long-run memory stability have not been demonstrated.

Reproduce the benchmark profile from the repository root:

```sh
OWNFRAME_PROFILE_DIR="$PWD/temp/android-runtime"
xvfb-run -a go test ./internal/page -run '^$' \
  -bench 'BenchmarkRedraw(Warm|HeavyWarm)$' -benchmem -benchtime=1s -count=3 \
  -cpuprofile "$OWNFRAME_PROFILE_DIR/cpu.pprof" \
  -memprofile "$OWNFRAME_PROFILE_DIR/memory.pprof" \
  -o "$OWNFRAME_PROFILE_DIR/page.test"
go tool pprof -top temp/android-runtime/cpu.pprof
go tool pprof -top -alloc_space temp/android-runtime/memory.pprof
```

## Pending device checks and limits

All seven requested Android scenarios remain pending native runtime validation:
login keyboard/focus, dashboard Flexbox/Grid, long settings scrolling, modal
stacking/input, portrait/landscape, system bars plus keyboard, and repeated
navigation plus background/resume. Test representative densities as well as
logical sizes, and exercise replay and bitmap fallback separately.

Input-to-render latency, scrolling frame pacing, Android CPU/RSS/GPU usage,
orientation performance, and navigation soak measurements need a device or
emulator. The desktop RSS interval is too short to establish memory stability.

Bitmap typography uses OpenType font drawing and has no full shaping contract.
FontFeatures and TextAutospaceGap currently fail with a rendering error. Bitmap
scrolling repaints a CPU viewport; full fallback pictures remain resident too.
Fixed projection applies at a committed normal viewport; fitted views and pending
resizes scale the full picture. Custom SetViewportPinZ counter-scroll behavior
requires replay. No persistent Activity/process restoration or accessibility
tree was added. Blinkless CSS semantics were not changed.

See [runtime behavior](../../documentation/android-runtime.md) and the
[build guide](../../documentation/android-build.md) for ownership and commands.
