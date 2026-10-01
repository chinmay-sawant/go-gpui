// Package window shows the login screen in a window.
//
// On a desktop, Run opens a normal window with a title bar. The person can
// drag the edges. The minimum size is login.MinWidth by login.MinHeight.
// On a phone, the mobile package hands NewGame to the system view, and the
// screen size is the phone. In a browser, build with GOOS=js GOARCH=wasm and
// Run uses the page instead of a desktop window. Resizing the browser changes
// the same frame.
//
// Mouse clicks and taps call login.App.Click. Typed text calls Type.
// Backspace calls Backspace. Enter calls Submit.
package window

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/chinmay-sawant/go-gpui/internal/login"
)

const (
	settleFrames         = 8
	backspaceRepeatAt    = 30
	backspaceRepeatEvery = 4
)

var errNilApp = errors.New("window: nil app")

// pageBackground matches the login page body color.
var pageBackground = color.RGBA{R: 0xf4, G: 0xf1, B: 0xea, A: 0xff}

// Run opens the desktop window, or the browser canvas when built for wasm.
// Run must be called from main, and it returns when the window closes.
func Run(app *login.App) error {
	if app == nil {
		return errNilApp
	}

	width, height := app.Size()
	ebiten.SetWindowTitle("go-gpui")
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowSizeLimits(login.MinWidth, login.MinHeight, -1, -1)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowDecorated(true)
	fmt.Println("opening a window")

	return ebiten.RunGame(NewGame(app))
}

// NewGame returns the screen loop used by Run and by the mobile package.
func NewGame(app *login.App) ebiten.Game {
	width, height := 0, 0
	if app != nil {
		width, height = app.Size()
	}

	return &shell{
		app:      app,
		ctx:      context.Background(),
		pendingW: width,
		pendingH: height,
		screenW:  width,
		screenH:  height,
	}
}

type shell struct {
	app      *login.App
	ctx      context.Context
	img      *ebiten.Image
	seq      uint64
	chars    []rune
	pendingW int
	pendingH int
	screenW  int
	screenH  int
	stable   int
}

func (s *shell) Update() error {
	if s.app == nil {
		return errNilApp
	}

	if err := s.keys(); err != nil {
		return err
	}

	if err := s.pointer(); err != nil {
		return err
	}

	if err := s.resize(); err != nil {
		return err
	}

	return s.syncImage()
}

func (s *shell) keys() error {
	s.chars = ebiten.AppendInputChars(s.chars[:0])
	text := strings.ReplaceAll(string(s.chars), "\r", "")
	text = strings.ReplaceAll(text, "\n", "")

	if text != "" {
		if err := s.app.Type(s.ctx, text); err != nil {
			return err
		}
	}

	if backspaceDue() {
		if err := s.app.Backspace(s.ctx); err != nil {
			return err
		}
	}

	enter := inpututil.IsKeyJustPressed(ebiten.KeyEnter)
	numpad := inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter)
	if enter || numpad {
		return s.app.Submit(s.ctx)
	}

	return nil
}

func backspaceDue() bool {
	held := inpututil.KeyPressDuration(ebiten.KeyBackspace)
	if held == 1 {
		return true
	}

	if held <= backspaceRepeatAt {
		return false
	}

	return (held-backspaceRepeatAt)%backspaceRepeatEvery == 0
}

func (s *shell) pointer() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if err := s.click(x, y); err != nil {
			return err
		}
	}

	for _, id := range inpututil.JustPressedTouchIDs() {
		x, y := ebiten.TouchPosition(id)
		if err := s.click(x, y); err != nil {
			return err
		}
	}

	return nil
}

func (s *shell) click(x, y int) error {
	frameW, frameH := s.frameSize()
	px, py := framePoint(x, y, frameW, frameH, s.screenW, s.screenH)

	return s.app.Click(s.ctx, px, py)
}

func (s *shell) frameSize() (int, int) {
	if s.img != nil {
		bounds := s.img.Bounds()

		return bounds.Dx(), bounds.Dy()
	}

	return s.app.Size()
}

func (s *shell) resize() error {
	wantW, wantH := login.ClampSize(s.pendingW, s.pendingH)
	haveW, haveH := s.app.Size()
	if wantW == haveW && wantH == haveH {
		return nil
	}

	s.stable++
	if s.stable < settleFrames {
		return nil
	}

	s.app.SetSize(wantW, wantH)

	return s.app.Redraw(s.ctx)
}

func (s *shell) syncImage() error {
	if s.app.Generation() == s.seq && s.img != nil {
		return nil
	}

	decoded, err := png.Decode(bytes.NewReader(s.app.PNG()))
	if err != nil {
		return err
	}

	if s.img != nil {
		s.img.Dispose()
	}

	s.img = ebiten.NewImageFromImage(decoded)
	s.seq = s.app.Generation()

	return nil
}

func (s *shell) Draw(screen *ebiten.Image) {
	screen.Fill(pageBackground)

	if s.img == nil {
		return
	}

	bounds := s.img.Bounds()
	if bounds.Dx() == 0 || bounds.Dy() == 0 || s.screenW == 0 || s.screenH == 0 {
		return
	}

	var op ebiten.DrawImageOptions
	op.GeoM.Scale(
		float64(s.screenW)/float64(bounds.Dx()),
		float64(s.screenH)/float64(bounds.Dy()),
	)
	screen.DrawImage(s.img, &op)
}

func (s *shell) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth < 1 {
		outsideWidth = 1
	}

	if outsideHeight < 1 {
		outsideHeight = 1
	}

	if outsideWidth != s.pendingW || outsideHeight != s.pendingH {
		s.pendingW = outsideWidth
		s.pendingH = outsideHeight
		s.stable = 0
	}

	s.screenW = outsideWidth
	s.screenH = outsideHeight

	return outsideWidth, outsideHeight
}
