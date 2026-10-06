#!/bin/sh
# Idle perf capture: 30-60s of an untouched window.
# Usage: scripts/perf-idle.sh [--example DIR] [--web] [--seconds N] [--url URL]
# POSIX sh. Prints a small CPU/RSS/heap/goroutine table.
set -eu
EX=examples/platform
WEB=0
SECS=60
URL=http://127.0.0.1:8128/debug/state
if [ ! -d "$EX" ]; then EX=examples/login; fi
while [ $# -gt 0 ]; do
  case "$1" in
    --example) EX="$2"; shift 2;;
    --web) WEB=1; shift;;
    --seconds) SECS="$2"; shift 2;;
    --url) URL="$2"; shift 2;;
    -h|--help) sed -n '2,6p' "$0"; exit 0;;
    *) echo "unknown flag $1" >&2; exit 1;;
  esac
done
echo "example=$EX web=$WEB secs=$SECS"
echo "Leave the window untouched for $SECS s (no input, no resize)."
if [ "$WEB" -eq 1 ]; then
  echo "Start: go run ./$EX -web &  then curl $URL"
  echo "Sampling $URL every 5s..."
  i=0
  printf '%-6s %-10s %-10s\n' "t(s)" "heap_alloc" "goroutines"
  while [ "$i" -lt "$SECS" ]; do
    if command -v curl >/dev/null 2>&1; then
      curl -s "$URL" 2>/dev/null | python3 -c \
        'import json,sys; d=json.load(sys.stdin); r=d.get("runtime") or {}; print("%-6s %-10s %-10s" % ("?", r.get("heap_alloc","-"), r.get("goroutines","-")))' \
        2>/dev/null || echo "state fetch failed (is -web running?)"
    else
      echo "curl not found; open $URL in a browser" >&2
    fi
    sleep 5; i=$((i+5))
  done
  exit 0
fi
echo "Manual desktop capture (no display automation here):"
echo "  1) go run ./$EX &  PID=\$!"
echo "  2) sleep $SECS  (do not touch the window)"
echo "  3) ps -o pid,pcpu,rss,etime -p \$PID"
echo "PID      %CPU     RSS(KB)  ELAPSED"
if command -v pgrep >/dev/null 2>&1; then
  pgrep -f "$EX" | head -5 | while read -r p; do
    ps -o pid=,pcpu=,rss=,etime= -p "$p" 2>/dev/null || true
  done
else
  echo "(pgrep missing: run ps against your window PID manually)"
fi
echo "Idle target: ~0% CPU, flat RSS across the window."
