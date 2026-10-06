package scene

import "testing"

func TestHistoryRowsPadOnePage(t *testing.T) {
	var h history
	h.apply(ScorePage{Index: 0, Total: 2, Entries: []ScoreEntry{
		{Rank: 1, Score: 9},
		{Rank: 2, Score: 5, Dummy: true},
	}})

	rows := h.rows()
	if len(rows) != PageSize {
		t.Fatalf("rows = %d, want %d", len(rows), PageSize)
	}

	if rows[0].Text == "" || rows[1].Text == "" {
		t.Fatal("loaded rows are empty")
	}

	for i := 2; i < PageSize; i++ {
		if rows[i].Text != "" {
			t.Fatalf("row %d = %q, want empty", i, rows[i].Text)
		}
	}

	if rows[1].Top <= rows[0].Top {
		t.Fatal("rows are not stacked downward")
	}
}

func TestHistoryEmptyMessages(t *testing.T) {
	var h history
	h.apply(ScorePage{Index: 0})

	if h.msg != "NO SCORES YET" {
		t.Fatalf("live msg = %q", h.msg)
	}

	h.demo = true
	h.apply(ScorePage{Index: 0})

	if h.msg != "NO DEMO SCORES" {
		t.Fatalf("demo msg = %q", h.msg)
	}

	if !h.canPrev() && h.page != 0 {
		t.Fatal("a fresh page has a previous page")
	}

	if h.canNext() {
		t.Fatal("an empty page has a next page")
	}
}

func TestHistoryBoundaries(t *testing.T) {
	var h history
	h.apply(ScorePage{Index: 1, More: true, Entries: []ScoreEntry{{Rank: 21}}})

	if !h.canPrev() || !h.canNext() {
		t.Fatal("a middle page must page both ways")
	}

	if got := h.want(-3); got != 0 {
		t.Fatalf("want(-3) = %d", got)
	}

	h.apply(ScorePage{Index: 2, More: false})

	if h.canNext() {
		t.Fatal("the last page has a next page")
	}
}

func TestHistoryModes(t *testing.T) {
	var h history
	if h.title() != "HIGH SCORES" || h.demoToggle() != "DEMO" {
		t.Fatal("live captions are wrong")
	}

	h.demo = true
	if h.title() != "DEMO SCORES" || h.demoToggle() != "LIVE" {
		t.Fatal("demo captions are wrong")
	}
}
