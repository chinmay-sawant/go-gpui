#!/usr/bin/env python3
"""Summarize repeated dumps and check the window's geometry."""
import gzip
import json
import statistics as st
from pathlib import Path

base = Path(__file__).resolve().parents[4] / 'temp' / 'complex-dump'
summary = {}
for mode in ['initial', 'cached', 'data', 'resize', 'windowed']:
    runs = [json.loads((base / f'run-{i}' / f'{mode}.json').read_text()) for i in range(1, 4)]
    # Initial sample zero carries process-first font work; keep it separate.
    samples = [r['Samples'][1:] if mode == 'initial' else r['Samples'] for r in runs]
    assert all(s['Stats']['PaintTime'] == 0 for ss in samples for s in ss)
    medians = [st.median(s['ElapsedNS'] / 1e6 for s in ss) for ss in samples]
    mid = sorted(range(3), key=lambda i: medians[i])[1]
    chosen = sorted(samples[mid], key=lambda s: s['ElapsedNS'])[len(samples[mid]) // 2]
    summary[mode] = {
        'run_medians_ms': medians, 'median_ms': st.median(medians),
        'alloc_MB': chosen['AllocBytes'] / 1e6, 'allocs': chosen['Allocs'],
        'stats': chosen['Stats'],
        'stage_medians_ms': {k: st.median(s['Stats'][k] / 1e6 for ss in samples for s in ss)
                             for k in ['LastTemplate', 'LayoutTime', 'DisplayListTime', 'PaintTime']},
    }
    if mode == 'initial':
        summary[mode]['process_first_ms'] = [r['Samples'][0]['ElapsedNS'] / 1e6 for r in runs]

checks = []
for i in range(1, 4):
    def layout(mode):
        with gzip.open(base / f'run-{i}' / f'{mode}-layout.json.gz') as f:
            return json.load(f)
    full, window = layout('cached'), layout('windowed')
    def named(d):
        return {b['ID']: b for b in d['boxes'] if b['ID']}
    a, b = named(full), named(window)
    assert (full['width'], full['height']) == (window['width'], window['height'])
    for key in ['grid'] + [f'row-{n}' for n in range(48)]:
        assert all(a[key][k] == b[key][k] for k in ['X', 'Y', 'W', 'H']), key
    assert a['grid']['H'] == 480 * 76
    assert all(a[f'row-{n}']['H'] == 76 for n in range(480))
    assert len(full['ops']) > 10000 and len(full['boxes']) > 9000
    assert len(window['ops']) < 1500
    # Ops use points. Compare all operations intersecting the 1000px viewport.
    def visible(d):
        return [o for o in d['ops'] if o['Y'] < 750 and o['Y'] + o['H'] >= 0]
    assert visible(full) == visible(window), 'visible operation geometry/text differs'
    checks.append({'run': i, 'canvas': [full['width'], full['height']],
                   'grid_height_px': a['grid']['H'], 'visible_ops': len(visible(full)),
                   'first_48_row_boxes_equal': True, 'visible_ops_equal': True})
summary['geometry_checks'] = checks
summary['windowing'] = {
    'time_reduction_percent': 100 * (1 - summary['windowed']['median_ms'] / summary['cached']['median_ms']),
    'allocation_reduction_percent': 100 * (1 - summary['windowed']['alloc_MB'] / summary['cached']['alloc_MB']),
    'speedup': summary['cached']['median_ms'] / summary['windowed']['median_ms'],
}
(base / 'summary.json').write_text(json.dumps(summary, indent=2) + '\n')
for k, v in summary.items():
    if 'median_ms' in v:
        print(k, round(v['median_ms'], 2), 'ms', round(v['alloc_MB'], 2), 'MB', v['stats']['Ops'], 'ops')
print(summary['windowing'])
print('Geometry checks passed for all three runs.')
