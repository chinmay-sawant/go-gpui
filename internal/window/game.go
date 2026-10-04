package window

import (
	"context"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// NewGame returns the screen loop used by Run and by BindMobile.
func NewGame(ctx context.Context, app host.Screen) ebiten.Game {
	if ctx == nil {
		ctx = context.Background()
	}

	width, height := 0, 0
	if app != nil {
		width, height = app.Size()
	}

	return &shell{
		app:      app,
		ctx:      ctx,
		chords:   newChordWatch(),
		watched:  newKeyWatch(),
		pendingW: width,
		pendingH: height,
		screenW:  width,
		screenH:  height,
	}
}

type shell struct {
	app      host.Screen
	ctx      context.Context
	img      *ebiten.Image
	display  *layout.Display
	fallback bool
	seq      uint64
	chars    []rune
	chords   chordWatch
	watched  keyWatch
	pendingW int
	pendingH int
	screenW  int
	screenH  int
	scrollX  int
	scrollY  int
	dragAxis int
	dragGrab float64

	mouseDown bool
	touches   []ebiten.TouchID

	replayBuf *ebiten.Image

	lastPoll time.Time
	lastNote string
}

func (s *shell) Update() error {
	if s.app == nil {
		return errNilApp
	}

	if err := s.ctx.Err(); err != nil {
		return err
	}

	if err := s.keys(); err != nil {
		return err
	}

	if err := s.pointer(); err != nil {
		return err
	}

	if err := s.dropPass(ebiten.DroppedFiles()); err != nil {
		return err
	}

	if err := s.resize(); err != nil {
		return err
	}

	if err := s.tickFrame(); err != nil {
		return err
	}

	s.pollReload()

	if err := s.syncImage(); err != nil {
		return err
	}

	s.wheel()

	return nil
}

// tickFrame runs the screen's per-frame callback when it has one.
func (s *shell) tickFrame() error {
	ticker, ok := s.app.(host.Ticker)
	if !ok {
		return nil
	}

	return ticker.Tick(s.ctx)
}
