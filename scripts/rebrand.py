#!/usr/bin/env python3
"""Apply or preview the tracked ownframe name migration."""
import argparse
import re
import subprocess
from pathlib import Path

root = Path(__file__).resolve().parent.parent
parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--apply', action='store_true')
parser.add_argument('--check', action='store_true', help='fail if source names still need migration')
parser.add_argument('--include-local', action='store_true', help='also migrate ignored temp Go probes')
args = parser.parse_args()
paths = subprocess.check_output(['git', 'ls-files', '-z'], cwd=root).decode().split('\0')
paths += subprocess.check_output(['git', 'ls-files', '--others', '--exclude-standard', '-z'], cwd=root).decode().split('\0')
if args.include_local:
    paths += [str(p.relative_to(root)) for p in (root / 'temp').rglob('*.go')]
text_types = {'.go', '.md', '.mod', '.sum', '.work', '.js', '.jsx', '.json', '.html',
              '.css', '.sh', '.py', '.java', '.xml', '.gradle', '.properties',
              '.yml', '.yaml', '.txt', '.toml', '.mjs', '.c', '.h'}
changed = []
for name in sorted(set(paths)):
    path = root / name
    if not name or name.startswith('docs/') or name in {'scripts/rebrand.py', 'documentation/rebrand.md'}:
        continue
    if path.suffix not in text_types and path.name not in {'Makefile', '.gitignore', 'go.work.sum'}:
        continue
    try:
        old = path.read_text()
    except (UnicodeError, FileNotFoundError):
        continue
    text = old
    # These are real upstream branch names and the existing checkout location.
    protected = ['chore/changes-for-go-gpui', str(root), 'go-gpui:theme',
                 'go-gpui:github-stars:v1', 'GPUI_PRINT_DEBUG', 'GPUI_FILEPICK_DEBUG',
                 'GPUI_BROWSER_PORT', 'GPUI_REPLAY_GPU_TEST',
                 'Formerly known as **go-gpui**.', 'formerly known as go-gpui.',
                 'VITE_GITHUB_REPOSITORY=chinmay-sawant/go-gpui',
                 'Migrate from go-gpui to ownframe']
    for i, value in enumerate(protected):
        text = text.replace(value, f'__REBRAND_KEEP_{i}__')
    text = text.replace('go-gpui', 'ownframe').replace('Ownframe', 'ownframe')
    text = text.replace('go_gpui', 'ownframe').replace('gogpui', 'ownframe')
    text = text.replace('GOGPUI', 'OWNFRAME').replace('GPUI_', 'OWNFRAME_').replace('OF_', 'OWNFRAME_')
    text = text.replace('gpui_test', 'ownframe_test')
    text = text.replace('data-gpui-', 'data-ownframe-').replace('data-of-', 'data-ownframe-')
    text = text.replace('`gpui`', '`ownframe`') if name != 'documentation/compare-rust-gpui.md' else text
    if path.suffix in {'.go', '.c', '.h'}:
        text = re.sub(r'(?m)^package (?:gpui|of)$', 'package ownframe', text)
        text = re.sub(r'(?m)^package of_test$', 'package ownframe_test', text)
        text = re.sub(r'\bgpui\b', 'ownframe', text)
        text = re.sub(r'\b(?:gpui|of)(?=[A-Z])', 'ownframe', text)
    else:
        text = re.sub(r'(?m)^package (?:gpui|of)$', 'package ownframe', text)
    text = re.sub(r'\b(?:gpui|of)\.(?=[A-Z])', 'ownframe.', text)
    text = re.sub(r'\b(?:gpui|of)\s+(?="github\.com/chinmay-sawant/ownframe")', '', text)
    text = re.sub(r'(\. |so )of\b', r'\1ownframe', text)
    text = re.sub(r'\bpackage of\b', 'package ownframe', text)
    text = text.replace('// of opens', '// ownframe opens').replace('the of page', 'the ownframe page')
    for prefix in ['of:', 'of print:', 'of-print-', 'of-cat-']:
        text = text.replace(prefix, prefix.replace('of', 'ownframe', 1))
    for i, value in enumerate(protected):
        text = text.replace(f'__REBRAND_KEEP_{i}__', value)
    if text != old:
        changed.append(name)
        if args.apply:
            path.write_text(text)
            if path.suffix == '.go':
                subprocess.run(['gofmt', '-w', str(path)], check=True)
print(f'{"Updated" if args.apply else "Would update"} {len(changed)} files.')
for name in changed:
    print(name)
if args.check and changed:
    raise SystemExit(1)
