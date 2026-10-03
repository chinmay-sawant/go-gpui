# Doom for go-gpui — phase-wise agent checklist

Plan status: ready, not started. This folder intentionally holds no Go code yet.
Execution note: this checklist is the artifact. The 12 agents below run only
after the plan is approved, one phase at a time, and no later phase starts
while its predecessor's gate is red.

## 1. Non-negotiables

### From scratch

- [ ] No WAD files, no Doom source, no Doom ports, no raycasting libraries,
      no game engines, no rendering libraries.
- [ ] Nothing new in `go.mod`. The only modules are the ones already there
      (`gowkhtmltopdf` via the replace, Ebiten, `golang.org/x/text`).
- [ ] Map format, levels, art, sprites, sounds, and game logic are written in
      this repo, by us.
- [ ] No downloaded assets. Pixels are CSS/HTML or generated in Go with the
      standard library. Sounds are synthesized, or the phase is skipped.

### HTML and CSS are the renderer

- [ ] The visible scene is one HTML document styled by CSS and painted
      through the library's display-list replay:
      `html.Parse` → `css.Apply` → `layout.DisplayList` → `internal/replay`.
- [ ] Per-frame motion mutates display ops inside `Page.SetTick`
      (`X`, `Y`, `W`, `H`, `R`, `G`, `B`, `Alpha`, `Text`).
      No per-frame `Redraw`; `Redraw` is 130–400 ms in the existing examples.
- [ ] No bitmap shortcut, no per-frame `SetImage` encodings, no bypass of the
      HTML/CSS path. If an element can be a styled op, it is a styled op.
- [ ] A `Redraw` replaces the display list, so a tick callback re-acquires its
      op handles after every redraw (Phase 1 builds the index for this).

### House rules (AGENTS.md)

- [ ] Every Go file is at most 2000 characters. Check with `wc -m`.
      Expect roughly 50 lines per file, so split by concern.
- [ ] `gofmt` on every edited Go file.
- [ ] `make test` is green at every gate. Tests ship with the feature,
      in the style of `examples/dino`.
- [ ] v1 is keyboard only: `Handlers.KeyDown` / `KeyUp`, lowercase names,
      one press and one release per real event. We track held keys ourselves.
- [ ] `Run` (desktop and wasm) is the target. `Serve` never ticks; document
      that a `-web` page shows the last `Redraw`.

## 2. Rendering budget

What the tick may touch each frame, and the ceiling for Phase 0 to prove:

| Element | Ops | Fields changed per frame |
| --- | --- | --- |
| View columns (walls) | `COLS` (start 320) | height, Y, R, G, B |
| Floor and ceiling | 2 × bands (4–8 per side) | R, G, B |
| Sprite slots | slots × frames, alpha-selected | X, Y, W, H, Alpha |
| HUD text | ~6 | `Text`, alpha |
| HUD bars, damage flash | ~4 | W, alpha |

- [ ] Total mutated ops stay under ~450 per frame.
- [ ] Tick target: under 8 ms median on this machine, measured in Phase 0
      and logged in the decision log.

## 3. The 12 agents

One writer per file. Later agents own new files only; A1 reviews integration.

| Agent | Role | Owns | Starts after |
| --- | --- | --- | --- |
| A1 | Lead, page skeleton, integration | `new.go`, `page.go`, `tick.go` | plan approved |
| A2 | Scene grid, op index | `html_*.go`, `css.go`, `ops.go`, `scene.go` | Phase 0 |
| A3 | Map format and levels | `map.go`, `levels.go` | Phase 1 |
| A4 | Movement, collision, doors | `collide.go`, `controls.go`, `door.go` | Phase 1 |
| A5 | Raycast view | `raycast.go`, `columns.go` | Phase 2 |
| A6 | Shading, floor, ceiling | `shade.go`, `planes.go`, `palette.go` | Phase 2 |
| A7 | Sprites and depth | `sprite.go`, `depth.go` | Phase 3 |
| A8 | Items, frames, animation | `items.go`, `anim.go` | Phase 3 |
| A9 | Enemies and AI | `enemy.go`, `ai.go`, `los.go` | Phase 4 |
| A10 | Weapons and damage | `weapon.go`, `projectile.go`, `damage.go` | Phase 4 |
| A11 | HUD, menus, game states | `hud.go`, `menu.go`, `state.go` | Phase 5 |
| A12 | Sound, polish, docs | `sound.go`, `README.md`, docs edits | Phase 6 |

