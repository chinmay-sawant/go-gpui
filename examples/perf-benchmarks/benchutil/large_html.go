package benchutil

// largeHTML is the Benchmark B page: a scrolling text-heavy list with one
// range over .Rows. Each row is three elements, so 1000 rows is 3000.
const largeHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Large</title>
<style>
html,body{margin:0}
body{background:#f4f1ea;font-family:sans-serif;color:#1c1915}
#head{background:#1c1915;color:#fff;padding:12px 16px;font-size:18px}
.row{display:flex;gap:12px;padding:8px 16px;border-bottom:1px solid #ddd6c7}
.num{width:64px;color:#8a8578}
.title{width:200px;font-weight:bold}
.body{color:#4a463c}
</style>
</head>
<body>
<div id="head">Large: {{.Total}} rows</div>
{{range .Rows}}<div class="row"><span class="num">{{.Num}}</span><span class="title">{{.Title}}</span><span class="body">{{.Body}}</span></div>{{end}}
</body>
</html>`
