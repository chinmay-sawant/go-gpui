package player

import "testing"

func TestBrowseScreenNavAndIDs(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)

	click(t, app, "nav-browse")

	if app.View().Nav != "browse" {
		t.Fatalf("Nav = %q", app.View().Nav)
	}

	redraw(t, app)

	ids := []string{"browse-hero-0", "browse-hero-2", "browse-tile-0", "browse-tile-11"}

	for _, id := range ids {
		if _, ok := findBox(app.Boxes(), id); !ok {
			t.Fatalf("no box %s", id)
		}
	}
}

func TestBrowseScreenTileSelect(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-browse")
	click(t, app, "browse-tile-4")

	v := app.View()
	if !v.Browse.Tiles[4].Active {
		t.Fatal("tile 4 not active")
	}

	for i, tile := range v.Browse.Tiles {
		if i != 4 && tile.Active {
			t.Fatalf("tile %d is active too", i)
		}
	}
}

func TestBrowseScreenHeroSelect(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-browse")
	click(t, app, "browse-hero-1")

	v := app.View()
	if !v.Browse.Hero[1].Active {
		t.Fatal("hero 1 not active")
	}

	for i, hero := range v.Browse.Hero {
		if i != 1 && hero.Active {
			t.Fatalf("hero %d is active too", i)
		}
	}
}
