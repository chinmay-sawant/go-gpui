package dino

// Overlay messages. Each one sits in its own band so no two text boxes
// overlap.
const (
	startText = "PRESS SPACE TO START"
	keysText  = "UP OR SPACE JUMP - DOWN DUCK"
	overText  = "G A M E  O V E R"
	againText = "PRESS SPACE, UP OR R TO RUN AGAIN"
)

// Touch messages stand in for the key hints when BindTouch is used.
const (
	tapStartText = "TAP TO START"
	tapKeysText  = "TAP TO JUMP"
	tapAgainText = "TAP TO RUN AGAIN"
)
