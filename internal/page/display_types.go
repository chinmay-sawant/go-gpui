package page

import "github.com/chinmay-sawant/blinkless/layout"

// Display is the retained vector list behind a replayable page.
type Display = layout.Display

// DisplayOp is one operation in a Display. A frame callback may change its
// paint fields; see Page.SetTick.
type DisplayOp = layout.DisplayOp

// DisplayOpFillRect, DisplayOpText, and DisplayOpLinkURI are the operation
// kinds the page and frame helpers look for.
const (
	DisplayOpFillRect = layout.DisplayOpFillRect
	DisplayOpText     = layout.DisplayOpText
	DisplayOpLinkURI  = layout.DisplayOpLinkURI
)
