#!/usr/bin/env bash
# Send the catnotify burst one message at a time, pausing between each so the
# cat can show every bubble instead of having them overwrite each other.
# Usage: bash burst.sh [delay-seconds]   (default 1.5)
set -u

response_file=$(mktemp)
trap 'rm -f "$response_file"' EXIT

dir="$(cd "$(dirname "$0")" && pwd)"
file="$dir/burst.jsonl"
delay="${1:-1.5}"
total=$(wc -l < "$file")
i=0

while IFS= read -r line; do
  [ -z "$line" ] && continue
  i=$((i + 1))
  code=$(curl --silent --show-error --max-time 3 \
    -o "$response_file" -w '%{http_code}' \
    http://127.0.0.1:6969/notify \
    -H 'Content-Type: application/json' \
    --data-binary "$line")
  printf '%02d/%d http=%s\n' "$i" "$total" "$code"
  if [ "$i" -lt "$total" ]; then
    sleep "$delay"
  fi
done < "$file"

printf 'done: %d notifications sent %ss apart\n' "$i" "$delay"
