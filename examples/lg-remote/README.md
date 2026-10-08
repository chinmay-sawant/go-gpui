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

Wi-Fi works the same way as the desktop app. Bluetooth makes the phone a HID keyboard. Tap Bluetooth, allow the radio prompt, then on the TV open the Bluetooth device list and pair this phone. Arrows, OK, back, digits, volume, channel, mute, and playback go over Bluetooth. App launches, colour keys, inputs, and screen power stay on Wi-Fi. Wake still uses the network.
