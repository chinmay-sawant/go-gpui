package player

import "testing"

func TestQueueScreenNav(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)

	click(t, app, "nav-queue")

	if app.View().Nav != "queue" {
		t.Fatalf("nav = %q", app.View().Nav)
	}

	redraw(t, app)

	ids := []string{"queue-row-0", "queue-row-5", "queue-from-0", "queue-from-1"}

	for _, id := range ids {
		if _, ok := findBox(app.Boxes(), id); !ok {
			t.Fatalf("no box %s", id)
		}
	}
}

func TestQueueScreenRow(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-queue")
	redraw(t, app)

	want := app.View().Queue.Next[2].Title

	click(t, app, "queue-row-2")
	waitAudio(t, app)

	v := app.View()
	if v.Now.Title != want {
		t.Fatalf("now = %q, want %q", v.Now.Title, want)
	}

	if !v.Queue.Next[2].Active {
		t.Fatal("row 2 not active")
	}

	for i, tr := range v.Queue.Next {
		if i != 2 && tr.Active {
			t.Fatalf("row %d still active", i)
		}
	}

	redraw(t, app)
}

func TestQueueScreenFrom(t *testing.T) {
	app := mustApp(t)
	redraw(t, app)
	click(t, app, "nav-queue")
	redraw(t, app)

	want := app.View().Shelf[1]

	click(t, app, "queue-from-1")
	waitAudio(t, app)

	v := app.View()
	if v.Now.Title != want.Title || v.Now.Cover != want.Cover {
		t.Fatalf("now = %+v, want %+v", v.Now, want)
	}

	if !v.Shelf[1].Active {
		t.Fatal("shelf 1 not active")
	}

	for i, tr := range v.Queue.Next {
		if tr.Active {
			t.Fatalf("queue row %d disturbed", i)
		}
	}

	redraw(t, app)
}
