package benchutil

// normalHTML is the Benchmark A page: a small desktop form with a counter,
// two buttons, one bound text input, and a greeting. Under 20 elements.
const normalHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Normal</title>
<style>
html,body{height:100%;margin:0}
body{background:#f4f1ea;font-family:sans-serif;color:#1c1915;display:flex;align-items:center;justify-content:center}
.card{width:320px;background:#fff;padding:24px}
h1{font-size:22px;margin:0 0 8px}
#count{font-size:48px;margin:8px 0}
#inc{display:block;background:#1a56db;color:#fff;padding:10px;text-align:center}
.field{display:block;box-sizing:border-box;width:100%;border:1px solid #c8c2b4;padding:8px;margin-top:4px}
#hello{font-size:14px;min-height:18px}
#reset{display:block;margin-top:12px;padding:8px;background:#e5e7eb;text-align:center}
</style>
</head>
<body>
<div class="card">
<h1>Normal</h1>
<p id="count">{{.Count}}</p>
<div id="inc">Click me</div>
<label>Name</label>
<input id="name" class="field" type="text" data-bind="Name">
<p id="hello">Hello, {{.Name}}</p>
<div id="reset">Reset</div>
</div>
</body>
</html>`
