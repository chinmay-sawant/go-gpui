package ownframe

import "github.com/chinmay-sawant/ownframe/internal/page"

// Display is the retained vector list behind a replayable page.
type Display = page.Display

// DisplayOp is one operation in a Display. A page tick may change its paint
// fields, which changes the next drawn frame without another Redraw.
type DisplayOp = page.DisplayOp

const (
	// DisplayOpFillRect is a filled rectangle.
	DisplayOpFillRect = page.DisplayOpFillRect

	// DisplayOpText is one shaped text run.
	DisplayOpText = page.DisplayOpText
)
