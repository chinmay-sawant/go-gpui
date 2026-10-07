package storage

import (
	"testing"
	"time"
)

// TestViewsRoundTrip checks save, replace, order, and delete.
func TestViewsRoundTrip(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	base := time.Now()

	first := View{ID: "v1", Name: "Top CPU", Query: "chrome", Sort: "cpu", Desc: true, Updated: base}
	if err := st.PutView(ctx, first); err != nil {
		t.Fatal(err)
	}
	if err := st.PutView(ctx, View{ID: "v2", Name: "Memory", Sort: "mem", Updated: base.Add(time.Second)}); err != nil {
		t.Fatal(err)
	}

	views, err := st.Views(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(views) != 2 || views[0].ID != "v2" || !views[1].Desc {
		t.Fatalf("views = %+v", views)
	}

	first.Name = "Top CPU (edited)"
	first.Updated = base.Add(2 * time.Second)
	if err := st.PutView(ctx, first); err != nil {
		t.Fatal(err)
	}

	views, err = st.Views(ctx)
	if err != nil || len(views) != 2 || views[0].Name != "Top CPU (edited)" {
		t.Fatalf("views = %+v err=%v", views, err)
	}

	if err := st.DeleteView(ctx, "v1"); err != nil {
		t.Fatal(err)
	}
	if views, err = st.Views(ctx); err != nil || len(views) != 1 {
		t.Fatalf("views = %+v err=%v", views, err)
	}

	if err := st.PutView(ctx, View{ID: "", Name: "x"}); err == nil {
		t.Fatal("view without an id accepted")
	}
}
