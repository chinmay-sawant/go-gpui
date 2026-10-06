# Showcase

Screenshots of the examples, captured on real hardware.

## Telegram demo on a Pixel 7

`examples/telegram` is the phone-first example. The screens below are the
same HTML template that `go run ./examples/telegram` opens on the desktop,
bound to Android with `ebitenmobile`. `sh scripts/android.sh install` binds
it, builds a debug APK with Gradle, and installs it over `adb`; the one-time
setup is in [examples/telegram/android/README.md](examples/telegram/android/README.md).

| Chats | Conversation | Contacts |
|-------|--------------|----------|
| ![Chat list with search, unread badges, pinned and muted chats](assets/telegram-chats.webp) | ![Conversation with bubbles, timestamps, read ticks, and a composer](assets/telegram-thread.webp) | ![Contacts tab with last-seen times](assets/telegram-contacts.webp) |

| Settings | Settings, dark | Chats, dark | Conversation, dark |
|----------|----------------|-------------|--------------------|
| ![Settings tab with a dark mode toggle](assets/telegram-settings.webp) | ![Settings in the dark theme](assets/telegram-settings-dark.webp) | ![Chat list in the dark theme](assets/telegram-chats-dark.webp) | ![Conversation in the dark theme](assets/telegram-thread-dark.webp) |

Taps, tabs, opening a chat, and drag-to-scroll work on the phone. Typing
needs a hardware keyboard, because this library version does not wire the
Android soft keyboard yet. The arm64 debug APK is 29 MB; binding all four
ABIs puts one APK at 188 MB.
