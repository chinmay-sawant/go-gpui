package scene

import (
	"fmt"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
	"github.com/chinmay-sawant/ownframe/examples/tetris/store"
)

func TestPageOfSlicesAndRanks(t *testing.T) {
	all := make([]game.Result, 45)
	for i := range all {
		all[i] = game.Result{ID: fmt.Sprint(i), Score: 100 - i}
	}

	p := pageOf(all, 1, PageSize*2)
	if p.Index != 1 || len(p.Entries) != PageSize || !p.More {
		t.Fatalf("page 1 = %+v", p)
	}

	if p.Entries[0].Rank != PageSize+1 {
		t.Fatalf("first rank = %d", p.Entries[0].Rank)
	}

	if p.Total != 45 {
		t.Fatalf("total = %d", p.Total)
	}

	p2 := pageOf(all, 2, PageSize*3)
	if len(p2.Entries) != 5 || p2.More {
		t.Fatalf("page 2 = %+v", p2)
	}

	p3 := pageOf(all, 9, PageSize*10)
	if len(p3.Entries) != 0 || p3.More {
		t.Fatalf("past-the-end page = %+v", p3)
	}
}

func TestSceneSettingsReadsTheTheme(t *testing.T) {
	dark := sceneSettings(store.Settings{Theme: themeDark})
	if !dark.Dark || dark.Control == nil {
		t.Fatalf("dark settings = %+v", dark)
	}

	light := sceneSettings(store.Settings{})
	if light.Dark {
		t.Fatal("the empty theme is not light")
	}

	if themeName(false) != "light" || themeName(true) != themeDark {
		t.Fatal("theme names are wrong")
	}
}
