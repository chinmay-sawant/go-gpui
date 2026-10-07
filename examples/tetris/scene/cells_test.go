package scene

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestClassAt(t *testing.T) {
	var f Frame
	f.Board[0][0] = game.Cell(game.PieceO)
	f.Piece = game.PieceS
	f.Cells = [4]game.Point{{X: 1, Y: 1}}
	f.Ghost = [4]game.Point{{X: 2, Y: 2}}

	cases := []struct {
		r, c int
		want string
	}{
		{0, 0, "k-o"},
		{1, 1, "k-s"},
		{2, 2, "k-g"},
		{9, 9, ""},
	}

	for _, c := range cases {
		if got := classAt(f, c.r, c.c); got != c.want {
			t.Fatalf("classAt(%d,%d) = %q, want %q", c.r, c.c, got, c.want)
		}
	}
}
