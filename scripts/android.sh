#!/bin/sh
# Build the Telegram demo as an Android debug APK.
#
# Usage:
#   scripts/android.sh           bind the Go page, build the APK, print it
#   scripts/android.sh install   build, then adb install -r the APK
#
# One-time setup is the Android SDK (platform-tools, a platform, build-tools,
# an NDK) plus ebitenmobile. examples/telegram/android/README.md has the
# commands for a fresh machine.
set -eu

root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)
android="$root/examples/telegram/android"
aar="$android/telegram/telegram.aar"
apk="$android/app/build/outputs/apk/debug/app-debug.apk"

if ! command -v ebitenmobile >/dev/null 2>&1; then
	echo "android: ebitenmobile is not on PATH; install it with" >&2
	echo "  go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@v2.10.4" >&2
	exit 1
fi

if [ -z "${ANDROID_HOME:-}" ] && [ -z "${ANDROID_SDK_ROOT:-}" ]; then
	if [ -d "$HOME/Android/Sdk" ]; then
		ANDROID_HOME="$HOME/Android/Sdk"
	else
		echo "android: set ANDROID_HOME to the Android SDK (see the README)" >&2
		exit 1
	fi
fi

export ANDROID_HOME="${ANDROID_HOME:-${ANDROID_SDK_ROOT:-}}"

# gomobile resolves the packages under the current directory, where the
# committed go.work points go-gpui at this checkout. Do not export GOWORK:
# the generated temp module it builds in cannot join the workspace.
cd "$root"

echo "==> ebitenmobile bind -target android"
ebitenmobile bind -target android \
	-javapkg com.chinmaysawant.telegram \
	-o "$aar" ./examples/telegram/mobile

echo "==> ./gradlew :app:assembleDebug"
(cd "$android" && ./gradlew --no-daemon :app:assembleDebug)

echo "APK: $apk"

if [ "${1:-}" = "install" ]; then
	if ! command -v adb >/dev/null 2>&1; then
		echo "android: adb is not on PATH; install platform-tools" >&2
		exit 1
	fi

	adb install -r "$apk"
fi
