package flappy

// buildHTML assembles the scene: the clouds, the pipe slots, the ground,
// the bird, and the text. Paint moves and hides the parts it owns.
func buildHTML() string {
	return `<!DOCTYPE html><html><head><meta charset="utf-8">` +
		`<title>Flappy Bird</title><style>` + styles + `</style></head>` +
		`<body><div class="scene" id="scene">` +
		cloudHTML +
		pipeHTML() +
		groundHTML +
		birdHTML +
		textHTML +
		`</div></body></html>`
}
