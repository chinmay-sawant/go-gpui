# v0.0.2 - Shared Android host and application migration checklist

> **Parent:** [v0.0.2 plan](README.md)
> **Status:** Planned. Implementation and device acceptance remain open.
> **Recorded:** 2026-10-09, against `master` at `d679bd8`.
> **Estimated effort:** 7-12 focused engineering days, plus device verification. This is a planning estimate; Samsung diagnosis and Android packaging choices can change it.

---

## Overview

Move reusable Android behavior from the LG remote into ownframe so LG Remote, Dino, Telegram, and future applications receive the same host, sizing, input, accessibility, and lifecycle support. Expose explicit defaults and application overrides. Migrate the three existing Android applications to the shared implementation while preserving their respective controls and navigation.

This checklist also tracks the Samsung screen-adjustment report. The Pixel 7 result does not establish correctness on Samsung devices, other densities, display zoom settings, or Android versions. Establish the cause before changing density conversion: Ebiten already converts its Android view dimensions from physical pixels to dp.

This is a new planning document. Creating it does not implement the phases, modify existing version ledgers, rebuild applications, install APKs, or authorize a new commit/push. Earlier implementation and verification are recorded as baseline context, not completion of this migration.

## Executive summary

- Shared support already exists for resizing, captured gestures, pointer press callbacks, `Page.PressedID()`, fitted rendering, coordinate conversion, scroll gating, renderer synchronization, and per-window performance state.
- Android activity setup, native button delivery, accessibility nodes, font reporting, keyboard/Back handling, and lifecycle integration still differ between examples.
- Current `LockView` combines fitting with disabling scrolling and all page zoom. Separate those policies so an Android app can disable page pinch while preserving scrolling or desktop zoom.
- General page zoom remains available. LG Remote, Dino, and Telegram will explicitly disable Android page pinch through their application profiles. Apps can override that setting. An app-owned gesture, such as zooming an image, remains a separate choice.
- Preserve the LG remote's no-scroll panels, Dino's landscape/fullscreen game controls, and Telegram's scrolling conversation, pinned composer, and attachment workflows.
- Diagnose Samsung sizing and verify final displayed touch targets before claiming automatic Android screen adjustment is complete.

## Checklist conventions

- `[ ]` means not implemented, not verified, or awaiting required evidence.
- `[x]` requires matching implementation and verification evidence recorded in this file.
- `[~]` means a documented partial/deferred item with the reason and next required check.
- Stable item IDs identify work and evidence even when rows move within this document.
- All phase rows start open. Source inspection establishes the baseline; it does not complete a device test or migration.
- Record physical-device results separately from unit tests, Android compilation, emulator checks, and desktop/web previews.
- Store generated APKs, screenshots, traces, and logs under ignored `temp/android-shared-host/`. Do not include saved pairing keys or private chat/media content in evidence.

## Current source baseline

| Area | Current implementation | Remaining boundary |
|---|---|---|
| Resize | [internal/window/resize.go](../../internal/window/resize.go) records window size and relayouts. | Android density/settings changes, insets, and recreation require end-to-end verification. |
| Density | Ebiten v2.10.4's Android view converts view pixels to dp before reporting layout. | Compare its density with the Android window/resources metrics on the reported Samsung device. Do not apply a second conversion without evidence. |
| Fit and zoom | [internal/window/view_lock.go](../../internal/window/view_lock.go), [zoom.go](../../internal/window/zoom.go), and [page_config.go](../../internal/page/page_config.go) implement `LockView`. | Fitting, page pinch, desktop zoom, and scrolling are coupled. |
| Final touch sizes | [phone_fit.go](../../examples/lg-remote/remote/phone_fit.go) uniformly fits the remote; [mobile_fit_test.go](../../examples/lg-remote/remote/mobile_fit_test.go) checks button sizes before fitting. | A logical 48-pixel button can shrink below the desired final Android touch size. |
| Android insets | LG pads the native container; Telegram forwards dp insets for app layout. Dino has a minimal fullscreen activity. | Establish one explicit inset contract with host-managed and app-managed modes. |
| Fonts | LG's activity reports `round(16 * fontScale)`; [bridge/phone.go](../../examples/lg-remote/bridge/phone.go) clamps the result to 16-48. | Smaller system text settings are ignored; Android nonlinear font scaling is not modeled correctly. |
| Native touches | [ButtonTouch.java](../../examples/lg-remote/android/app/src/main/java/com/chinmaysawant/lgremote/ButtonTouch.java) and the LG bridge queue down/up/cancel events. | Delivery is example-specific, with a hardcoded pad exception. |
| Hold repeat | [remote/repeat.go](../../examples/lg-remote/remote/repeat.go) repeats volume/channel/direction controls. | Move the reusable timer/cancellation behavior; keep TV dispatch in the remote. |
| Accessibility | [RemoteAccessibility.java](../../examples/lg-remote/android/app/src/main/java/com/chinmaysawant/lgremote/RemoteAccessibility.java) exposes virtual controls. | It depends on LG metadata and skips bounds rebuilding when the snapshot is unchanged, even if native view dimensions changed. |
| Lifecycle | Each example owns its activity and mobile bridge. | Unify pause/resume, cancellation, listener ownership, and recreation behavior. |
| Back | LG dismisses the IME explicitly on API 30+ before panel navigation; Telegram has its own Back callback. | Verify Android 9-10 keyboard dismissal and share the dispatch order. |
| Bluetooth recreation | LG creates a new `Hid` controller with `wanted = false`; the bound Go application can retain Bluetooth mode. | Prove whether recreation leaves UI mode and native connection intent inconsistent, then restore intent explicitly. |

