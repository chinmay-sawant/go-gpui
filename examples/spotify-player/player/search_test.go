package player

import "testing"

func TestSearchScreenNavAndBoxes(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-search")

	if v := app.View(); v.Nav != "search" {
		t.Fatalf("Nav = %q", v.Nav)
	}

	redraw(t, app)

	for _, id := range []string{"search-tile-0", "search-tile-11", "search-recent-0", "search-recent-4"} {
		if _, ok := findBox(app.Boxes(), id); !ok {
			t.Fatalf("no box %s", id)
		}
	}
}

func TestSearchScreenTileSelect(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-search")
	redraw(t, app)
	click(t, app, "search-tile-3")

	v := app.View()
	if !v.Search.Tiles[3].Active {
		t.Fatalf("tile 3 = %+v", v.Search.Tiles[3])
	}

	for i, tile := range v.Search.Tiles {
		if i != 3 && tile.Active {
			t.Fatalf("tile %d also active", i)
		}
	}

	redraw(t, app)
}

func TestSearchScreenRecentPlays(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-search")
	redraw(t, app)
	click(t, app, "search-recent-1")

	v := app.View()
	if v.Now.Title != v.Search.Recent[1].Title {
		t.Fatalf("Now = %+v", v.Now)
	}

	if v.Progress != 0 || !v.Playing {
		t.Fatalf("progress = %d playing = %v", v.Progress, v.Playing)
	}

	redraw(t, app)
}

func TestSearchScreenResults(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-search")

	app.view.Query = "nova"
	app.page.SetData(app.view)
	redraw(t, app)

	if _, ok := findBox(app.Boxes(), "search-song-0"); !ok {
		t.Fatal("no box search-song-0")
	}

	click(t, app, "search-song-2")

	v := app.View()
	if !v.Tracks[2].Active || v.Now.Title != v.Tracks[2].Title {
		t.Fatalf("Now = %+v", v.Now)
	}

	redraw(t, app)
}
