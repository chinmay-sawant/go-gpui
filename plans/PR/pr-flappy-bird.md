## Summary

Add a playable Flappy Bird example, and move the game examples' HTML and CSS out of Go strings into `template/` files. The bird, pipes, clouds, and ground are plain HTML boxes styled by CSS; `Page.SetTick` runs the physics, spawns the pipes, scores, and moves the retained display-list operations each frame, so the game animates without a `Redraw`. The branch also adds a `make build` compile check that links nothing.

## Motivation / context

- Requested example: a Flappy Bird built with HTML and CSS directly, alongside `examples/dino`.
- The dino and flappy-bird examples were the only two that kept HTML/CSS in Go; every other example already embeds a separate `.html`. Both games now follow the same `template/` convention.
- No issue is linked; the repo has no tracking tickets for this work.

## Changes

### Example: `examples/flappy-bird`

- `go run ./examples/flappy-bird` opens the desktop window; `-web` serves the static picture on `127.0.0.1:8121`.
- The simulation applies gravity and flap velocity, spawns and culls pipe pairs, shrinks the gap and grows the speed with distance, ends the run on ground or pipe contact, and tracks score and best (`motion.go`, `control.go`, `spawn.go`, `collide.go`, `pose.go`).
- Space, up, W, Enter, or a click flaps; R or Space restarts after a crash and keeps the best. The first press flaps and starts the run.
- The scene is one 480x720 HTML/CSS page; `bind.go` and `paint_*.go` move the cached display-list fills and text each tick.
- Hidden fills move off the canvas as well as collapsing to zero size, because a zero-sized rounded rect still painted its corner.
- 27 tests cover the scene ids, physics, collisions, spawning, painting, and the HUD.

### Templates: dino and flappy-bird

- `template/index.html` and `template/styles.css` hold each game's scene and styles. `html.go` embeds both with `//go:embed` and inlines the stylesheet at the `<!-- stylesheet -->` marker, because the engine reads one HTML string and does not fetch a linked stylesheet.
- Dino's overlay message constants move to `text.go`.
- A `template_test.go` in each package guards the marker.
- The assembled HTML is byte-identical to the old Go-built output, and the CSS files are byte-identical to the old `styles` strings.

### Build guidance

- `Makefile` gains a `build` target that runs `go vet -p 1 ./...`, a compile check for every package that links nothing and writes no binaries. Raise the limit with `make build BUILD_P=4`.
- `AGENTS.md` tells agents to run `make build` instead of `go build ./...`, which links an executable for each example, takes about a minute, and leaves the binaries in the repo root.

### Docs

- `examples/readme.md`: flappy-bird row with port 8121; the game-keys note now says flappy-bird also takes clicks.
- `documentation/frames.md`: the tick examples paragraph lists dino and flappy-bird.
- `documentation/features-examples.md`: flappy-bird entry and the test-file count refreshed.

## Impact

| Area | Effect |
|---|---|
| Performance | The tick moves display operations; a click redraws the page once, as the click path always does. |
| Memory | No new long-lived allocations; moving pipes reuses the slice in place. |
| Behavior/correctness | New example only; dino is unchanged apart from the template refactor, proven byte-identical. |
| API/CLI | None; `examples/` only. |
| Dependencies | None. |
| Binary size/build time | Not measured. |

## Breaking changes / migration

| Item | Migration |
|---|---|
| None | - |

## Test plan

- [x] `make test`
- [x] `make build`
- [x] `go vet ./examples/dino/... ./examples/flappy-bird/...`
- [x] `gofmt -l examples/dino/ examples/flappy-bird/` prints nothing
- [x] Desktop window: ready screen, gameplay with pipes and score, and game-over board draw through the replay path
- [x] `buildHTML()` byte-identical before and after the template refactor, checked against a `HEAD` worktree
- [ ] No CI is configured in this repo.

### Commands

```sh
make test
make build
go run ./examples/flappy-bird
```

## Screenshots / sample output

Desktop captures taken while building the example (not committed; scratch under `/tmp/opencode`): the ready screen with the bobbing bird, a mid-run frame with a pipe pair and the score, and the game-over board. The hidden-board sparkle seen in the first captures is fixed and gone in the final ones.

## Related issues

None. No issues exist in this repo; this PR is the reference for the work.

## PR metadata checklist (author)

- [x] Self-assigned with `--assignee "@me"`.
- [x] Labels applied (`enhancement`, `documentation`).
- [x] Body copy committed at `plans/PR/pr-flappy-bird.md`.
- [ ] Ticket IDs linked: not applicable, no issues exist.

## Follow-ups (out of scope)

- The other examples keep their CSS inline in their `.html`; split those into separate stylesheets if the convention should apply repo-wide.

## Reviewer checklist

- [ ] Behavior matches summary and test plan
- [ ] No unrelated changes in the diff
- [ ] Public API / CLI changes documented (none)
- [ ] Related issues use correct Closes/Relates keywords
- [ ] No secrets or generated artifacts committed
- [ ] Diff-stat-by-extension table pasted at the bottom

## Diff stat by extension

| Extension | Files | Insertions | Deletions |
| --- | ---: | ---: | ---: |
| `.css` | 2 | 22 | 5 |
| `.go` | 36 | 1477 | 94 |
| `.html` | 2 | 132 | 0 |
| `.md` | 4 | 11 | 5 |
| No extension | 1 | 10 | 0 |
| **Total** | **45** | **1652** | **104** |
