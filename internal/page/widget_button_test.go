package page

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

func TestRewriteButtonsSubmit(t *testing.T) {
	t.Parallel()

	src := `<input id="sub" class="go" type="submit" value="Send" data-action="send">`
	got := rewriteButtons(src)
	if !strings.Contains(got, `<button id="sub" class="go" data-action="send">Send</button>`) {
		t.Fatalf("%s", got)
	}
	if strings.Contains(got, "type=") || strings.Contains(got, "value=") {
		t.Fatalf("%s", got)
	}
}

func TestRewriteButtonsDefaults(t *testing.T) {
	t.Parallel()

	if got := rewriteButtons(`<input type="reset">`); !strings.Contains(got, ">Reset</button>") {
		t.Fatalf("%s", got)
	}
	if rewriteButtons("plain") != "plain" {
		t.Fatal("plain text changed")
	}
}

func TestButtonCSSInjected(t *testing.T) {
	t.Parallel()

	got := rewriteControls(`<button id="go">Sign in</button>`, nil, nil, "", false)
	if !strings.Contains(got, buttonCSS) || !strings.Contains(got, `<button id="go">Sign in</button>`) {
		t.Fatalf("%s", got)
	}
	if rewriteControls("hi", nil, nil, "a", false) != "hi" {
		t.Fatal("plain text changed")
	}
}

func TestButtonGetsPaintBox(t *testing.T) {
	t.Parallel()

	src := rewriteControls(`<button id="go" data-action="go">Sign in</button>`, nil, nil, "", false)
	_, boxes, err := render.Paint(context.Background(), src, 400, 200)
	if err != nil {
		t.Fatal(err)
	}
	for _, b := range boxes {
		if b.ID != "go" {
			continue
		}
		if b.W <= 0 || b.H <= 0 {
			t.Fatalf("button box %+v", b)
		}

		return
	}
	t.Fatal("no button box")
}
