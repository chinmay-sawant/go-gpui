package transfer

import (
	"strings"
	"testing"
)

// TestSafeName covers the untrusted-path cases.
func TestSafeName(t *testing.T) {
	cases := map[string]string{
		"https://host/path/report.pdf":       "report.pdf",
		"https://host/a/b/":                  "b",
		"":                                   "download",
		"https://host/":                      "host",
		"https://host/CON":                   "_CON",
		"https://host/con.txt":               "_con.txt",
		"https://host/nul":                   "_nul",
		"https://host/lpt9.log":              "_lpt9.log",
		"https://host/..":                    "download",
		"https://host/.":                     "download",
		"https://host/a:b*c?.bin":            "a_b_c",
		"https://host/héllo wörld.txt":       "héllo wörld.txt",
		"https://host/name.":                 "name",
		"https://host/name ":                 "name",
		"https://host/back\\slash.txt":       "slash.txt",
		"https://host/ctrl\x01char":          "ctrlchar",
		"https://host/file.bin?token=secret": "file.bin",
		"https://host/file.bin#frag":         "file.bin",
	}

	for raw, want := range cases {
		if got := SafeName(raw); got != want {
			t.Errorf("SafeName(%q) = %q, want %q", raw, got, want)
		}
	}

	long := strings.Repeat("x", 300) + ".bin"
	got := SafeName(long)
	if len(got) > 120 || !strings.HasSuffix(got, ".bin") {
		t.Errorf("long name not clipped: %d %q", len(got), got)
	}
}
