package web

import (
	"bytes"
	"errors"
	"html/template"
	"strings"
	"testing"
)

func renderShell(t *testing.T, data shellData) string {
	t.Helper()

	tpl, err := template.New("shell").Parse(shellHTML)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		t.Fatal(err)
	}

	return buf.String()
}

func TestShellRefreshesFrameWhileWatching(t *testing.T) {
	t.Parallel()

	on := renderShell(t, shellData{Reload: true})
	if !strings.Contains(on, `id="frame"`) || !strings.Contains(on, "setInterval") {
		t.Fatalf("watching shell:\n%s", on)
	}

	off := renderShell(t, shellData{})
	if strings.Contains(off, "setInterval") {
		t.Fatalf("string shell has a script:\n%s", off)
	}
}

func TestReloadNotePrintsOncePerDistinctError(t *testing.T) {
	t.Parallel()

	s := &server{}
	err := errors.New("/tmp/index.html: stat: no such file")

	line, show := s.reloadNote(err)
	if !show || line != "hot reload: /tmp/index.html: stat: no such file" {
		t.Fatalf("line = %q, show = %v", line, show)
	}

	if _, show := s.reloadNote(err); show {
		t.Fatal("the same error printed twice")
	}
}
