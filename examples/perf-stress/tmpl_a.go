package main

// tmplA opens the document with the dashboard stylesheet.
const tmplA = `<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>Stress</title><style>
html,body{margin:0}
body{background:#14161c;font-family:sans-serif;color:#e8e6df;font-size:13px}
#top{background:#1f2330;padding:10px 16px;display:flex;gap:12px}
#brand{font-size:16px;font-weight:bold;color:#fff}
#tick{color:#9aa3b2}
#wrap{display:flex}
#side{width:150px;background:#1b1e29;padding:8px}
.nav{padding:7px 10px;margin:3px 0;background:#262b3b;color:#cfd4e0}
#main{padding:12px}
#cards{display:flex;gap:10px}
.card{width:150px;background:#222736;padding:10px;border:1px solid #343a4f}
.cv{font-size:20px;color:#fff}
.ch{color:#9aa3b2;font-size:12px}
#prow{display:flex;gap:10px;margin:12px 0}
#ptrack{width:420px;height:14px;background:#343a4f}
#seek{width:60px;height:14px;background:#1db954}
#pct{color:#9aa3b2}
#eq{display:flex;gap:4px;height:56px;margin:8px 0}
.eqbar{width:18px;height:28px;background:#1db954}
#form{display:flex;gap:8px;margin:10px 0}
.row{display:flex;gap:10px;height:36px;box-sizing:border-box;padding:5px 8px;border-bottom:1px solid #2b3042}
.num{width:52px;color:#9aa3b2}
.t{width:130px;font-weight:bold;color:#fff}
.st{width:110px}
.d{width:60px;color:#7ee2a8}
</style></head><body>
`
