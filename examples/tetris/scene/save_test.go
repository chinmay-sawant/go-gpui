package scene

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestCompletedGameIsSaved(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	res := game.Result{ID: "g1", Score: 10}
	m.done = &res
	m.rep = &game.Replay{Seed: 7}

	tickAt(t, s, clock, time.Second/60)

	if len(st.saves) != 1 || st.saves[0].ID != "g1" {
		t.Fatalf("saves = %v", st.saves)
	}

	if len(st.replays) != 1 || st.replays[0] == nil {
		t.Fatal("the replay was not passed to the store")
	}

	if st.clears != 1 {
		t.Fatalf("resume clears = %d, want 1", st.clears)
	}

	st.push(Result{ID: st.saveID, Kind: ResultSaved})
	tickAt(t, s, clock, time.Second/60)

	if s.pendingSave != nil {
		t.Fatal("the completed game is still pending")
	}

	if s.status != "SCORE SAVED" {
		t.Fatalf("status = %q", s.status)
	}
}

func TestRefusedSaveRetries(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	res := game.Result{ID: "g2", Score: 4}
	m.done = &res
	st.refuse = 1

	tickAt(t, s, clock, time.Second/60)

	if len(st.saves) != 0 {
		t.Fatal("a refused save was recorded")
	}

	if s.pendingSave == nil {
		t.Fatal("the refused save was dropped")
	}

	tickAt(t, s, clock, time.Second)

	if len(st.saves) != 0 {
		t.Fatal("the save retried inside the backoff")
	}

	tickAt(t, s, clock, 3*time.Second)

	if len(st.saves) != 1 {
		t.Fatalf("saves after the retry = %d, want 1", len(st.saves))
	}
}
