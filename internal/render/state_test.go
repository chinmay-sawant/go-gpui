package render

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/blinkless/css"
)

func stateWidth(t *testing.T, source string, state State) float64 {
	t.Helper()

	display, err := DisplayListState(context.Background(), source, 320, 200, state)
	if err != nil {
		t.Fatal(err)
	}

	for _, b := range display.Boxes {
		if b.ID == "e" {
			return b.W
		}
	}

	t.Fatal("no box e")

	return 0
}

func TestDisplayListStatePseudoClasses(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>` +
		`#e{width:100px;height:10px}` +
		`#e:focus{width:120px}` +
		`#e:hover{width:130px}` +
		`#e:active{width:140px}` +
		`</style></head><body><div id="e"></div></body></html>`

	if got := stateWidth(t, source, State{}); got != 100 {
		t.Fatalf("plain width = %v", got)
	}

	if got := stateWidth(t, source, State{Focus: "e"}); got != 120 {
		t.Fatalf("focus width = %v", got)
	}

	if got := stateWidth(t, source, State{Hover: "e"}); got != 130 {
		t.Fatalf("hover width = %v", got)
	}

	if got := stateWidth(t, source, State{Active: "e"}); got != 140 {
		t.Fatalf("active width = %v", got)
	}
}

func TestDisplayListStateChecked(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>` +
		`input{width:20px;height:10px}` +
		`input:checked{width:40px}` +
		`</style></head><body><input id="e" type="checkbox" checked></body></html>`

	if got := stateWidth(t, source, State{}); got != 40 {
		t.Fatalf("checked width = %v", got)
	}
}

func TestDisplayListStateTheme(t *testing.T) {
	t.Parallel()

	sheet, err := css.Parse(`#e{width:150px}`)
	if err != nil {
		t.Fatal(err)
	}

	source := `<html><head><style>` +
		`#e{width:100px;height:10px}` +
		`</style></head><body><div id="e"></div></body></html>`

	if got := stateWidth(t, source, State{Theme: sheet}); got != 150 {
		t.Fatalf("theme width = %v, want 150", got)
	}
}
