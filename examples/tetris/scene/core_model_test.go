package scene

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestCoreFrameMapsRunningGame(t *testing.T) {
	c := NewCore(1)
	c.g.Start()

	f := c.Frame()
	if f.Phase != game.PhaseRunning {
		t.Fatalf("phase = %v", f.Phase)
	}

	if !f.Piece.Valid() || !f.Next.Valid() {
		t.Fatalf("piece = %v next = %v", f.Piece, f.Next)
	}

	if f.Level != 1 || f.Score != 0 {
		t.Fatalf("level = %d score = %d", f.Level, f.Score)
	}
}

func TestCoreInputReachesTheGame(t *testing.T) {
	c := NewCore(2)
	c.g.Start()
	x := c.g.X

	c.Down("arrowleft")
	c.Step(game.FixedStep)

	if c.g.X != x-1 {
		t.Fatalf("x = %d, want %d", c.g.X, x-1)
	}

	if len(c.rec) != 1 || c.rec[0].Action != game.ActionLeft {
		t.Fatalf("recording = %v", c.rec)
	}
}

func TestCoreRecordingSkipsPause(t *testing.T) {
	c := NewCore(3)
	c.g.Start()

	c.Down("p")
	c.Step(game.FixedStep)

	if len(c.rec) != 0 {
		t.Fatal("a pause action was recorded")
	}

	if c.g.Phase != game.PhasePaused {
		t.Fatalf("phase = %v", c.g.Phase)
	}
}

func TestCoreRestartClearsTheRecording(t *testing.T) {
	c := NewCore(9)
	c.g.Start()
	c.Down("arrowleft")
	c.Step(game.FixedStep)

	c.Restart()

	if len(c.rec) != 0 || c.done != nil || c.taken {
		t.Fatal("a restart kept the old run")
	}

	if c.g.Phase != game.PhaseRunning {
		t.Fatalf("restart phase = %v", c.g.Phase)
	}
}
