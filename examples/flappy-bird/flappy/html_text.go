package flappy

// textHTML is the HUD: the score, the ready lines, and the game-over board
// with its lines. Paint writes the live values and empties the rest.
const textHTML = `<div id="board" class="board" style="left:84px;top:236px;width:312px;height:208px"></div>` +
	`<div id="t-score" class="score">0</div>` +
	`<div id="t-title" class="title">FLAPPY BIRD</div>` +
	`<div id="t-hint" class="hint">SPACE OR CLICK TO FLAP</div>` +
	`<div id="t-over" class="over">GAME OVER</div>` +
	`<div id="t-oscore" class="line" style="top:322px">SCORE 0</div>` +
	`<div id="t-obest" class="line" style="top:356px">BEST 0</div>` +
	`<div id="t-again" class="again">PRESS R OR SPACE</div>`
