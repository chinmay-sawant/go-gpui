package dino

// stepSky drifts the clouds and the ground marks left. The ready screen
// drifts too; a crash freezes the scene.
func (a *App) stepSky(dt float64) {
	if a.game.phase == over {
		return
	}

	for i := range a.clouds {
		a.clouds[i].x -= a.game.speed * 0.14 * dt
		if a.clouds[i].x < -80 {
			a.clouds[i].x = sceneW + 40 + float64(i)*60
		}
	}

	for i := range a.pebbles {
		a.pebbles[i] -= a.game.speed * dt
		if a.pebbles[i] < -16 {
			a.pebbles[i] = sceneW + 20 + float64(i)*13
		}
	}
}
