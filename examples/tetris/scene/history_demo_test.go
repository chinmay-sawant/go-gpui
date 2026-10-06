package scene

import (
	"testing"
	"time"
)

func TestHistoryPagesForward(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	click(t, s, "btn-scores")

	st.push(Result{ID: st.scoresID, Kind: ResultScores, Scores: ScorePage{
		Index: 0, More: true, Entries: []ScoreEntry{{Rank: 1, Score: 5}},
	}})
	tickAt(t, s, clock, time.Second/60)

	click(t, s, "btn-next")

	if len(st.scoresCalls) != 2 || st.scoresCalls[1] != 1 {
		t.Fatalf("score requests = %v", st.scoresCalls)
	}

	if !s.hist.loading {
		t.Fatal("the next page is not marked loading")
	}
}

func TestHistoryDemoToggle(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	click(t, s, "btn-scores")
	click(t, s, "btn-demo")

	if !s.hist.demo {
		t.Fatal("the demo mode did not turn on")
	}

	if st.demoID == 0 {
		t.Fatal("no demo request was made")
	}

	st.push(Result{ID: st.demoID, Kind: ResultDummy, Scores: ScorePage{
		Entries: []ScoreEntry{{Rank: 1, Score: 9, Dummy: true}},
	}})
	tickAt(t, s, clock, time.Second/60)

	if !hasText(s, "DEMO SCORES") || !hasText(s, "DEMO") {
		t.Fatal("the demo page is not labelled")
	}
}
