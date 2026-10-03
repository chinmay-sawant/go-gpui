package dino

// styles paints the scene. Every visible part is a plain box or a text
// run, so the page replays as display-list fills and text with no bitmap.
const styles = `
html, body { margin:0; padding:0; background:#f7f7f7; }
.scene { position:relative; width:900px; height:300px; background:#f7f7f7; overflow:hidden; font-family:monospace; color:#535353; }
.ground { position:absolute; left:0; top:248px; width:900px; height:2px; background:#535353; }
.ink { position:absolute; background:#535353; }
.eye { position:absolute; background:#ffffff; }
.cloud { position:absolute; background:#e9e9e9; }
.pebble { position:absolute; background:#cfcfcf; }
.hidden { position:absolute; background:#f7f7f7; }
.hud { position:absolute; right:14px; width:220px; text-align:right; font-size:12px; letter-spacing:1px; color:#9a9a9a; }
.hud.big { font-size:18px; font-weight:700; letter-spacing:2px; color:#535353; }
.overlay { position:absolute; left:0; width:900px; text-align:center; font-size:26px; font-weight:700; letter-spacing:6px; color:#f7f7f7; }
.overlay.small { font-size:13px; font-weight:400; letter-spacing:2px; }
`
