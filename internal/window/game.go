package window

import (
	"context"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/host"
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

type shell struct {
	transparent bool
	interactive func(int, int) bool
	draggable   func(int, int) bool
	perf        bool
	windowDrag  windowDrag
	passthrough bool
	app         host.Screen
	ctx         context.Context
	img         *ebiten.Image
	display     *layout.Display
	fallback    bool
	seq         uint64
	contentW    int
	contentH    int
	contentGen  uint64
	chars       []rune
	chords      chordWatch
	watched     keyWatch
	pendingW    int
	pendingH    int
	screenW     int
	screenH     int
	scrollX     int
	scrollY     int
	// redrawX and redrawY are the scroll offsets the current display was
	// built with; viewport-pinned layers draw against them.
	redrawX  int
	redrawY  int
	dragAxis int
	dragGrab float64
	hold     longPressWatch

	mouseDown        bool
	fingers          touchGesture
	pageZoom         float64
	tabEaten         bool
	f11Eaten         bool
	clicks           clickWatch
	dragActive       bool
	dragX, dragY     float64
	menu             menuState
	cursor           ebiten.CursorShapeType
	setCursor        func(ebiten.CursorShapeType)
	readFullscreen   func() bool
	applyFullscreen  func(bool)
	moving           bool
	lastRelayout     time.Time
	cursorX, cursorY int

	replayBuf *ebiten.Image

	lastPoll time.Time
	lastNote string
	viewport viewportState
	partial  partialState
	dev      devState
	ime      imeState
	commits  uint64
	skipped  uint64
}
