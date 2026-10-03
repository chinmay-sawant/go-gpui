package flappy

// stepScene drifts the clouds and the ground marks left. The ready screen
// drifts too; a crash freezes the scene.
func (a *App) stepScene(dt float64) {
	if a.game.phase == over {
		return
	}

	drift := a.game.speed * dt

	for i := range a.clouds {
		a.clouds[i] -= drift * 0.18

		if a.clouds[i] < -90 {
			a.clouds[i] = sceneW + 60 + float64(i)*40
		}
	}

	for i := range a.stripes {
		a.stripes[i] -= drift

		if a.stripes[i] < -50 {
			a.stripes[i] = sceneW + 30 + float64(i)*12
		}
	}
}
