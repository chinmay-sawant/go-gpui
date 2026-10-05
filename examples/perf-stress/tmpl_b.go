package main

// tmplB holds the header, sidebar, cards, progress, equalizer, and form.
const tmplB = `<div id="top"><span id="brand">🚀 Stress {{.Title}}</span><span id="tick">{{.Tick}}</span><span>❤️ live</span></div>
<div id="wrap"><div id="side">{{range .Nav}}<div class="nav" id="nav-{{.}}">{{.}}</div>{{end}}</div>
<div id="main"><div>Active: {{.Active}} ⚙️</div>
<div id="cards">{{range .Cards}}<div class="card"><div class="cl">{{.Label}}</div><div class="cv">{{.Value}}</div><div class="ch">{{.Hint}}</div></div>{{end}}</div>
<div id="prow"><div id="ptrack"><div id="seek"></div></div><span id="pct">{{.Progress}}%</span></div>
<div id="eq"><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div><div class="eqbar"></div></div>
<div id="form"><input id="q" type="text" value="filter orders..."><input id="bell" type="checkbox" checked><select id="sort"><option>newest</option><option>oldest</option></select><button id="go">Apply ✅</button><button id="reset">Reset</button></div>
`
