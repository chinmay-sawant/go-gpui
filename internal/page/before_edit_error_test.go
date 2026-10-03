package page

import (
	"context"
	"errors"
	"testing"
)

func TestBeforeEditErrorAbortsType(t *testing.T) {
	ctx := context.Background()
	p := editPage(t, `<input id="e" type="text" style="display:block;width:120px;height:20px">`)
	clickBox(t, p, "e")

	gen := p.Generation()
	boom := errors.New("stop")
	p.Handle(Handlers{BeforeEdit: func(context.Context, Box) error { return boom }})

	if err := p.Type(ctx, "x"); !errors.Is(err, boom) {
		t.Fatalf("err = %v", err)
	}

	if p.FormValue("e") != "" {
		t.Fatalf("value = %q", p.FormValue("e"))
	}

	if p.Generation() != gen {
		t.Fatal("page redrew after abort")
	}
}

func TestBeforeEditFiresOnCheckbox(t *testing.T) {
	p := editPage(t, `<input id="c" type="checkbox">`)

	var got Box
	p.Handle(Handlers{BeforeEdit: func(_ context.Context, b Box) error {
		got = b
		return nil
	}})

	clickBox(t, p, "c")

	if got.ID != "c" || !p.FormChecked("c") {
		t.Fatalf("box = %+v checked = %v", got, p.FormChecked("c"))
	}
}
