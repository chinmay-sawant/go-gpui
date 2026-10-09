#!/bin/sh
# Build the LG remote as an Android debug APK.
#
# Usage:
#   scripts/android-lg.sh           bind the Go page, build the APK, print it
#   scripts/android-lg.sh install   build, then adb install -r the APK
#
# From WSL2 the phone is visible to the Windows adb, not the Linux one.
# The script uses adb.exe under /mnt/c/Users/*/platform-tools when ADB is unset.
# When ADB ends in .exe the APK path is handed over with wslpath -w.
set -eu

root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)
android="$root/examples/lg-remote/android"
aar="$android/lgremote/lgremote.aar"
apk="$android/app/build/outputs/apk/debug/app-debug.apk"
target=${ANDROID_TARGET:-android/arm64}

if [ -z "${ADB:-}" ]; then
	for candidate in \
		/mnt/c/Users/*/platform-tools/adb.exe \
		/mnt/c/Users/*/AppData/Local/Android/Sdk/platform-tools/adb.exe
	do
		if [ -x "$candidate" ]; then
			ADB=$candidate
			break
		fi
	done
fi

adb=${ADB:-adb}

if ! command -v ebitenmobile >/dev/null 2>&1; then
	echo "android: ebitenmobile is not on PATH; install it with" >&2
	echo "  go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@v2.10.4" >&2
	exit 1
fi

if [ -z "${ANDROID_HOME:-}" ] && [ -z "${ANDROID_SDK_ROOT:-}" ]; then
	if [ -d "$HOME/Android/Sdk" ]; then
		ANDROID_HOME="$HOME/Android/Sdk"
	else
		echo "android: set ANDROID_HOME to the Android SDK" >&2
		exit 1
	fi
fi

export ANDROID_HOME="${ANDROID_HOME:-${ANDROID_SDK_ROOT:-}}"

# gomobile resolves the packages under the current directory, where the
# committed go.work points ownframe at this checkout. Do not export GOWORK.
cd "$root"

echo "==> ebitenmobile bind -target $target"
ebitenmobile bind -target "$target" \
	-ldflags "-s -w -extldflags '-Wl,-z,max-page-size=16384 -Wl,-z,common-page-size=16384'" \
	-javapkg com.chinmaysawant.lgremote \
	-o "$aar" ./examples/lg-remote/mobile

echo "==> ./gradlew :app:assembleDebug"
(cd "$android" && ./gradlew --no-daemon :app:assembleDebug)

echo "APK: $apk"

if [ "${1:-}" = "install" ]; then
	if ! command -v "$adb" >/dev/null 2>&1 && [ ! -x "$adb" ]; then
		echo "android: adb is not on PATH; set ADB to the Windows adb.exe" >&2
		exit 1
	fi

	install_path=$apk
	case "$adb" in
	*.exe) install_path=$(wslpath -w "$apk") ;;
	esac

	"$adb" install -r "$install_path"
fi
