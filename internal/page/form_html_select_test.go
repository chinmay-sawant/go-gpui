package page

import (
	"strings"
	"testing"
)

func TestRewriteSelect(t *testing.T) {
	t.Parallel()

	src := `<select id="s" name="s"><option value="a">Ant</option><option value="b">Bee</option></select>`
	ctrl := Control{
		ID: "s", Tag: "select", Name: "s", Value: "b",
		Options: []Option{{Value: "a", Label: "Ant"}, {Value: "b", Label: "Bee"}},
	}
	got := rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "", false)
	if !strings.Contains(got, ">Bee</select>") || strings.Contains(got, "Ant") || strings.Contains(got, "<option") {
		t.Fatalf("%s", got)
	}

	ctrl.Value = "nope"
	got = rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "", false)
	if !strings.Contains(got, ">Ant</select>") || strings.Contains(got, "Bee") {
		t.Fatalf("%s", got)
	}

	ctrl.Options = []Option{}
	ctrl.Value = "missing"
	got = rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "", false)
	if !strings.Contains(got, ">missing</select>") {
		t.Fatalf("%s", got)
	}
}
