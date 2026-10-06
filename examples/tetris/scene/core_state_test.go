package scene

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
	"github.com/chinmay-sawant/ownframe/examples/tetris/input"
	"github.com/chinmay-sawant/ownframe/examples/tetris/store"
)

func TestCoreSavePointAndRestore(t *testing.T) {
	c := NewCore(6)
	if _, ok := c.SavePoint(); ok {
		t.Fatal("a ready game has a save point")
	}

	c.g.Start()
	snap, ok := c.SavePoint()
	if !ok {
		t.Fatal("a running game has no save point")
	}

	c2 := NewCore(7)
	if err := c2.Restore(snap); err != nil {
		t.Fatal(err)
	}

	f := c2.Frame()
	if f.Score != c.g.Score || f.Level != c.g.Level || f.Piece != c.g.Piece {
		t.Fatalf("restored frame = %+v", f)
	}

	if f.Phase != game.PhasePaused {
		t.Fatalf("restored phase = %v, want paused", f.Phase)
	}
}

func TestCoreConfigureSwapsKeymap(t *testing.T) {
	c := NewCore(8)
	c.g.Start()

	km := input.DefaultKeymap()
	km.Left = []string{"j"}
	c.Configure(Settings{Control: store.Settings{Keymap: km}})

	c.Down("arrowleft")
	c.Step(game.FixedStep)

	if len(c.rec) != 0 {
		t.Fatal("the old keymap still fires")
	}

	c.Down("j")
	c.Step(game.FixedStep)

	if len(c.rec) != 1 || c.rec[0].Action != game.ActionLeft {
		t.Fatalf("recording = %v", c.rec)
	}
}
