## Summary

Add a shared Android host and per-app profiles for the mobile examples. The branch also ships release APKs, routes Android Back through each app, and improves Telegram touch scrolling while keeping its header and composer pinned.

---

## Motivation / context

The Android examples had separate activity and input handling, which made shared behavior such as insets, Back, lifecycle callbacks, and touch policy inconsistent. This change moves common host behavior into reusable code while leaving app-specific settings and actions with each example.

The work follows [the v0.0.2 Android shared-host checklist](../v0.0.2/android-shared-host-checklist.md). Samsung sizing, TalkBack, and several lifecycle and gesture cases remain device-verification items in that checklist.

---

## Changes

- Add shared Android host support and app profiles, including common viewport, inset, font-scale, Back, lifecycle, and input-policy callbacks.
- Apply Android profiles to LG Remote, Dino, and Telegram. Add Android pinch policy, shared host integration for Dino and Telegram, and release APKs for the mobile examples.
- Add reusable hold-repeat scheduling and preserve LG Remote's rapid button taps. Improve Telegram touch fling scrolling and keep its top bar and composer pinned during message scrolling.
- Document the shared host and release process, and track remaining work in the v0.0.2 checklist.

---

## Impact

| Area | Impact |
|------|--------|
| Android behavior | Common host callbacks and per-app profiles provide a shared place for platform behavior while preserving app-specific policies. Device coverage remains incomplete. |
| Input and scrolling | Android pinch can be controlled per app; hold-repeat scheduling is reusable. Telegram retains its pinned controls while touch scrolling uses the shared window path. |
| Build and distribution | Release APKs are included for the mobile examples. Build intermediates remain excluded. |
| Compatibility | Existing desktop and browser behavior continues through the existing window paths; Android profile overrides control app-specific behavior. |

---

## Breaking changes / migration

| Item | Migration |
|------|-----------|
| None known | Existing non-Android callers keep their current behavior. Android examples use explicit host profiles. |

---

## Test plan

- [x] `make test TEST_P=4` passed on the feature branch.
- [x] Android release APKs were built and included for the mobile examples.
- [x] Telegram Back behavior and touch scrolling were exercised on Pixel 7 during development.
- [ ] Verify sizing, density, cutouts, keyboard behavior, and touch target alignment on the reported Samsung device.
- [ ] Verify TalkBack traversal, Bluetooth reconnect behavior, activity recreation, and app-specific gesture handling on physical Android devices.

No claim is made that the remaining physical-device checks passed. See the checklist for per-area status and evidence boundaries.

---

## Screenshots / sample output

No screenshots are attached. The release APKs are included in the branch.

---

## Related issues

No related issue was found for this work.

---

## PR metadata checklist (author)

- [x] Self-assigned with `--assignee "@me"`.
- [x] Labeled `enhancement` and `documentation`.
- [x] Related issues recorded as none; no placeholder issue IDs.
- [x] Filled body saved under `plans/PR/pr-android-shared-host.md`.

---

## Follow-ups (out of scope)

- Complete the remaining Samsung and Pixel device checks listed in the shared-host checklist.
- Finish broader Android host migration and verify TalkBack, keyboard/insets, Bluetooth reconnects, and lifecycle recreation.
- Continue the wider engine and window race audit tracked outside this branch's app-level changes.

---

## Reviewer checklist

- [ ] Confirm the shared host owns only Android behavior and app-specific policies remain overridable.
- [ ] Confirm the three app profiles preserve each app's required interaction model.
- [ ] Confirm release APKs are the intended deliverables and no build intermediates are included.
- [ ] Confirm the stated tests and device-test limitations match the evidence.
- [x] PR creation specifies an assignee and labels.
- [x] Related issues contain no fabricated IDs.
- [x] Diff-stat-by-extension table is generated from `scripts/pr-diff-stat.sh` and included below.

---

## Diff stat by extension

| Extension | Files | Insertions | Deletions |
| --- | ---: | ---: | ---: |
| `.apk` | 5 | Binary | Binary |
| `.bat` | 2 | 184 | 0 |
| `.go` | 45 | 856 | 127 |
| `.gradle` | 15 | 96 | 0 |
| `.jar` | 2 | Binary | Binary |
| `.java` | 13 | 524 | 158 |
| `.md` | 11 | 592 | 13 |
| `.properties` | 4 | 14 | 0 |
| `.sh` | 2 | 71 | 0 |
| `.xml` | 5 | 41 | 0 |
| No extension | 8 | 520 | 0 |
| **Total** | **112** | **2898** | **298** |
