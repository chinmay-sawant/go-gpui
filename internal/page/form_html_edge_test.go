package page

import (
	"strings"
	"testing"
)

func TestRewriteOverlap(t *testing.T) {
	t.Parallel()

	src := "0123456789"
	spans := []controlSpan{
		{Control: Control{ID: "c", Type: "text", Value: "CC"}, Start: 6, End: 10},
		{Control: Control{ID: "z", Type: "text", Value: "ZZ"}, Start: -1, End: 2},
		{Control: Control{ID: "a", Type: "text", Value: "AA"}, Start: 0, End: 4},
		{Control: Control{ID: "b", Type: "text", Value: "BB"}, Start: 2, End: 6},
		{Control: Control{ID: "y", Type: "text", Value: "YY"}, Start: 8, End: 20},
	}
	got := rewriteControls(src, spans, nil, "a")
	if strings.Contains(got, "BB") || strings.Contains(got, "ZZ") || strings.Contains(got, "YY") {
		t.Fatalf("%s", got)
	}
	if !strings.Contains(got, ">AA</span>45") || !strings.Contains(got, ">CC</span>") {
		t.Fatalf("%s", got)
	}
	if !strings.Contains(got, focusStyle) {
		t.Fatalf("%s", got)
	}
}

func TestRewriteHead(t *testing.T) {
	t.Parallel()

	src := "<html><HEAD><title>t</title></HEAD><body>X</body></html>"
	at := strings.Index(src, "X")
	sp := controlSpan{Control: Control{ID: "a", Type: "text", Value: "Z"}, Start: at, End: at + 1}
	got := rewriteControls(src, []controlSpan{sp}, nil, "")
	style := strings.Index(got, "<style>")
	head := strings.Index(strings.ToLower(got), "</head>")
	if style < 0 || head < 0 || style > head || !strings.Contains(got, ">Z</span>") {
		t.Fatalf("%s", got)
	}
	if rewriteControls("hi", nil, nil, "a") != "hi" {
		t.Fatal("empty spans changed")
	}
}

func TestRewriteFile(t *testing.T) {
	t.Parallel()

	src := `<input type="file" id="f">`
	ctrl := Control{ID: "f", Tag: "input", Type: "file"}
	got := rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "")
	if !strings.Contains(got, ">No file</span>") {
		t.Fatalf("%s", got)
	}
	ctrl.Value = "a.txt"
	got = rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "")
	if !strings.Contains(got, ">a.txt</span>") {
		t.Fatalf("%s", got)
	}
}
