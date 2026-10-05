#!/bin/sh
# Long-run soak driver: 10-30min resize/scroll/click/text/animate loop.
# Usage: scripts/perf-longrun.sh [--minutes N] [--url URL] [--out CSV]
# POSIX sh. Samples /debug/state periodically into a CSV log.
set -eu
MINS=15
URL=http://127.0.0.1:8128/debug/state
OUT=/tmp/opencode/perf-soak.csv
while [ $# -gt 0 ]; do
  case "$1" in
    --minutes) MINS="$2"; shift 2;;
    --url) URL="$2"; shift 2;;
    --out) OUT="$2"; shift 2;;
    -h|--help) sed -n '2,4p' "$0"; exit 0;;
    *) echo "unknown flag $1" >&2; exit 1;;
  esac
done
mkdir -p "$(dirname "$OUT")"
echo "Soak plan (${MINS}min). Drive the app by hand or via browser:"
echo "  loop: resize | scroll | click | type text | toggle animation |"
echo "        create/remove elements (repeat, vary order)"
echo "Sampling $URL every 30s -> $OUT"
echo "ts,redraws,parses,cascades,layouts,repaints,heap_alloc,goroutines" > "$OUT"
END=$((MINS * 60)); EL=0
while [ "$EL" -lt "$END" ]; do
  TS=$(date -u +%FT%TZ)
  if command -v curl >/dev/null 2>&1; then
    curl -s "$URL" 2>/dev/null | python3 -c \
      'import json,sys; d=json.load(sys.stdin); s=d.get("stats") or {}; r=d.get("runtime") or {}; print(",".join(map(str,["'"$TS"'",s.get("Redraws",""),s.get("Parses",""),s.get("Cascades",""),s.get("Layouts",""),s.get("Repaints",""),r.get("heap_alloc",""),r.get("goroutines","")])))' \
      >> "$OUT" 2>/dev/null || echo "$TS,fetch-failed" >> "$OUT"
  else
    echo "$TS,no-curl" >> "$OUT"
  fi
  sleep 30; EL=$((EL+30))
done
echo "Done. Watch for: RSS/heap growth, redraw count exploding, errors."
echo "Log: $OUT"
