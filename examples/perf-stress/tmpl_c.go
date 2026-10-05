package main

// tmplC holds the data-grid window: spacers above and below keep the full
// laid-out height, so only visible rows become boxes and operations.
const tmplC = `<div id="grid"><div id="padtop" style="height:{{.TopPad}}px"></div>{{range .Rows}}<div class="row" id="row-{{.Seq}}"><span class="num">{{.Num}}</span><span class="t">{{.Title}}</span><span class="b">{{.Body}}</span><span class="st">{{.Status}}</span><span class="d">{{.Delta}}</span></div>{{end}}<div id="padbot" style="height:{{.BotPad}}px"></div></div>
</div></div>
</body></html>
`

// buildHTML assembles the dashboard page from the template parts.
func buildHTML() string { return tmplA + tmplB + tmplC }
