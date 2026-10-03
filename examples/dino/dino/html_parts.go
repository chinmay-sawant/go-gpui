package dino

// cloudHTML is the six puffs of the two clouds in their starting places.
const cloudHTML = `<div id="c0a" class="cloud" style="left:620px;top:62px;width:26px;height:10px"></div>` +
	`<div id="c0b" class="cloud" style="left:628px;top:54px;width:34px;height:14px"></div>` +
	`<div id="c0c" class="cloud" style="left:650px;top:62px;width:22px;height:10px"></div>` +
	`<div id="c1a" class="cloud" style="left:300px;top:94px;width:26px;height:10px"></div>` +
	`<div id="c1b" class="cloud" style="left:308px;top:86px;width:34px;height:14px"></div>` +
	`<div id="c1c" class="cloud" style="left:330px;top:94px;width:22px;height:10px"></div>`

// dinoHTML is the upright dinosaur at its starting place: tail, body, neck,
// head, eye, arm, and the two legs.
const dinoHTML = `<div id="d-tail" class="ink" style="left:70px;top:218px;width:8px;height:8px"></div>` +
	`<div id="d-body" class="ink" style="left:76px;top:218px;width:26px;height:24px"></div>` +
	`<div id="d-neck" class="ink" style="left:96px;top:210px;width:10px;height:14px"></div>` +
	`<div id="d-head" class="ink" style="left:94px;top:198px;width:22px;height:16px"></div>` +
	`<div id="d-eye" class="eye" style="left:108px;top:202px;width:4px;height:4px"></div>` +
	`<div id="d-arm" class="ink" style="left:102px;top:224px;width:10px;height:4px"></div>` +
	`<div id="d-leg1" class="ink" style="left:80px;top:242px;width:8px;height:8px"></div>` +
	`<div id="d-leg2" class="ink" style="left:92px;top:242px;width:8px;height:8px"></div>`

// textHTML is the frames-per-second readout and the high-score line. The
// frame rate sits at the top-right corner; the score line is under it.
const textHTML = `<div id="t-fps" class="hud" style="top:12px">060 FPS</div>` +
	`<div id="t-score" class="hud big" style="top:30px">HI 00000 00000</div>`
