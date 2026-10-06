package scene

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestCoreCaptureAndReplay(t *testing.T) {
	c := NewCore(4)
	c.g.Start()
	c.Down("arrowleft")
	c.Step(game.FixedStep)

	c.g.Phase = game.PhaseOver
	c.watch()

	res, rep, ok := c.TakeCompleted()
	if !ok {
		t.Fatal("no completed run")
	}

	if res.Score != c.g.Score {
		t.Fatalf("result score = %d", res.Score)
	}

	if rep == nil || len(rep.Events) != 1 || rep.Events[0].Step != 0 {
		t.Fatalf("replay = %v", rep)
	}

	if _, _, ok := c.TakeCompleted(); ok {
		t.Fatal("the run was returned twice")
	}
}

func TestCoreReplayReproducesTheRun(t *testing.T) {
	c := NewCore(5)
	c.g.Start()

	c.Down("arrowleft")
	c.Step(game.FixedStep)
	c.Down("space")
	c.Step(game.FixedStep)

	c.g.Phase = game.PhaseOver
	c.watch()

	_, rep, ok := c.TakeCompleted()
	if !ok {
		t.Fatal("no completed run")
	}

	got, err := rep.Play(0)
	if err != nil {
		t.Fatal(err)
	}

	if got.Score != c.g.Score {
		t.Fatalf("replay score = %d, live = %d", got.Score, c.g.Score)
	}
}
