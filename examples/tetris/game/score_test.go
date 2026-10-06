package game

import "testing"

func TestLineScoresAndLevels(t *testing.T) {
	if LineScores[1] != 100 || LineScores[2] != 300 ||
		LineScores[3] != 500 || LineScores[4] != 800 {
		t.Fatalf("line table is %v", LineScores)
	}

	if got := lineScore(0); got != 0 {
		t.Fatalf("zero rows score %d", got)
	}

	if got := lineScore(9); got != 800 {
		t.Fatalf("the table is not capped: %d", got)
	}

	cases := map[int]int{0: 1, 9: 1, 10: 2, 19: 2, 20: 3, MaxLines: MaxLevel}
	for lines, want := range cases {
		if got := levelFor(lines); got != want {
			t.Fatalf("levelFor(%d) = %d, want %d", lines, got, want)
		}
	}
}

func TestAddScoreCapsAtMaxScore(t *testing.T) {
	if got := addScore(MaxScore-5, 10); got != MaxScore {
		t.Fatalf("score overflowed to %d", got)
	}

	if got := addScore(10, 0); got != 10 {
		t.Fatalf("zero points changed the score to %d", got)
	}
}

func TestClearScoresAtTheLevelAndLevelsUp(t *testing.T) {
	g, err := NewFromFixture("clear-1", 1)
	if err != nil {
		t.Fatal(err)
	}

	g.Lines = 9

	ev := g.Apply(ActionHardDrop)

	if g.Score != LineScores[1] {
		t.Fatalf("score = %d, want %d", g.Score, LineScores[1])
	}

	if g.Level != 2 {
		t.Fatalf("level = %d, want 2", g.Level)
	}

	if !haveEvent(ev, EventLevelUp) {
		t.Fatal("crossing ten lines reported no level up")
	}
}

func TestScoreAndLinesStayBounded(t *testing.T) {
	g, err := NewFromFixture("clear-1", 1)
	if err != nil {
		t.Fatal(err)
	}

	g.Score = MaxScore
	g.Lines = MaxLines

	g.Apply(ActionHardDrop)

	if g.Score != MaxScore {
		t.Fatalf("score left its bound: %d", g.Score)
	}

	if g.Lines != MaxLines {
		t.Fatalf("lines left their bound: %d", g.Lines)
	}

	if g.Level != MaxLevel {
		t.Fatalf("level = %d, want %d", g.Level, MaxLevel)
	}
}
