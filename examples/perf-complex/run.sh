#!/bin/sh
set -eu
ROOT=$(CDPATH= cd -- "$(dirname "$0")/../.." && pwd)
cd "$ROOT"
BASE=examples/perf-complex
OUT=temp/complex-dump
mkdir -p "$OUT"
# One named executable, outside the checkout. Never build all examples.
BIN=$(mktemp /tmp/gpui-complex-dump.XXXXXX)
trap 'rm -f "$BIN"' EXIT HUP INT TERM
go build -o "$BIN" "./$BASE"
go version -m "$BIN" > "$OUT/build.txt"
python3 - <<'MANIFEST'
import datetime,hashlib,json,os,platform,subprocess
from pathlib import Path
base=Path('examples/perf-complex')
files=[p for p in base.iterdir() if p.is_file()]
m={'captured_utc':datetime.datetime.now(datetime.timezone.utc).isoformat(),
   'revision':subprocess.check_output(['git','rev-parse','HEAD'],text=True).strip(),
   'branch':subprocess.check_output(['git','branch','--show-current'],text=True).strip(),
   'host':platform.platform(),'cpu':'Intel Core i7-13700HX, WSL2, 24 logical CPUs',
   'gomaxprocs_env':os.environ.get('GOMAXPROCS','unset'),
   'go_version':subprocess.check_output(['go','version'],text=True).strip(),
   'source_sha256':{str(p):hashlib.sha256(p.read_bytes()).hexdigest() for p in files}}
(Path('temp/complex-dump')/'manifest.json').write_text(json.dumps(m,indent=2)+'\n')
MANIFEST
for run in 1 2 3; do
  for mode in initial cached data resize windowed; do
    "$BIN" -dump -mode "$mode" -out "$OUT/run-$run"
    go tool pprof -top -nodecount=35 "$OUT/run-$run/$mode-cpu.pprof" > "$OUT/run-$run/$mode-cpu.txt"
    go tool pprof -top -sample_index=alloc_space -nodecount=25 "$OUT/run-$run/$mode-heap.pprof" > "$OUT/run-$run/$mode-alloc.txt"
    go tool pprof -top -sample_index=inuse_space -nodecount=25 "$OUT/run-$run/$mode-heap.pprof" > "$OUT/run-$run/$mode-retained.txt"
  done
done
python3 "$BASE/analyze.py" > "$OUT/analysis.txt"
