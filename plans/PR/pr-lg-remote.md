## Summary

Add an LG webOS remote for desktop and Android, with Wi-Fi pairing, Bluetooth HID controls, Wake-on-LAN, app shortcuts, and a pointer pad. The phone layout fits each panel without scrolling and provides readable themes, accessible controls, and immediate press feedback. Rapid taps remain separate commands, and holding volume, channel, or direction buttons repeats with increasing speed.

---

## Motivation / context

- The remote needs to work on a Pixel 7 with system insets, enlarged text, keyboard input, and short taps.
- Device feedback identified a startup crash, layout movement, unreadable button text, delayed volume/channel commands, and missing press-and-hold behavior.
- Volume/channel requests previously blocked the command worker until each TV reply arrived. Android button events now queue both touch edges instead of depending on frame polling.
- Filled copy of `skills/PR/PR_TEMPLATE.md`, adapted to ownframe. No tracker issue exists for this work.

---

## Changes

### TV connection and controls

- Add SSAP pairing and saved client keys, SSDP discovery, model/MAC metadata, Wake-on-LAN, pointer control, and TV/app/input/media commands.
- Keep healthy pointer sockets connected, monitor connection health, and remove the extra volume-status request.
- Send volume and channel steps in order while collecting acknowledgements separately. Bound outstanding steps to 16 and report timeouts without replaying a possibly applied command.
- Use the pointer fallback after an explicit SSAP rejection. Allow commands when the local power indicator is stale.
- Update Power/Wake state immediately and ignore stale power results.
- Replace the Netflix shortcut with Disney+ Hotstar artwork and discover the installed Hotstar application ID.

### Phone interface and gestures

- Use named buttons, shared control metadata, responsive rows, and a reserved status height.
- Fit Remote, Pad, and Numbers panels into the viewport, with scrolling disabled. Keep fitted drawing, hit testing, and native accessibility bounds aligned.
- Improve light/dark contrast and highlight presses/focus through button backgrounds.
- Capture pad gestures, coalesce pointer movement, and add a 0.25x to 3.00x sensitivity slider with keyboard and native accessibility actions.
- Preserve quick Android down/up events in a bounded native-action queue and paint each burst once per tick.
- Start hold repeats after 350 ms, then repeat every 120 ms, 80 ms after two seconds, and 50 ms after four seconds.
- Stop new repeats on release, cancellation, sliding off, a second finger, activity pause, or a connection mode/IP change. Discard unsent repeats from an earlier hold.

### Android and Bluetooth

- Add an Android Gradle project and `scripts/android-lg.sh`, including Windows adb support under WSL2.
- Respect system bars, display cutouts, IME insets, system font scale, and Back navigation.
- Expose visible controls through native virtual accessibility nodes. Guard event dispatch when accessibility is disabled to fix the observed Pixel 7 startup crash.
- Dispatch Bluetooth commands through a native listener, bound HID work, and reconnect to the last TV across connection/service/lifecycle changes.

### Library support and documentation

- Add optional pointer-press and captured-drag callbacks, disabled-button handling, and `Page.PressedID()` for hold tracking.
- Honor locked views for fitting, scrolling, zoom, and coordinate conversion.
- Synchronize renderer adapter calls around the pinned engine's shared registration state and keep performance hooks owned by each window.
- Publish saved TV metadata atomically. Document the remote, Android build, hold behavior, asynchronous step replies, and interaction APIs.

---

## Impact

| Area | Impact |
|------|--------|
| Performance | Volume/channel sends no longer wait for the preceding acknowledgement. Pad movement is coalesced and native tap bursts share a redraw. Actual TV latency has no measured before/after result. |
| Memory | Wi-Fi jobs, native actions, Bluetooth work, and outstanding step requests are bounded. The remote retains connection metadata and accessibility snapshots. |
| Behavior / correctness | Short taps are retained, held incremental controls repeat, fitted controls keep matching hit bounds, and status messages do not move the layout. |
| API / CLI | Additive press/drag handlers and `Page.PressedID()`. Add the LG remote example, its web preview, and Android build/install script. |
| Dependencies | No Go dependency or engine-pin change. Android wrapper/build files are included; generated APK/AAR outputs remain local. |
| Binary size / build time | The final Android debug build succeeded. No comparative binary-size or build-time benchmark was run. |

