package flappy

// cloudHTML is the three drifting clouds in their starting places. Each
// cloud is one rounded fill; the tick drifts it left and wraps it.
const cloudHTML = `<div id="c0" class="cloud" style="left:40px;top:96px;width:78px;height:26px"></div>` +
	`<div id="c1" class="cloud" style="left:210px;top:300px;width:64px;height:26px"></div>` +
	`<div id="c2" class="cloud" style="left:380px;top:120px;width:72px;height:26px"></div>`

// birdHTML is the six fills of the bird, facing right: the body, the belly,
// the wing, the eye, the pupil, and the beak.
const birdHTML = `<div id="b-body" class="bird" style="left:113px;top:347px;width:34px;height:26px"></div>` +
	`<div id="b-belly" class="belly" style="left:117px;top:360px;width:26px;height:11px"></div>` +
	`<div id="b-wing" class="wing" style="left:115px;top:354px;width:15px;height:10px"></div>` +
	`<div id="b-eye" class="eye" style="left:138px;top:349px;width:9px;height:9px"></div>` +
	`<div id="b-pupil" class="pupil" style="left:141px;top:351px;width:4px;height:4px"></div>` +
	`<div id="b-beak" class="beak" style="left:145px;top:356px;width:11px;height:8px"></div>`
