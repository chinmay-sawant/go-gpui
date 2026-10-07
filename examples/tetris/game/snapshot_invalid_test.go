package game

import "testing"

func TestInvalidSnapshotsAreRejected(t *testing.T) {
	g := New(9)
	g.Start()

	good := g.Snapshot()

	cases := []struct {
		name   string
		mutate func(*Snapshot)
	}{
		{"id", func(s *Snapshot) { s.ID = "" }},
		{"ruleset", func(s *Snapshot) { s.Ruleset = "other" }},
		{"fixture", func(s *Snapshot) { s.FixtureVersion = 99 }},
		{"board", func(s *Snapshot) { s.Board = s.Board[:19] }},
		{"piece", func(s *Snapshot) { s.Piece = 0 }},
		{"position", func(s *Snapshot) { s.X = 99 }},
		{"score", func(s *Snapshot) { s.Score = -1 }},
		{"level", func(s *Snapshot) { s.Level = 0 }},
		{"timers", func(s *Snapshot) { s.LockMS = -1 }},
		{"queue", func(s *Snapshot) { s.Next = []Piece{9} }},
		{"overlap", func(s *Snapshot) { s.Board[0] = "JJJJJJJJJJ" }},
	}

	for _, tc := range cases {
		s := good
		s.Board = append([]string(nil), good.Board...)
		s.Next = append([]Piece(nil), good.Next...)
		s.Bag = append([]Piece(nil), good.Bag...)

		tc.mutate(&s)

		if err := s.Validate(); err == nil {
			t.Fatalf("%s snapshot passed validation", tc.name)
		}

		if _, err := FromSnapshot(s); err == nil {
			t.Fatalf("%s snapshot was rebuilt", tc.name)
		}
	}
}

func TestSnapshotDoesNotAliasTheLiveGame(t *testing.T) {
	g := New(3)
	g.Start()

	snap := g.Snapshot()
	snap.Board[0] = "JJJJJJJJJJ"

	if g.Board.rows()[0] == snap.Board[0] {
		t.Fatal("the snapshot board shares memory with the game")
	}
}