---

## Breaking changes / migration

| Item | Migration |
|------|-----------|
| None | Existing callers keep their behavior unless they opt into the new handlers or locked-view controls. |

---

## Test plan

- [x] Earlier mobile/layout work passed `make test`, `make build`, focused LG race checks, and core page/window/renderer race checks before the rapid-tap/hold patch.
- [x] Earlier mobile matrix covered 108 viewport/theme/font/panel cases.
- [x] Earlier Pixel 7 checks reproduced and resolved the accessibility-disabled startup crash and inspected fitted controls, native bounds, scrolling, and the Hotstar icon.
- [x] Final rapid-tap/hold source compiled through `sh scripts/android-lg.sh`; Gradle `:app:assembleDebug` succeeded.
- [x] Final APK installed on the connected Pixel 7 with Windows adb `install -r`, returning `Success`.
- [x] Go edits formatted and checked against the 2000-character file limit; `git diff --check` passed.
- [ ] Go unit/race tests after the rapid-tap/hold patch. Skipped at the user's explicit request. Earlier test results are not presented as validation of that patch.
- [ ] Automated or measured TV-response validation of rapid taps, repeat acceleration, and release behavior. The installed app is available for device testing.

### Commands

```sh
# Earlier mobile/layout validation, before the rapid-tap/hold patch:
make test
make build

# Final rapid-tap/hold build:
sh scripts/android-lg.sh

# Final installation through Windows adb under WSL2:
adb.exe -P 5038 -s 2C161FDH200GBQ install -r <Windows path to app-debug.apk>

git diff --check
```

PDF reference metrics and the template's gowkhtmltopdf timing gates do not apply to this remote example. No tests were rerun for the final patch, as requested.

---

## Screenshots / sample output

```text
BUILD SUCCESSFUL in 13s
33 actionable tasks: 3 executed, 30 up-to-date
Performing Streamed Install
Success
```

Screenshots, logs, saved TV credentials, and generated APK/AAR files remain local under ignored paths.

---

## Related issues

- None. No tracker issue exists for this work.

---

## PR metadata checklist (author)

- [x] Self-assigned with `--assignee @me`.
- [x] Labels specified as `enhancement`, `bug`, and `accessibility`.
- [x] Related issues explicitly recorded as none; no placeholder issue IDs.
- [x] Filled body committed under `plans/PR/pr-lg-remote.md`.

---

## Follow-ups (out of scope)

- Complete TalkBack navigation and keyboard/cutout checks across Android versions.
- Exercise Bluetooth reconnects/permission denial and actual TV response to short taps and held controls.
- Add focused regression coverage for the final native-touch and step-pipeline changes when tests are requested.

---

## Reviewer checklist

- [ ] Behavior matches the summary and test plan, including repeat cancellation on release.
- [x] Unrelated generated dino/telegram `.class` files are excluded.
- [x] Public interaction changes are documented.
- [ ] Final native-touch/step-pipeline changes have regression coverage; explicitly deferred above.
- [x] PR creation specifies assignee and labels.
- [x] Related issues contain no fabricated IDs.
- [x] Saved client keys and generated build outputs are excluded. Source artwork and the Gradle wrapper are intentional assets.
- [x] Diff-stat-by-extension table generated with `scripts/pr-diff-stat.sh` is included below.

---

## Diff stat by extension

| Extension | Files | Insertions | Deletions |
| --- | ---: | ---: | ---: |
| `.bat` | 1 | 92 | 0 |
| `.css` | 1 | 101 | 0 |
| `.go` | 159 | 6267 | 206 |
| `.gradle` | 4 | 48 | 0 |
| `.html` | 2 | 137 | 0 |
| `.jar` | 1 | Binary | Binary |
| `.java` | 5 | 898 | 0 |
| `.md` | 5 | 228 | 0 |
| `.png` | 4 | Binary | Binary |
| `.properties` | 2 | 9 | 0 |
| `.sh` | 1 | 77 | 0 |
| `.svg` | 1 | 254 | 0 |
| `.xml` | 1 | 29 | 0 |
| No extension | 2 | 255 | 0 |
| **Total** | **189** | **8395** | **206** |
