package dino

import "testing"

func TestDuckShrinksTheBox(t *testing.T) {
	app, _ := newTestApp(t)
	press(t, app, "space")

	if _, _, _, tall := app.game.dinoBox(); tall != 48 {
		t.Fatalf("standing box height = %v, want 48", tall)
	}

	press(t, app, "arrowdown")

	x, y, w, h := app.game.dinoBox()
	if h != 26 || w != 46 || x != dinoX+4 || y != groundY-26 {
		t.Fatalf("duck box = %v %v %v %v", x, y, w, h)
	}

	release(t, app, "arrowdown")

	if _, _, _, tall := app.game.dinoBox(); tall != 48 {
		t.Fatal("the box stayed low after the release")
	}
}

func TestBigCactusForcesAJump(t *testing.T) {
	app, _ := newTestApp(t)
	press(t, app, "space")

	// A cactus at the dino's feet hits the standing box.
	app.game.obstacles = []obstacle{{kind: cactusBig, x: dinoX, w: 30, h: 48}}
	if !app.game.hits() {
		t.Fatal("the standing dino missed a cactus at its feet")
	}

	// The same cactus passes under a jump.
	app.game.onGround = false
	app.game.feet = 80

	if app.game.hits() {
		t.Fatal("the airborne dino still hit the cactus")
	}
}
