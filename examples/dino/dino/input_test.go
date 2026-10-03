package dino

import "testing"

func TestKeysControlTheGame(t *testing.T) {
	app, _ := newTestApp(t)

	press(t, app, "arrowup")

	if app.game.phase != running {
		t.Fatal("arrowup did not start the run")
	}

	release(t, app, "arrowup")
	press(t, app, "arrowdown")

	if !app.game.ducking {
		t.Fatal("arrowdown did not duck")
	}

	release(t, app, "arrowdown")

	if app.game.ducking {
		t.Fatal("the dino stayed ducked")
	}

	press(t, app, "w")

	if app.game.onGround {
		t.Fatal("w did not jump")
	}

	release(t, app, "w")
}

func TestDuckStopsJump(t *testing.T) {
	app, _ := newTestApp(t)
	press(t, app, "space")
	press(t, app, "arrowdown")
	press(t, app, "arrowup")

	if !app.game.onGround {
		t.Fatal("the dino jumped while ducking")
	}
}

func TestRestartKey(t *testing.T) {
	app, _ := newTestApp(t)
	press(t, app, "space")
	app.game.phase = over
	app.game.score = 90

	press(t, app, "r")

	if app.game.phase != running || app.game.score != 0 {
		t.Fatal("r did not start a new run")
	}

	if app.game.high != 90 {
		t.Fatal("the best score was lost")
	}
}

func TestUnknownKeysDoNothing(t *testing.T) {
	app, _ := newTestApp(t)
	before := app.game

	press(t, app, "q")
	release(t, app, "q")

	if app.game.phase != before.phase || app.game.score != before.score ||
		app.game.feet != before.feet || app.game.ducking != before.ducking {
		t.Fatal("an unknown key changed the game")
	}
}
