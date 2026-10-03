package dino

// Overlay messages. Each one sits in its own band so no two text boxes
// overlap.
const (
	startText = "PRESS SPACE TO START"
	keysText  = "UP OR SPACE JUMP - DOWN DUCK"
	overText  = "G A M E  O V E R"
	againText = "PRESS SPACE, UP OR R TO RUN AGAIN"
)

// overlayHTML is the ready and game-over messages. They start in the
// background colour; the paint step colours and fills them in.
func overlayHTML() string {
	return `<div id="t-start" class="overlay" style="top:92px">` + startText + `</div>` +
		`<div id="t-skeys" class="overlay small" style="top:134px">` + keysText + `</div>` +
		`<div id="t-over" class="overlay" style="top:168px">` + overText + `</div>` +
		`<div id="t-again" class="overlay small" style="top:208px">` + againText + `</div>`
}