## 4. Phases

### Phase 0 — Feasibility spike (A1, A2)

Goal: prove the HTML/CSS op grid can carry a game before writing a game.

- [ ] Create `examples/doom/doom/` with a package comment file.
- [ ] Build a page with `COLS = 320` column divs and a tick that recolors
      every column from a sine.
- [ ] Build a second variant with columns split in two, to learn headroom.
- [ ] Write a test that counts ops after `Redraw` and proves one tick changed
      them without a redraw.
- [ ] Measure median and worst tick time over 300 frames; log both.
- [ ] Decide the sprite technique and the floor technique, with numbers:
      alpha-toggled image ops vs stacked fill ops.
- [ ] No game logic yet. Code here becomes the seed of the real package.

Gate: window runs, numbers are in the decision log, tick stays under budget,
technique decisions recorded. Red gate means shrink `COLS` before Phase 1.

### Phase 1 — Page skeleton and op index (A1, A2)

- [ ] `New()` in the style of `examples/dino/dino/new.go`: `gpui.New`,
      `Handle`, `SetTick`, and the example `main.go`.
- [ ] Template builder split across `html_*.go` and `css.go`, each file
      under the cap, following the dino file shape.
- [ ] Stable slot map: one op per view column, band, sprite slot, HUD field.
- [ ] Index rebuild: after any `Redraw`, re-find each op and its slot, then
      fail loudly if a slot is missing.
- [ ] Held-key map for `KeyDown`/`KeyUp` with tests.
- [ ] `make test` target coverage for every new file.

Gate: `go run ./examples/doom` opens, grid and HUD placeholder visible,
index test and key tests pass.

### Phase 2 — World and movement (A3, A4)

- [ ] Map format from scratch: ASCII rows plus a legend (wall types, doors,
      spawn, exit, items). No binary formats.
- [ ] Three hand-authored levels: corridor, arena, maze.
- [ ] Collision against the grid with a player radius; slide, do not stick.
- [ ] Movement: forward, back, strafe, turn, run; dt capped like dino's.
- [ ] Doors: closed / opening / open / closing, player-triggered.
- [ ] Camera constants in one place (FOV, horizon, movement speeds).

Gate: tests prove no clipping, sliding works, doors cycle and block while
closed, unknown map symbols fail at load.

### Phase 3 — Raycast view (A5, A6)

- [ ] One DDA ray per column over the map grid, fisheye correction.
- [ ] Column height from distance, Y from the horizon, wall-type color.
- [ ] Distance shading into R/G/B; no gradient bands that need new ops.
- [ ] Floor and ceiling bands colored from the same palette module.
- [ ] Golden tests: fixed player poses map to exact expected column heights
      and colors, computed by hand in the test.
- [ ] Wall types: stone, metal, door frames, exit switch, all as CSS colors.

Gate: golden tests pass; view matches the map; tick budget still holds with
movement and doors running.

### Phase 4 — Sprites and depth (A7, A8)

- [ ] Use the Phase 0 decision. Preferred: pre-placed image ops generated in
      Go with `image/png`, frame selection by `Alpha`, motion by `X/Y/W/H`.
- [ ] Depth: reuse the column pass to build a per-column depth buffer.
- [ ] Occlusion: split a sprite into vertical strips whose alpha follows the
      depth buffer, so a wall hides the near half.
- [ ] Items: health, ammo, armor, keys; pickup radius, respawn rules.
- [ ] Idle animations for items and decorations.
- [ ] Sprite frames are ours: drawn in Go or composed from fill ops.

Gate: test proves a sprite behind a wall is hidden, in front is visible;
pickups change the values the HUD will read.

### Phase 5 — Combat and AI (A9, A10)

