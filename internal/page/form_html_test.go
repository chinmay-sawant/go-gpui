package page

import (
	"strings"
	"testing"
)

func whole(src string, ctrl Control) controlSpan {
	return controlSpan{Control: ctrl, Start: 0, End: len(src)}
}

func openOf(html, tag string) string {
	i := strings.Index(strings.ToLower(html), "<"+tag)
	if i < 0 {
		return ""
	}
	j := strings.Index(html[i:], ">")
	if j < 0 {
		return ""
	}

	return html[i : i+j]
}

func TestRewritePassword(t *testing.T) {
	t.Parallel()

	src := `<input type="password" id="pw" value="secret">`
	ctrl := Control{ID: "pw", Tag: "input", Type: "password", Value: "secret"}
	live := map[string]Control{
		"pw": {Tag: "input", Type: "password", Value: "nope"},
	}
	got := rewriteControls(src, []controlSpan{whole(src, ctrl)}, live, "")
	if !strings.HasPrefix(got, "<style>") || !strings.Contains(got, `<span id="pw"`) {
		t.Fatalf("%s", got)
	}
	if !strings.Contains(got, "••••") || strings.Contains(got, "secret") || strings.Contains(got, "nope") {
		t.Fatalf("%s", got)
	}
	if !strings.Contains(got, plainStyle) {
		t.Fatalf("%s", got)
	}
}
