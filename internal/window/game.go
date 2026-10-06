package window

import (
	"context"
	"github.com/chinmay-sawant/ownframe/internal/host"
	"github.com/hajimehoshi/ebiten/v2"
)

// NewGame returns the screen loop Run and BindMobile use.
func NewGame(ctx context.Context, app host.Screen) ebiten.Game {
	if ctx == nil {
		ctx = context.Background()
	}

	width, height := 0, 0
	if app != nil {
		width, height = app.Size()
	}

	game := &shell{
		app:      app,
		ctx:      ctx,
		chords:   newChordWatch(),
		watched:  newKeyWatch(),
		pageZoom: 1,
		dev:      devState{watch: newKeyWatch()},
		pendingW: width,
		pendingH: height,
		screenW:  width,
		screenH:  height,
		perf:     perfEnabled(app),
	}
	game.wirePerf()
	game.imeInit()

	return game
}
