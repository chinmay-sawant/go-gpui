package game

import "testing"

func TestTPieceRotationsMatchSRS(t *testing.T) {
	cases := []struct {
		r    Rotation
		want []Point
	}{
		{Rot0, []Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}}},
		{RotR, []Point{{2, 1}, {1, 0}, {1, 1}, {1, 2}}},
		{Rot2, []Point{{1, 2}, {2, 1}, {1, 1}, {0, 1}}},
		{RotL, []Point{{0, 1}, {1, 0}, {1, 1}, {1, 2}}},
	}

	for _, tc := range cases {
		if got := PieceT.Cells(tc.r); !cellsEqual(got, tc.want...) {
			t.Fatalf("T %v = %v, want %v", tc.r, got, tc.want)
		}
	}
}

func TestIPieceRotationsMatchSRS(t *testing.T) {
	cases := []struct {
		r    Rotation
		want []Point
	}{
		{RotR, []Point{{2, 0}, {2, 1}, {2, 2}, {2, 3}}},
		{Rot2, []Point{{0, 2}, {1, 2}, {2, 2}, {3, 2}}},
		{RotL, []Point{{1, 0}, {1, 1}, {1, 2}, {1, 3}}},
	}

	for _, tc := range cases {
		if got := PieceI.Cells(tc.r); !cellsEqual(got, tc.want...) {
			t.Fatalf("I %v = %v, want %v", tc.r, got, tc.want)
		}
	}
}

func TestCellsDoNotAliasTheBase(t *testing.T) {
	base := PieceL.Cells(Rot0)

	_ = PieceL.Cells(RotR)

	if got := PieceL.Cells(Rot0); !cellsEqual(got, base...) {
		t.Fatalf("rotating changed the spawn cells: %v", got)
	}
}

func TestRotationCycles(t *testing.T) {
	r := Rot0
	for i := 0; i < 4; i++ {
		r = r.CW()
	}

	if r != Rot0 {
		t.Fatalf("four clockwise turns reached %v", r)
	}

	if Rot0.CCW() != RotL {
		t.Fatalf("counter-clockwise from 0 is %v", Rot0.CCW())
	}
}

func TestPieceNamesAndValidity(t *testing.T) {
	want := map[Piece]string{
		PieceI: "I", PieceO: "O", PieceT: "T", PieceS: "S",
		PieceZ: "Z", PieceJ: "J", PieceL: "L",
	}

	for p, name := range want {
		if p.String() != name {
			t.Fatalf("%d names %q, want %q", p, p.String(), name)
		}

		if !p.Valid() {
			t.Fatalf("%d should be valid", p)
		}
	}

	if Piece(0).Valid() || Piece(8).Valid() {
		t.Fatal("zero and eight must not be valid pieces")
	}
}
