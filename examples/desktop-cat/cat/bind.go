package cat

import (
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
)

func (a *animation) bind(d *gpui.Display) error {
	for i := range d.Ops {
		op := &d.Ops[i]
		if data, _, _ := op.ImageBytes(); len(data) != 0 {
			a.display, a.image = d, op
			a.y = op.Y

			return nil
		}
	}

	return fmt.Errorf("desktop cat: no registered PNG in the display list")
}
