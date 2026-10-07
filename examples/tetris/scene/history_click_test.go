package scene

import (
	"strings"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestHistoryOpenClosePreservesGame(t *testing.T) {
	m := &fakeModel{f: Frame{Phase: game.PhaseRunning, Score: 77}}
	st := &fakeStore{}
	s, _ := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	before := m.f
	click(t, s, "btn-scores")

	if !s.hist.open {
		t.Fatal("history did not open")
	}

	if m.pauses != 1 || m.clears != 1 {
		t.Fatalf("pauses = %d clears = %d", m.pauses, m.clears)
	}

	if !strings.Contains(s.page.HTML(), "btn-prev") {
		t.Fatal("the history template is not loaded")
	}

	if len(st.scoresCalls) != 1 || st.scoresCalls[0] != 0 {
		t.Fatalf("score requests = %v", st.scoresCalls)
	}

	click(t, s, "btn-close")

	if s.hist.open {
		t.Fatal("history did not close")
	}

	if m.resumes != 1 {
		t.Fatalf("resumes = %d, want 1", m.resumes)
	}

	if m.f != before {
		t.Fatal("the board state changed while history was open")
	}

	if !strings.Contains(s.page.HTML(), "btn-pause") {
		t.Fatal("the game template is not back")
	}
}

func TestHistoryShowsALoadedPage(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	click(t, s, "btn-scores")

	st.push(Result{ID: st.scoresID, Kind: ResultScores, Scores: ScorePage{
		Index: 0, Total: 1, Entries: []ScoreEntry{{Rank: 1, ID: "g-1", Score: 1234, Lines: 3, Level: 2}},
	}})
	tickAt(t, s, clock, time.Second/60)

	if !hasText(s, "1234") || !hasText(s, "g-1") {
		t.Fatal("the loaded score row is missing")
	}
}
