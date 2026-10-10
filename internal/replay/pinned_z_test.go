package replay

import (
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

func TestPinZSeparatesScrollingAndViewportOperations(t *testing.T) {
	body := &layout.DisplayOp{ZIndex: 1, ZIndexSet: true}
	bar := &layout.DisplayOp{ZIndex: 2, ZIndexSet: true}
	fixed := &layout.DisplayOp{Fixed: true, ZIndex: 3, ZIndexSet: true}
	if !scrollContentOp(body, 2) || scrollContentOp(bar, 2) || scrollContentOp(fixed, 2) {
		t.Fatal("scroll buffer included a pinned or fixed operation")
	}
	if atPinZ(body, 2) || !atPinZ(bar, 2) {
		t.Fatal("viewport pin threshold selected the wrong layer")
	}
}
