# LG remote

A remote for an LG webOS TV, aimed at a UP7750PTZ. The page starts dark. The corner control switches to the light theme.

Desktop control is Wi-Fi only. The Android build adds Bluetooth.

## Desktop

The TV and the computer have to be on the same network. The TV has to be on for the first pairing.

```
go run ./examples/lg-remote
```

On open the page searches the Wi-Fi and fills in the TV it finds. Tap Scan to search again. Accept the prompt on the TV the first time. The client key is saved under the user config directory, in `ownframe/lg-remote.json`.

Scan looks for a webOS TV with SSDP. Wake sends a Wake-on-LAN packet to the MAC learned while the TV was on. Power turns the TV off. Channel, volume, mute, the pad, digits, colours, playback, and the rest of the webOS button names are on the page.

`-web` serves the same picture on `127.0.0.1:8131`. That mode does not tick, so connection status stays on the last draw.

## Android

`scripts/android-lg.sh` binds `examples/lg-remote/mobile` and builds a debug APK. The activity is `com.chinmaysawant.lgremote`.

```
sh scripts/android-lg.sh install
```

From WSL2 the script prefers the Windows `adb.exe` under `/mnt/c/Users/*/platform-tools`, because the Linux adb cannot see a phone plugged into the Windows host. Set `ADB` to override it.

Wi-Fi works the same way as the desktop app. Bluetooth makes the phone a HID keyboard. Tap Bluetooth, allow the radio prompt, then on the TV open the Bluetooth device list and pair this phone. Arrows, OK, back, digits, volume, channel, mute, and playback go over Bluetooth. Power, home, and menu send HID consumer controls whose support depends on the TV. App launches, colour keys, and inputs need Wi-Fi. Wake still uses the network.

The phone shows each panel without scrolling. The connection controls share a row, and the whole panel fits within the available screen when its content is taller. Buttons have a minimum 48 CSS pixel target before fitting and rows wrap on narrow screens. Both themes highlight pressed and focused buttons through their background colour. The activity reserves space for system bars, display cutouts, and the soft keyboard, and passes the system text scale to the page. Back closes the keyboard first, returns Pad or Numbers to Remote, then leaves the activity.

Drag inside the Wi-Fi pad to move the TV pointer while your finger is down. The sensitivity slider adjusts movement from 0.25x to 3.00x. Bluetooth disables controls it cannot send. Its queue preserves accepted commands and reports when it is full. "Queued for Bluetooth" means the command was accepted locally; Android reports whether the HID report was sent, which does not confirm that the TV acted on it. Wi-Fi commands run in order on one worker.

The HTML uses named buttons and a labelled IP field. Tab moves through enabled controls; Enter or Space activates the focused button. The Android activity exposes visible controls, status, hints, text editing, and scrolling through a native accessibility node provider because Ebiten draws a canvas rather than browser DOM elements.

Layout tests cover widths from 240 to 800 CSS pixels, portrait and landscape viewports, both themes, all three panels, long status text, and base font sizes from 16 to 48 pixels. Real-device checks are still needed for TalkBack navigation, keyboard and cutout handling across Android versions, Bluetooth permission denial and reconnects, and command delivery to a paired TV.
