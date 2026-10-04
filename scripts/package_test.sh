#!/bin/sh
# Check the dry-run layout of scripts/package.sh against the layouts in
# documentation/packaging.md. No build and no network, so it stays fast.
set -eu

root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)
cd "$root"

fail=0

want() { # $1 output, $2 literal line, $3 label
	case $1 in
	*"$2"*) ;;
	*)
		echo "package layout: missing $3: $2" >&2
		fail=1
		;;
	esac
}

linux=$(GOOS=linux sh scripts/package.sh -n print)
want "$linux" "archive: dist/go-gpui-print-linux-" "Linux archive"
want "$linux" "  go-gpui-print/print" "Linux binary"
want "$linux" "  go-gpui-print/print.desktop" "Linux desktop entry"
want "$linux" "  go-gpui-print/README.md" "Linux README"
want "$linux" "  go-gpui.wasm" "wasm binary"
want "$linux" "  wasm_exec.js" "wasm loader"
want "$linux" "  index.html (browser/index.html)" "wasm page"
want "$linux" "checksums: dist/SHA256SUMS" "checksums"

darwin=$(GOOS=darwin sh scripts/package.sh -n print)
want "$darwin" "archive: dist/go-gpui-print-macos-" "macOS archive"
want "$darwin" "  go-gpui-print.app/Contents/Info.plist" "macOS plist"
want "$darwin" "  go-gpui-print.app/Contents/MacOS/print" "macOS binary"

windows=$(GOOS=windows sh scripts/package.sh -n print)
want "$windows" "archive: dist/go-gpui-print-windows-" "Windows archive"
want "$windows" "  go-gpui-print.exe" "Windows executable"

if [ "$fail" -ne 0 ]; then
	exit 1
fi

echo "package layout ok"
