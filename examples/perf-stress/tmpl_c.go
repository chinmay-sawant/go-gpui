package main

// tmplC holds the 280-row data grid and closes the document.
const tmplC = `<div id="grid">{{range .Rows}}<div class="row"><span class="num">{{.Num}}</span><span class="t">{{.Title}}</span><span class="b">{{.Body}}</span><span class="st">{{.Status}}</span><span class="d">{{.Delta}}</span></div>{{end}}</div>
</div></div>
</body></html>
`

// buildHTML assembles the dashboard page from the template parts.
func buildHTML() string { return tmplA + tmplB + tmplC }