The LG branch was merged through [PR #17](https://github.com/chinmay-sawant/ownframe/pull/17). Earlier layout/core checks passed before the final rapid-tap/hold patch. That final patch compiled and was installed on Pixel 7; automated tests were skipped at the user's request. No Samsung reproduction or measurements were recorded. Preserve these limits when reporting progress.

## Common Android feature contract

These are proposed product defaults and override requirements. Phase 2 chooses the final API names and packaging; the table does not claim these options already exist.

| Common feature | Expected shared behavior/default | Application override |
|---|---|---|
| Viewport and density | Follow the current app window in logical units; convert native pixels exactly once. | Apps may select a logical game viewport or minimum size. Drawing, input, and accessibility must use the same resulting transform. |
| Resize/configuration | Refresh layout when window size, density, display zoom, font settings, or insets change. | Apps may choose responsive relayout or an aspect-preserving game viewport. They may observe changes through a callback. |
| Safe areas | Report all four system-bar and cutout insets in declared units. Host-managed padding is the ordinary UI default. | App-managed inset handling for pinned UI or fullscreen games; exactly one layer owns padding. |
| Keyboard | Report IME visibility and space; keep focused fields/composers reachable. | Resize, app-managed adjustment, or a documented no-keyboard game profile. |
| Font scale | Respect Android text scaling, including supported smaller and enlarged settings and nonlinear conversion. | Set a base text size, explicit app text policy, or a custom resolver. Keep user-visible controls readable. |
| Page pinch | Preserve general ownframe page-pinch availability. | Independent Android page-pinch enable/disable option; the three migrated apps explicitly disable it. |
| Desktop/web zoom | Preserve existing zoom behavior for apps that currently support it. | Configure separately from Android pinch; respect an existing app's deliberate locked view. |
| Scroll and fit | Ordinary content can scroll and relayout without forced whole-page shrinking. | Independent scroll axes and fit policy. LG selects no-scroll panels; Telegram selects scrolling content. |
| Touch delivery | Deliver ordered native edges, retain short taps, and cancel safely on pointer/lifecycle changes. | Apps choose press/release activation and which areas capture movement or app-owned gestures. |
| Touch feedback | Expose consistent pressed/disabled/focused state immediately. | Apps supply colors, styles, or a custom renderer. Preserve LG's background-only feedback. |
| Hold repeat | Disabled unless a control opts in; one initial action per press. | Configure repeat delay, interval, acceleration, enabled controls, and action callback. |
| Accessibility | Publish visible semantic controls with correct native bounds and disabled/focus state. | Apps supply labels, custom roles/actions, or a custom provider through the same geometry contract. |
| Back | Dismiss keyboard first, then ask the app to handle navigation, then finish the activity if unhandled. | An app callback may consume navigation. Integrate both legacy and predictive Back paths. |
| Lifecycle | Suspend drawing/input appropriately, cancel active gestures/repeats, and detach old listeners. | App hooks pause/resume game state or restore connection intent without owning the host cleanup. |
| Orientation/system UI | Host follows the app's declared window/orientation policy. | Portrait, sensor-landscape, fullscreen, edge-to-edge, and system-bar appearance belong to app profiles. |
| Permissions/services | Offer reusable lifecycle-safe request/result hooks. | Apps request only their capabilities, such as Bluetooth or camera/photo access; ordinary apps inherit no unrelated permissions. |
| Diagnostics | Optional platform/window/transform/input diagnostics with bounded output, disabled by default. | Apps may opt in and add domain fields without exposing credentials or message content. |

Overrides configure behavior. Density facts, correct coordinate conversion, UI-thread ownership, and cleanup remain invariants. An overridden feature must still obey its documented thread and lifecycle contract.

### Expected application profiles

| Setting | LG Remote on Android | Dino on Android | Telegram on Android |
|---|---|---|---|
| Orientation | Portrait | Sensor-landscape | Preserve its current declaration; support current window changes |
| Android page pinch | Disabled | Disabled | Disabled |
| Content scroll | Disabled for remote panels | Disabled for the game surface | Enabled for chat lists and messages |
| Fit/layout | Compact responsive panels; preserve final target sizes | Stable aspect/game coordinates | Reflow with pinned header/composer |
| Insets | Host-managed safe container unless an equivalent app-managed policy is selected | Fullscreen game with explicit cutout/gesture-safe controls | App-managed insets for pinned UI |
| Keyboard | IP field and text editing | No ordinary text-entry flow | Composer/search fields and keyboard dismissal |
| Repeat | Volume, channel, and directions opt in | Jump/duck semantics owned by the game; no automatic timer repeat | No ordinary button repeat unless a specific control opts in |
| Back | Keyboard, panel, activity | Game callback if needed, then activity | Keyboard, conversation/navigation callback, activity |
| App-owned gestures | Pointer pad and sensitivity slider | Tap/swipe jump and duck | Scrolling and attachment flows; media gestures only where implemented |
| Native capabilities | Bluetooth and local-network access | Existing game capabilities only | Camera/photo-picker capabilities only where used |

All three profiles consume common defaults and override only what their application needs. Verify that changing one app's profile cannot affect another app or another page instance.

---

## Phase 1: Reproduce sizing failures and establish the baseline

### 1.1 Samsung and Pixel comparison

- [ ] A1.01 Record the Samsung model, Android/One UI version, app build revision, physical resolution, logical window size, display zoom/density override, font setting, and navigation mode. Proof: a sanitized device record under `temp/android-shared-host/`.
- [ ] A1.02 Verify Pixel 7 and Samsung run the same APK by recording its hash and installed package information. Proof: matching artifact identity before comparing layout.
- [ ] A1.03 Reproduce the Samsung report in Remote, Pad, and Numbers panels. Record whether the failure is tiny controls, clipping, excess space, misplaced hit areas, or failure to relayout.
- [ ] A1.04 Capture the native container bounds, Ebiten logical viewport, page layout size, content extent, density, font conversion, insets, and final fit transform for both devices. Proof: one trace of the complete size pipeline per device.
- [ ] A1.05 Compare default and changed display-zoom/font settings without adding another density conversion. Proof: identify the first incorrect value or transformation, or retain the cause as unconfirmed.
- [ ] A1.06 Measure final displayed button sizes in dp after fitting. Proof: distinguish pre-fit CSS dimensions from final Android hit dimensions.

### 1.2 Existing app behavior

- [ ] A1.07 Record Dino's current landscape layout, jump/duck/tap/swipe behavior, pinch response, and pause/resume behavior on a named Android device.
- [ ] A1.08 Record Telegram's chat scrolling, pinned composer, keyboard behavior, Back navigation, attachments, and pinch response on a named Android device.
- [ ] A1.09 Inventory each example's native activity, exported Go bridge, gesture policy, inset handling, and listener/timer ownership. Proof: a migration map naming what moves and what stays app-specific.

Phase exit: the Samsung symptom has reproducible evidence or a precise missing-device record; app baselines exist before migration. A missing Samsung device leaves Samsung acceptance open while independent shared work can proceed.

## Phase 2: Define overridable shared contracts and Android packaging

### 2.1 API and ownership

- [ ] A2.01 Define independent Android policies for page pinch, scroll axes, viewport fitting, text scale, and inset ownership. Proof: no single flag must disable unrelated behaviors to satisfy an app profile.
- [ ] A2.02 Define unset/explicit-enabled/explicit-disabled option semantics. Proof: a zero-valued option cannot accidentally override a default or make a false override impossible.
- [ ] A2.03 Define option precedence as shared defaults, platform defaults, app profile, then explicit instance override. Proof: deterministic resolution tests cover conflicts and explicit false values.
- [ ] A2.04 Choose the mobile API entry point alongside [run.go](../../run.go), preserving existing `BindMobile` calls. Proof: document how existing callers obtain unchanged behavior and how new callers override it.
- [ ] A2.05 Define whether runtime overrides are supported per option. Proof: changing a gesture policy cancels incompatible active gestures; creation-only settings are identified explicitly.
- [ ] A2.06 Define a viewport snapshot with units, available window dimensions, four insets, IME state, font information, transform, and a revision. Proof: stale native geometry cannot override newer Go layout state.
- [ ] A2.07 Define generic control/action metadata for buttons, editable fields, ranges, capture regions, and repeat policies. Proof: the contract contains no LG control IDs, SSAP commands, Dino actions, or Telegram routes.
- [ ] A2.08 Define Back and lifecycle callbacks with thread ownership and handled/unhandled results. Proof: Android UI work and Go page mutation execute on their respective loops.

### 2.2 Reuse across Gradle projects

- [ ] A2.09 Select one shared Android host module/source location consumed by LG, Dino, and Telegram. Proof: a single native-host fix reaches all three builds without copying activity implementations.
- [ ] A2.10 Define how each generated mobile AAR supplies its package-specific callback adapter to the shared host. Proof: the shared Java code has no hardcoded import of LG's generated `Mobile` package.
- [ ] A2.11 Confirm the common-feature and application-profile tables above against the final API. Proof: every configurable feature has a documented override or callback, scope, default, and precedence.

Phase exit: contracts and packaging are reviewable, defaults preserve existing clients, and each app can express its own policy without duplicating host logic.

## Phase 3: Implement viewport, insets, font, and usable layout behavior

### 3.1 Correct sizing

- [ ] A3.01 Fix the evidenced Samsung sizing cause at the layer where the first incorrect value originates. Proof: the Phase 1 reproduction passes on the same Samsung configuration.
- [ ] A3.02 Carry native window/view changes through the shared viewport snapshot and existing `internal/window/resize.go` path. Proof: the committed layout matches the current usable window after resize settles.
- [ ] A3.03 Verify native-pixel to logical-unit conversion happens once for layout, touch input, and accessibility. Proof: multiple densities and display overrides yield equivalent logical placement.
- [ ] A3.04 Keep actual drawable window size distinct from app minimum-size clamps and logical game viewport size. Proof: short keyboard windows and narrow multi-window views still have correct fitting and hit areas.
- [ ] A3.05 Use one authoritative fit/coordinate transform for rendering, native input, and accessibility. Proof: letterboxed corners and control centers map to the same element.
- [ ] A3.06 Recompute geometry after display zoom, density, font settings, split-screen bounds, orientation, and activity recreation. Proof: final values do not depend on a stale first-launch snapshot.

### 3.2 Safe areas and text

- [ ] A3.07 Implement host-managed safe-area padding for all four system-bar/cutout edges. Proof: gesture navigation, button navigation, and landscape cutouts do not cover controls.
- [ ] A3.08 Implement app-managed inset reporting with explicit units and consumption rules. Proof: Telegram receives composer/header insets without an additional host-padding offset.
- [ ] A3.09 Implement IME visibility/space reporting across supported Android versions, including API 28-29 fallback behavior. Proof: focus, show/hide, and Back restore the usable viewport without double adjustment.
- [ ] A3.10 Use Android-supported sp conversion for nonlinear text scaling and refresh it on relevant configuration changes. Proof: small/default/large system settings produce the documented text sizes.
- [ ] A3.11 Replace LG's implicit 16-48 clamp with the resolved shared/app font policy. Proof: supported smaller text settings and enlarged text are handled intentionally.
- [ ] A3.12 Adapt LG's compact layouts/panel allocation so final hit targets remain usable while panels remain scroll-free. Proof: use reflow or additional panels where necessary; verify final Android targets rather than only layout boxes.
- [ ] A3.13 Verify app viewport/font overrides remain local to the selected app instance. Proof: Dino's game viewport does not impose LG's font or Telegram's composer geometry.

Phase exit: sizing, insets, and fonts have one unit contract; Samsung acceptance and final target-size acceptance remain open until device evidence exists.

## Phase 4: Share touch delivery, configurable pinch, and optional hold repeat

### 4.1 Native events and gesture ownership

- [ ] A4.01 Move ordered down/up/cancel delivery into the shared Android input adapter. Proof: taps shorter than a rendered frame produce exactly one accepted activation.
- [ ] A4.02 Define bounded input-queue behavior that preserves release/cancellation even under pressure. Proof: saturation cannot leave a button, game action, or repeat timer stuck down.
- [ ] A4.03 Route generic controls through the page's enabled state and metadata. Proof: disabled controls cannot activate through native touch or accessibility actions.
- [ ] A4.04 Replace the hardcoded LG pad exception with generic capture/gesture metadata. Proof: the pad, Dino's game surface, and Telegram's message scroll each receive the intended events.
- [ ] A4.05 Ensure native events and Ebiten's mouse/touch paths cannot activate the same physical press twice. Proof: mixed native/canvas traces yield one initial action per gesture.
- [ ] A4.06 Preserve per-control activation-on-press or activation-on-release overrides. Proof: a release-activated control cancels correctly when the pointer leaves its permitted area.

### 4.2 Pinch and repeat policies

- [ ] A4.07 Add an Android page-pinch gate independent of scrolling and fit. Proof: Telegram can scroll with page pinch disabled; an opted-in demo can still pinch.
- [ ] A4.08 Preserve supported desktop/browser zoom when only Android pinch is disabled. Proof: platform-specific options do not disable desktop keyboard/wheel zoom or change existing deliberate locks.
- [ ] A4.09 Verify an app-owned multi-touch/captured gesture can override page gesture ownership without accidental page zoom. Proof: capture, second-finger cancellation, and pointer release are deterministic.
- [ ] A4.10 Extract optional repeat scheduling with one initial action, delay, interval, and acceleration overrides. Proof: ordinary buttons remain single-action and only opted-in controls repeat.
- [ ] A4.11 Cancel repeats and discard unsent old-generation repeat work on release, slide-off, disable/removal, second finger, pause, and policy changes. Proof: queued repeats cannot begin after cancellation; already transmitted domain actions are reported separately.
- [ ] A4.12 Verify the LG profile preserves 350 ms start delay and 120/80/50 ms intervals after 0/2/4 seconds of holding. Proof: deterministic timer checks and recorded device behavior agree with the configured profile.
- [ ] A4.13 Batch native input state changes into bounded frame work without delaying action delivery behind unnecessary redraws. Proof: a three-tap burst within one second is retained and feedback timing is measured on named devices.

Phase exit: general pinch works for an opted-in Android app; the three migrated profiles can block page pinch independently; rapid taps and repeat cancellation are proven at the actual input boundary.

## Phase 5: Share native accessibility and geometry invalidation

- [ ] A5.01 Publish generic page control semantics from ownframe rather than LG-only key tables. Proof: ordinary app buttons, text fields, ranges, and visible text can be exposed without remote-specific metadata.
- [ ] A5.02 Move the native virtual-node provider into the shared Android host. Proof: the same provider operates in LG, Dino's menus/controls, and Telegram's lists/composer.
- [ ] A5.03 Rebuild native bounds when the viewport, native view size, density, transform, or layout revision changes, even if semantic JSON is unchanged. Proof: a resize with identical content leaves no stale hit regions.
- [ ] A5.04 Map bounds using the shared draw transform, including offsets, letterboxing, app-managed insets, and scroll position. Proof: spoken control targets match the visible controls.
- [ ] A5.05 Preserve the accessibility-disabled event guard. Proof: launching and interacting with accessibility disabled cannot dispatch invalid native events.
- [ ] A5.06 Preserve focus by stable control identity while controls survive, and clear or relocate it when they disappear. Proof: panel switches, chat transitions, resize, and disabled controls have valid focus outcomes.
- [ ] A5.07 Implement enabled/selected/pressed/editable/range semantics and route actions through the same app control dispatch as ordinary input. Proof: touch and TalkBack cannot bypass disabled state or create duplicate activation.
- [ ] A5.08 Expose text editing and range progress through generic callbacks. Proof: LG's IP field/sensitivity slider and Telegram's composer/search remain usable.
- [ ] A5.09 Define restrained status announcements without announcing every repeat acknowledgement. Proof: rapid remote steps do not overwhelm TalkBack speech.
- [ ] A5.10 Support app-provided names, custom actions, and a custom provider override without losing viewport invalidation. Proof: overrides remain compatible with shared geometry and lifecycle handling.
- [ ] A5.11 Verify TalkBack traversal and activation on Pixel 7 and the Samsung device. Proof: record device/version, traversal order, field/slider actions, and control alignment.

Phase exit: the shared provider works in more than one application and survives both accessibility-enabled and accessibility-disabled paths. Device checks remain open until performed.

## Phase 6: Share lifecycle, Back, and platform callback handling

- [ ] A6.01 Centralize Ebiten view resume/suspend ownership in the shared host. Proof: no duplicate game registration, resumed draw loop, or listener remains after recreation.
- [ ] A6.02 Cancel active gestures/repeat state before pausing and detach obsolete activity callbacks. Proof: backgrounding during a hold or drag cannot leave actions running.
- [ ] A6.03 Define app pause/resume hooks and verify they execute once per relevant transition. Proof: Dino can stop gameplay work without introducing TV or chat behavior into the host.
- [ ] A6.04 Reapply app policy and viewport/font state after recreation and process restart. Proof: options do not revert silently to defaults on a new activity.
- [ ] A6.05 Provide lifecycle-safe native capability callbacks so the LG adapter restores Bluetooth intent consistently with the Go mode. Proof: recreate in Bluetooth mode without forcing the user to toggle modes first.
- [ ] A6.06 Define ownership of queued actions at suspend/recreation, including stale generation rejection and important-action retention. Proof: an obsolete activity cannot send a late action into its replacement.
- [ ] A6.07 Share IME-first Back handling on API 28-29 and API 30+. Proof: Back dismisses the keyboard without also leaving the screen or activity.
- [ ] A6.08 Route an unconsumed Back action to an app callback, then to activity finish. Proof: LG panel navigation and Telegram conversation navigation each happen once.
- [ ] A6.09 Integrate predictive and legacy Back through the same callback order. Proof: one user action cannot run both paths or close an already-handled screen.
- [ ] A6.10 Add reusable permission/service result hooks with activity-generation ownership. Proof: denial, revocation, and late results do not crash a replacement activity or pause the game twice.
- [ ] A6.11 Keep permission sets app-specific. Proof: migrating Dino does not request Bluetooth/camera access, and Telegram does not inherit LG's TV pairing logic.

Phase exit: keyboard/Back and lifecycle behavior are shared; platform services retain app-owned intent and cannot act through obsolete activity instances.

## Phase 7: Migrate LG Remote, Dino, and Telegram

### 7.1 LG Remote

- [ ] A7.01 Replace LG's duplicated activity host, touch queue, and accessibility provider with the shared host and its generated-AAR adapter. Proof: the Android project consumes the shared native implementation.
- [ ] A7.02 Configure portrait, Android page pinch disabled, no panel scrolling, font policy, and hold-repeat controls through the LG app profile. Proof: explicit per-instance overrides resolve to the expected policies.
- [ ] A7.03 Keep SSAP/Bluetooth command mapping, pairing, saved keys, volume/channel reply handling, Power/Wake sequencing, and Hotstar discovery in the LG example. Proof: generic host packages import no TV domain code.
- [ ] A7.04 Verify all three panels, rapid taps, hold release, pad movement, sensitivity, background press feedback, and status-layout stability after migration. Proof: compare against the recorded baseline and final target-size requirements.

### 7.2 Dino

- [ ] A7.05 Update Dino's activity and mobile binding to consume the shared host. Proof: its build uses the same viewport/lifecycle/input implementation as LG.
- [ ] A7.06 Set Android page pinch disabled, game scrolling disabled, sensor-landscape, fullscreen behavior, and the game viewport policy explicitly. Proof: two-finger motion cannot scale or pan the page unexpectedly.
- [ ] A7.07 Preserve tap/swipe jump and duck behavior and the game's existing held-jump semantics. Proof: the common hold timer does not introduce repeated jumps or interfere with game input.
- [ ] A7.08 Verify cutout/gesture-safe controls, landscape rotation, display changes, activity recreation, and pause/resume. Proof: game state is handled by its documented app policy and controls remain reachable.
- [ ] A7.09 Expose meaningful menu/control accessibility where applicable, with a game-specific override for unsupported interactions. Proof: do not present the canvas as a list of LG-style remote buttons.

### 7.3 Telegram

- [ ] A7.10 Update Telegram's activity and mobile binding to consume the shared host while retaining the existing app-managed pinned-layout inset model. Proof: no host/app double padding.
- [ ] A7.11 Disable Android page pinch independently while keeping chat-list/message scrolling enabled. Proof: a two-finger gesture cannot scale the entire chat page and one-finger scrolling still works.
- [ ] A7.12 Preserve header/composer pinning, focus, search, typing, keyboard show/hide, and Back-to-chat-list behavior. Proof: resize and font changes keep the composer above the keyboard.
- [ ] A7.13 Preserve camera/photo-picker requests and delivered photo bytes through app-specific callbacks. Proof: returning from an external activity resumes the correct conversation without duplicate results.
- [ ] A7.14 Expose chat navigation, composer/search fields, and attachment controls through the shared accessibility provider. Proof: traversal and text editing use Telegram semantics and labels.

### 7.4 Remove duplicated implementations

- [ ] A7.15 Delete migrated private host/input/accessibility copies after all three callers use the shared APIs. Proof: one canonical implementation remains for each common feature; no unused legacy native classes remain.
- [ ] A7.16 Demonstrate an override in each app and an Android demo that explicitly enables page pinch. Proof: the shared defaults are reusable and every app profile is independently configurable.

Phase exit: all three Android builds consume common features and preserve their application workflows. Source migration alone does not satisfy the device acceptance rows.

## Phase 8: Verification, documentation, and completion

### 8.1 Focused automated evidence

- [ ] A8.01 Add meaningful tests for option precedence, explicit false overrides, runtime policy changes, and instance isolation. Proof: platform/app overrides are behaviorally exercised.
- [ ] A8.02 Add tests for coordinate conversion, native view-only resize invalidation, clamped/virtual viewports, insets, and final displayed touch sizes. Proof: tests reach the full size/transform chain.
- [ ] A8.03 Add event-path tests for sub-frame taps, multiple presses within one second, duplicate native/Ebiten delivery, queue saturation, cancellation, and repeat generations. Proof: exercise the actual dispatch seam rather than a standalone timer only.
- [ ] A8.04 Add native checks for accessibility-disabled launch, virtual-node geometry changes, keyboard Back, callback detachment, and activity recreation. Proof: Android host integration is exercised beyond Go unit tests.
- [ ] A8.05 Run `gofmt` on changed Go files and confirm each edited/new Go file is at most 2000 characters using `wc -m`. Proof: record formatting and size results.
- [ ] A8.06 Run `make test TEST_P=4` once on the final implementation tree, or use the default serial setting if memory pressure requires it. Proof: record command, revision, exit status, and log; do not substitute `go build ./...`.
- [ ] A8.07 Run `make build BUILD_P=4` on the final tree. Proof: the root and examples compile through the repository's vet-based build check.
- [ ] A8.08 Run focused race checks for the changed host/bridge/page/window state boundaries. Proof: avoid repeating the full expensive layout matrix under race detection without a new concern.
- [ ] A8.09 Build LG, Dino, and Telegram Android artifacts through their supported scripts/Gradle tasks. Proof: list each artifact hash, revision, supported ABI, SDK/toolchain, and successful build result.

### 8.2 Real-device acceptance

- [ ] A8.10 Complete the mandatory rows in the device matrix below on Pixel 7 and the reported Samsung device. Proof: screenshots/trace records identify app, build, device, settings, expected behavior, and result.
- [ ] A8.11 Verify an Android 9-10 path and an Android 11+ path for keyboard/insets/Back, using physical devices or clearly labeled emulators. Proof: do not equate an emulator pass with Samsung hardware acceptance.
- [ ] A8.12 Measure app input feedback separately from TV command acknowledgement and actual TV action. Proof: report observed measurements and transport conditions; no claim of instant response from queue acceptance alone.
- [ ] A8.13 Check bounded queues, listeners, timers, and retained activity state through repeated pause/resume/recreation cycles. Proof: no growing backlog or old activity callback remains; record the cycle count and observation method.

### 8.3 Documentation and closure

- [ ] A8.14 Document common Android setup, feature defaults, override precedence, runtime/creation-only options, inset ownership, and pinch/scroll independence. Proof: a new app can bind without copying an example activity.
- [ ] A8.15 Update LG, Dino, and Telegram usage/build documentation with their actual Android profiles and remaining device limitations. Proof: docs match the migrated behavior and artifact commands.
- [ ] A8.16 Update this file's evidence log and check only rows with matching implementation and verification. Proof: unavailable Samsung, TalkBack, Bluetooth, and actual TV checks remain open.
- [ ] A8.17 Close the workstream only when all common features have documented overrides, all three apps are migrated, final verification passes, and mandatory device checks are complete.

Phase exit: shared Android behavior and its overrides are demonstrated across the three apps, with truthful device evidence and no unresolved mandatory acceptance rows.

## Device and configuration matrix

Every row is planned verification. Add more cases when the Samsung reproduction reveals a device-specific dependency.

| ID | Scenario | Expected behavior | Required applications/evidence |
|---|---|---|---|
| D01 | Pixel 7 and reported Samsung, identical APK build | Logical layout follows the available window; no stale or double density conversion | LG panels plus shared host trace; record both device identities |
| D02 | Default, smaller, and enlarged display zoom/font settings | Reflow and text conversion follow policy; controls remain usable after any fit | All three apps; final dp target measurements |
| D03 | Gesture navigation and three-button navigation | Controls avoid reserved system areas; available height updates | All three apps on available navigation modes |
| D04 | Keyboard open/close and Back | Focused IP field/composer remains reachable; Back dismisses IME first | LG and Telegram; Android 9-10 and 11+ paths |
| D05 | Portrait/landscape and short/narrow multi-window bounds | App orientation/viewport policy holds; letterboxing and hit areas agree | Dino landscape; LG/Telegram declared-orientation and allowed window changes |
| D06 | Cutout on either landscape edge | Four-edge insets protect controls, with one owner of padding | Dino and a supported rotated/windowed UI case |
| D07 | Two-finger page pinch | LG/Dino/Telegram Android pages keep their scale; an opted-in demo zooms correctly | All migrated apps plus explicit pinch-enabled demo |
| D08 | Scrolling with page pinch disabled | Telegram scrolls normally; LG and Dino do not pan their page | All three app profiles |
| D09 | Three taps within one second and a longer rapid-tap sequence | One initial app action per accepted press; no dropped edge or duplicate activation | LG steps/menu controls; Dino actions; Telegram navigation/controls |
| D10 | Hold, slide off, add finger, pause, disable/remove control | Only opted-in controls repeat; cancellation stops new repeats and stale unsent work | LG repeat profile and generic repeat harness |
| D11 | TalkBack enabled/disabled and size change | No startup crash; virtual nodes, focus, fields and ranges match the canvas | LG, Dino menus, Telegram conversation/composer |
| D12 | Activity recreation/process restart | Reapply options and geometry; restore documented app state; no obsolete callbacks | All three apps; include LG Bluetooth mode |
| D13 | Background/foreground and external permission/media activity | Suspend/resume occurs once; gestures stop and results reach the current app | Dino gameplay; LG permissions; Telegram camera/photo picker |
| D14 | Bluetooth off/on, unavailable TV, denial/revocation and reconnect | LG reports real connection state and retries through its app-owned adapter safely | LG only; paired TV evidence |
| D15 | Actual TV volume/channel response to taps and holds | Accepted actions map to observed TV steps; distinguish network latency from input delivery | LG only; connected TV, timing method and counts |
| D16 | Desktop/browser regression after Android-only policy changes | Existing supported zoom, scrolling, fields and controls retain their policy | Representative desktop/wasm runs; no Android evidence substitution |

## Dependencies and implementation order

1. Phase 1 supplies reproduction evidence and current app behavior. Phase 2 defines the shared interfaces and packaging.
2. Phase 3 depends on the viewport/inset/font contracts from Phase 2. The Samsung-specific closure depends on access to the affected device and settings.
3. Phase 4 uses the geometry contract and independent gesture options. Phase 5 uses the same geometry and action metadata.
4. Phase 6 uses the callback and event-generation contracts. It must be ready before activity recreation is accepted during migration.
5. Phase 7 requires the shared host, touch/accessibility adapters, and lifecycle behavior. Migrate every named caller and delete its old implementation in the same completed migration unit.
6. Phase 8 checks each implementation as it lands and performs final cross-app/device closure. Do not postpone basic compile and focused correctness checks until all migrations are complete.

The public API remains at the module root. Generic page state belongs in `internal/page`; host interfaces belong in `internal/host`; frame/input/viewport integration belongs in `internal/window`. Select a shared native Android module in Phase 2. Application packages keep navigation, game rules, TV protocols, connection policies, camera/photo workflows, and visual design.

Existing [`dynamic-resize.md`](dynamic-resize.md) and [`input-interaction.md`](input-interaction.md) remain the historical workstream plans. This file owns the new Android common-feature/override/migration work. It does not reopen their completed desktop work or duplicate their active rows.

## Validation policy and current scope

For this planning-only change, check Markdown structure, local links, row IDs, and worktree scope. Do not run application tests, lint, builds, or device automation just to create this file.

For implementation, use the verification rows above and the repository's actual Make targets. The checklist guide's generic `make lint` example does not match this repository; `make build` runs the required vet-based compile check. A later explicit user instruction to skip checks takes precedence: record skipped checks and leave the matching acceptance rows open.

Maintain the current Go module boundaries, the pinned blinkless engine, and the no-JavaScript architecture. Do not add a dependency or upgrade the engine to conceal an unexplained Samsung issue. Leave `~/.Xauthority` untouched and keep clipboard tests on `clipboard.UseMemory`.

## Evidence log

| Date/revision | Item IDs | Command or workflow | Platform/device/settings | Result | Artifact |
|---|---|---|---|---|---|
| 2026-10-09 / `d679bd8` | Baseline only | Read current host, LG/Dino/Telegram activities, mobile bindings, policies and LG tests | WSL2 source review | Shared support and example-specific duplication identified; no device acceptance closed | Source links in this file; [LG PR record](../PR/pr-lg-remote.md) |

For each completed row, append its exact IDs, source revision, command/manual workflow, device or renderer, relevant settings, observed result, and ignored artifact path. Include failure evidence and the next check for partial work. A screenshot alone does not prove tap routing, repeat cancellation, or network response.

## Reference material

- [Android density guidance](https://developer.android.com/training/multiscreen/screendensities)
- [Responsive Android views](https://developer.android.com/develop/ui/views/layout/responsive-adaptive-design-with-views)
- [Android nonlinear font scaling](https://developer.android.com/about/versions/14/features#non-linear-font-scaling)
- [Android accessibility and touch targets](https://developer.android.com/guide/topics/ui/accessibility/views/apps-views)
- [ownframe interaction guide](../../documentation/interaction.md)
- [ownframe feature guide](../../documentation/features.md)
- [Repository checklist guide](../../skills/phase-wise-checklist/SKILLS.md)
