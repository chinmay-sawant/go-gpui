package game

import "testing"

func TestReplayReproducesTheSameResult(t *testing.T) {
	events := make([]InputEvent, 0, 30)
	for i := 0; i < 30; i++ {
		events = append(events, InputEvent{Step: uint64(i), Action: ActionHardDrop})
	}

	r := Replay{
		Seed:           42,
		Ruleset:        Ruleset,
		FixtureVersion: FixtureVersion,
		Events:         events,
	}

	first, err := r.Play(200)
	if err != nil {
		t.Fatal(err)
	}

	second, err := r.Play(200)
	if err != nil {
		t.Fatal(err)
	}

	if first.Score != second.Score || first.Lines != second.Lines ||
		first.Pieces != second.Pieces || first.Level != second.Level {
		t.Fatalf("replays differ: %+v vs %+v", first, second)
	}

	if first.Pieces == 0 {
		t.Fatal("the replay did nothing")
	}
}

func TestReplayFromAFixture(t *testing.T) {
	r := Replay{
		Seed:           1,
		Ruleset:        Ruleset,
		FixtureVersion: FixtureVersion,
		Fixture:        "clear-4",
		Events:         []InputEvent{{Step: 0, Action: ActionHardDrop}},
	}

	res, err := r.Play(100)
	if err != nil {
		t.Fatal(err)
	}

	if res.Lines != 4 || res.Score != LineScores[4] {
		t.Fatalf("fixture replay scored %d lines, %d points", res.Lines, res.Score)
	}
}

func TestReplayValidation(t *testing.T) {
	bad := []Replay{
		{Ruleset: "other"},
		{FixtureVersion: 99},
		{Fixture: "nope"},
		{Events: []InputEvent{
			{Step: 5, Action: ActionHardDrop},
			{Step: 4, Action: ActionLeft},
		}},
		{Events: []InputEvent{{Action: ActionNone}}},
	}

	for i, r := range bad {
		if err := r.Validate(); err == nil {
			t.Fatalf("bad replay %d passed validation", i)
		}

		if _, err := r.Play(10); err == nil {
			t.Fatalf("bad replay %d played", i)
		}
	}
}
