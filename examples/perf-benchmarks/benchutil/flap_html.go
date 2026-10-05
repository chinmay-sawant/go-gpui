package benchutil

// flapHTML is the Benchmark C scene: one bird, three pipe pairs, ground,
// score, and hint. Pipe geometry is inline so the first frame is sane
// before the tick moves the retained ops.
const flapHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Flappy</title>
<style>
html,body{height:100%;margin:0}
body{background:#87ceeb;font-family:sans-serif}
#scene{position:relative;width:480px;height:720px;overflow:hidden}
#bird{position:absolute;width:34px;height:24px;background:#f5c518;border-radius:12px}
.pipe{position:absolute;width:62px;background:#2e9e44}
#ground{position:absolute;left:0;top:624px;width:480px;height:96px;background:#d8c27a}
#score{position:absolute;left:0;top:24px;width:480px;text-align:center;font-size:40px;color:#fff;margin:0}
#hint{position:absolute;left:0;top:660px;width:480px;text-align:center;font-size:14px;color:#5b4a1e;margin:0}
</style>
</head>
<body>
<div id="scene">
<div id="p0t" class="pipe" style="left:210px;top:0;height:205px"></div>
<div id="p0b" class="pipe" style="left:210px;top:395px;height:325px"></div>
<div id="p1t" class="pipe" style="left:390px;top:0;height:125px"></div>
<div id="p1b" class="pipe" style="left:390px;top:315px;height:405px"></div>
<div id="p2t" class="pipe" style="left:100px;top:0;height:180px"></div>
<div id="p2b" class="pipe" style="left:100px;top:370px;height:350px"></div>
<div id="bird" style="left:130px;top:348px"></div>
<div id="ground"></div>
<p id="score">0</p>
<p id="hint">Space or click to flap</p>
</div>
</body>
</html>`
