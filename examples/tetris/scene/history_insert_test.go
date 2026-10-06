package scene

import (
	"testing"
	"time"
)

// TestHistoryRefetchesAfterAnInsert pages forward, then back, after a new
// top score landed between the two pages.
func TestHistoryRefetchesAfterAnInsert(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	click(t, s, "btn-scores")

	st.push(Result{ID: st.scoresID, Kind: ResultScores, Scores: ScorePage{
		Index: 0, More: true, Entries: []ScoreEntry{{Rank: 1, Score: 100}},
	}})
	tickAt(t, s, clock, time.Second/60)

	click(t, s, "btn-next")

	st.push(Result{ID: st.scoresID, Kind: ResultScores, Scores: ScorePage{
		Index: 1, Entries: []ScoreEntry{{Rank: 21, Score: 5}},
	}})
	tickAt(t, s, clock, time.Second/60)

	click(t, s, "btn-prev")

	if len(st.scoresCalls) != 3 || st.scoresCalls[2] != 0 {
		t.Fatalf("score requests = %v", st.scoresCalls)
	}

	st.push(Result{ID: st.scoresID, Kind: ResultScores, Scores: ScorePage{
		Index: 0, More: true, Entries: []ScoreEntry{{Rank: 1, Score: 200}},
	}})
	tickAt(t, s, clock, time.Second/60)

	if !hasText(s, "200") {
		t.Fatal("the refetched page kept the stale score")
	}
}
