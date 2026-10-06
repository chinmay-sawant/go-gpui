package page_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

func formPage(t *testing.T, body string) *page.Page {
	t.Helper()

	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:8px;font-family:sans-serif">` + body + `</body>`,
		Width:  640,
		Height: 480,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err = screen.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return screen
}

func clickID(t *testing.T, screen *page.Page, id string) {
	t.Helper()

	x, y := boxCenter(t, screen, id)
	if err := screen.Click(context.Background(), x, y); err != nil {
		t.Fatal(err)
	}
}

const wide = ` style="display:block;width:220px;height:28px"`

func TestFormTypeText(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen := formPage(t, `<input id="t" type="text"`+wide+`>`)
	clickID(t, screen, "t")
	if err := screen.Type(ctx, "ab"); err != nil {
		t.Fatal(err)
	}

	if screen.FormValue("t") != "ab" || boxText(t, screen, "t") != "ab" {
		t.Fatalf("value %q text %q", screen.FormValue("t"), boxText(t, screen, "t"))
	}
}

func TestFormPasswordCopy(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen := formPage(t, `<input id="pw" type="password"`+wide+`>`)
	clickID(t, screen, "pw")
	if err := screen.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	if screen.FormValue("pw") != "secret" {
		t.Fatalf("value %q", screen.FormValue("pw"))
	}

	got := boxText(t, screen, "pw")
	if strings.Contains(got, "secret") || !maskText(got) {
		t.Fatalf("box %q", got)
	}

	text, ok, err := screen.Copy(ctx)
	if err != nil || !ok || text != "secret" {
		t.Fatalf("copy %q %v %v", text, ok, err)
	}
}

func maskText(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		switch r {
		case '•', '*', '●', '·', '∗', '⬤':
		default:
			return false
		}
	}

	return true
}
