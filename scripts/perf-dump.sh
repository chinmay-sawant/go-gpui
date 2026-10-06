#!/bin/sh
# Automatic perf dump: headless Redraw of each bench app to JSON.
# Usage: scripts/perf-dump.sh [--out DIR]
# Display-free: go test only, never opens a window.
set -eu
ROOT=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)
OUT="$ROOT/temp/perf-dumps"
while [ $# -gt 0 ]; do
  case "$1" in
    --out) OUT="$2"; shift 2;;
    -h|--help) sed -n '2,4p' "$0"; exit 0;;
    *) echo "unknown flag $1" >&2; exit 1;;
  esac
done
mkdir -p "$OUT"
PERF_DUMP_DIR="$OUT" go test -p 1 ./examples/perf-stress/ -run TestDump -v -count 1
echo "Dumps in $OUT"
