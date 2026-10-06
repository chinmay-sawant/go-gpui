// Package telegram is the Telegram-like demo.
// One HTML template holds the chat list, the contacts tab, the settings tab,
// and the open conversation. A tap opens a chat or switches tabs, typing
// filters the list or writes a message, and Enter sends it. ownframe opens the
// window. This package does not.
package telegram

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 390
	DefaultHeight = 720

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 320
	MinHeight = 480
)
