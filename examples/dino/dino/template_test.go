package dino

import (
	"strings"
	"testing"
)

// TestTemplateInlinesTheStylesheet guards the marker buildHTML replaces.
func TestTemplateInlinesTheStylesheet(t *testing.T) {
	if !strings.Contains(indexHTML, "<!-- stylesheet -->") {
		t.Fatal("template/index.html lost its stylesheet marker")
	}

	if !strings.Contains(buildHTML(), "<style>") {
		t.Fatal("the stylesheet was not inlined")
	}
}
