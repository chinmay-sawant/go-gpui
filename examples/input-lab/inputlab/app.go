package inputlab

import "github.com/chinmay-sawant/go-gpui"

// App is the combined input screen.
type App struct {
	page     *gpui.Page
	view     View
	lastClip string
	lastEdit string
	eUndo    []string
	eRedo    []string
	cUndo    []csnap
	cRedo    []csnap
	lUndo    []lsnap
	lRedo    []lsnap
}

// csnap saves the two clipboard fields.
type csnap struct {
	Left  string
	Right string
}

// lsnap saves the two sign-in fields.
type lsnap struct {
	Email string
	Pass  string
}
