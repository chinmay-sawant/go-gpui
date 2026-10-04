package player

import "testing"

func TestLibraryScreenNavShowsGrid(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)

	click(t, app, "nav-library")

	if v := app.View(); v.Nav != "library" {
		t.Fatalf("Nav = %q", v.Nav)
	}

	redraw(t, app)

	ids := []string{"library-filter-0", "library-filter-1", "library-filter-4", "library-card-0"}
	for _, id := range ids {
		if _, ok := findBox(app.Boxes(), id); !ok {
			t.Fatalf("no box %s", id)
		}
	}
}

func TestLibraryScreenFilterNarrowsCards(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-library")
	redraw(t, app)

	click(t, app, "library-filter-1")

	v := app.View()
	if v.Library.Filter != 1 {
		t.Fatalf("Filter = %d", v.Library.Filter)
	}

	if len(v.Library.Cards) == 0 || len(v.Library.Cards) == len(v.Library.All) {
		t.Fatalf("cards = %d", len(v.Library.Cards))
	}

	for i, c := range v.Library.Cards {
		if c.Kind != 1 || c.Index != i || c.Active {
			t.Fatalf("card %d = %+v", i, c)
		}
	}

	redraw(t, app)

	if _, ok := findBox(app.Boxes(), "library-card-3"); !ok {
		t.Fatal("filtered card missing")
	}

	if _, ok := findBox(app.Boxes(), "library-card-9"); ok {
		t.Fatal("stale card drew")
	}
}

func TestLibraryScreenCardActivates(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-library")
	redraw(t, app)

	click(t, app, "library-card-1")

	v := app.View().Library
	if !v.Cards[1].Active {
		t.Fatalf("card 1 = %+v", v.Cards[1])
	}

	for i, c := range v.Cards {
		if i != 1 && c.Active {
			t.Fatalf("card %d also active", i)
		}
	}

	redraw(t, app)
}
