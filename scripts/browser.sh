#!/bin/sh
# Build the login window as WebAssembly and serve it.
# Open the printed address. Resizing the browser changes the frame.
set -eu

root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)
out=$(mktemp -d)
trap 'rm -rf "$out"' EXIT

GOOS=js GOARCH=wasm go build -C "$root" -o "$out/go-gpui.wasm" ./cmd/go-gpui

goroot=$(go env GOROOT)
if [ -f "$goroot/lib/wasm/wasm_exec.js" ]; then
  cp "$goroot/lib/wasm/wasm_exec.js" "$out/wasm_exec.js"
else
  cp "$goroot/misc/wasm/wasm_exec.js" "$out/wasm_exec.js"
fi

cp "$root/browser/index.html" "$out/index.html"

echo "http://127.0.0.1:8092/"
exec python3 -m http.server 8092 --directory "$out"
