package flappy

import (
	"fmt"
	"strings"
)

// groundHTML is the grass line at the scene floor, the sand under it, and
// the six scrolling dashes.
const groundHTML = `<div id="gtop" class="grass" style="left:0;top:624px;width:480px;height:16px"></div>` +
	`<div id="ground" class="sand" style="left:0;top:640px;width:480px;height:80px"></div>` +
	`<div id="s0" class="dash" style="left:0px;top:662px;width:44px;height:10px"></div>` +
	`<div id="s1" class="dash" style="left:80px;top:662px;width:44px;height:10px"></div>` +
	`<div id="s2" class="dash" style="left:160px;top:662px;width:44px;height:10px"></div>` +
	`<div id="s3" class="dash" style="left:240px;top:662px;width:44px;height:10px"></div>` +
	`<div id="s4" class="dash" style="left:320px;top:662px;width:44px;height:10px"></div>` +
	`<div id="s5" class="dash" style="left:400px;top:662px;width:44px;height:10px"></div>`

// pipeHTML is the three pipe slots. Paint moves and hides every part, so
// these starting places only matter before the first frame.
func pipeHTML() string {
	var b strings.Builder

	for slot := range pipeSlots {
		x := 210 + slot*90

		fmt.Fprintf(&b, `<div id="%s" class="pipe" style="left:%dpx;top:0;width:62px;height:200px"></div>`,
			pipeID(slot, "t"), x)
		fmt.Fprintf(&b, `<div id="%s" class="cap" style="left:%dpx;top:174px;width:74px;height:26px"></div>`,
			pipeID(slot, "tc"), x-6)
		fmt.Fprintf(&b, `<div id="%s" class="pipe" style="left:%dpx;top:426px;width:62px;height:198px"></div>`,
			pipeID(slot, "b"), x)
		fmt.Fprintf(&b, `<div id="%s" class="cap" style="left:%dpx;top:400px;width:74px;height:26px"></div>`,
			pipeID(slot, "bc"), x-6)
	}

	return b.String()
}
