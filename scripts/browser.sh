#!/bin/sh
# Build one example as WebAssembly and serve it.
# The first argument is the example directory name; login is the default.
# Open the printed address. Resizing the browser changes the frame. A file
# dropped on the canvas arrives through the same Update as a desktop drop.
set -eu

name=${1:-login}
port=${OWNFRAME_BROWSER_PORT:-${GPUI_BROWSER_PORT:-8092}}

root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)

if [ ! -d "$root/examples/$name" ]; then
  echo "no example named $name under examples/" >&2
  exit 1
fi

out=$(mktemp -d)
trap 'rm -rf "$out"' EXIT

GOOS=js GOARCH=wasm go build -C "$root" -p 1 -o "$out/ownframe.wasm" "./examples/$name"

goroot=$(go env GOROOT)
if [ -f "$goroot/lib/wasm/wasm_exec.js" ]; then
  cp "$goroot/lib/wasm/wasm_exec.js" "$out/wasm_exec.js"
else
  cp "$goroot/misc/wasm/wasm_exec.js" "$out/wasm_exec.js"
fi

if [ -f "$root/examples/$name/browser.html" ]; then
  cp "$root/examples/$name/browser.html" "$out/index.html"
else
  cp "$root/browser/index.html" "$out/index.html"
fi

echo "http://127.0.0.1:$port/"
python3 -m http.server "$port" --directory "$out"
