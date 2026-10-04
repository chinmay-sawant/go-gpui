#!/bin/sh
# Build one example as WebAssembly and serve it.
# The first argument is the example directory name; login is the default.
# Open the printed address. Resizing the browser changes the frame. A file
# dropped on the canvas arrives through the same Update as a desktop drop.
set -eu

name=${1:-login}

root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)

if [ ! -d "$root/examples/$name" ]; then
  echo "no example named $name under examples/" >&2
  exit 1
fi

out=$(mktemp -d)
trap 'rm -rf "$out"' EXIT

GOOS=js GOARCH=wasm go build -C "$root" -o "$out/go-gpui.wasm" "./examples/$name"

goroot=$(go env GOROOT)
if [ -f "$goroot/lib/wasm/wasm_exec.js" ]; then
  cp "$goroot/lib/wasm/wasm_exec.js" "$out/wasm_exec.js"
else
  cp "$goroot/misc/wasm/wasm_exec.js" "$out/wasm_exec.js"
fi

cp "$root/browser/index.html" "$out/index.html"

echo "http://127.0.0.1:8092/"
exec python3 -m http.server 8092 --directory "$out"