- [ ] Enemy states: idle, chase, attack, pain, death, corpse.
- [ ] Line of sight through the same DDA code; no shooting through walls.
- [ ] Three enemy kinds, one ranged, one melee, one fast; state tests.
- [ ] Weapons: pistol, shotgun, chaingun; cooldown, ammo, spread.
- [ ] Hitscan and one projectile type with wall collision.
- [ ] Damage both ways, armor absorbed, death, level restart on death.
- [ ] Balance constants live in one file, not sprinkled through ai.go.

Gate: scripted fight test (enemy dies after N hits), no-damage-through-wall
test, death and restart test.

### Phase 6 — HUD, menus, game states (A11, with A1 and A2)

- [ ] State machine: title, playing, paused, dead, level complete, finale.
- [ ] Menus on the `Redraw` path only: new game, controls, pause, quit.
- [ ] HUD: face built from fill ops, health / ammo / armor / keys text.
- [ ] Damage flash by mutating one full-screen fill's alpha.
- [ ] Level transitions, exit switch objective, three-level finale.
- [ ] Headless tests for every transition.

Gate: title → play → win and title → play → death both reachable in tests;
manual run confirms the same.

### Phase 7 — Sound (A12)

- [ ] Synthesize short effects with the pattern in `examples/music`:
      shoot, hurt, enemy death, door, pickup, menu blip.
- [ ] Event-to-sound mapping, one mute toggle, no audio in tests.
- [ ] No downloaded files; no new modules.

Gate: with sound off the game behaves identically; `go run` plays effects.
This phase may be skipped if it puts the schedule at risk, and the skip is
recorded.

### Phase 8 — Integration, QA, docs (all; A1 leads)

- [ ] `make test` green from the repo root.
- [ ] `wc -m` audit: every Go file at or under 2000 characters.
- [ ] `gofmt -l` clean.
- [ ] Perf log updated: median and worst tick, worst frame observed.
- [ ] `examples/doom/README.md`: controls, rules, how to run, what is from
      scratch, why the HTML/CSS path is used.
- [ ] Update the example lists: root `README.md`, `AGENTS.md`,
      `documentation/features-examples.md`.
- [ ] Manual play-through: every weapon, every door, every key, death, win.
- [ ] `GOOS=js GOARCH=wasm go build` check; note the `Serve` no-tick limit.
- [ ] Dead code and unused ops removed.

Gate: the definition of done below is fully checked.

### Phase 9 — Stretch: mouse-look (separate approval, not v1)

This is the only planned change to library code, so it does not start by
default.

- [ ] `internal/window/pointer.go`: cursor capture via Ebiten, motion deltas.
- [ ] `Handlers` gains a move callback; `internal/page` plumbs it through.
- [ ] Docs `documentation/keys.md` (or a new `pointer.md`) and AGENTS.md.
- [ ] Tests with a fake screen, like the existing input tests.

Gate: not required for v1. Record the decision.

## 5. Definition of done

- [ ] `go run ./examples/doom` plays a full loop: title, three levels,
      combat, keys, exit, win, death, restart.
- [ ] Every visual element is an HTML/CSS display op; per-frame motion is a
      tick mutation; `Redraw` happens only on real content changes.
- [ ] No third-party rendering or game code; `go.mod` unchanged.
- [ ] All Go files pass the 2000-character cap and `gofmt`.
- [ ] `make test` green; each system above has tests.
- [ ] Docs updated in this folder and in the repo lists.

## 6. Parallel-work rules

- [ ] Phases run in order. Within a phase, agents with disjoint files run in
      parallel; A1 resolves overlaps.
- [ ] One writer per file at any time. Agents add new files instead of
      editing another agent's file, then A1 consolidates.
- [ ] Every agent leaves the repo in a compiling state and `make test` green.
- [ ] Every contested decision gets one row in the log below before the code
      that depends on it.

## 7. Decision log (append-only)

| Date | Decision | Why | Evidence |
| --- | --- | --- | --- |
| — | Plan uses 12 agents, 10 phases, HTML/CSS rendering only | required by this checklist | this file |
| — | `COLS`, sprite technique, floor technique | Phase 0 | tick measurements |
| — | Sound phase kept or skipped | Phase 7 | schedule check |
