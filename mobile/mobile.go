// Package mobile is the Android and iOS entry.
//
// Desktop and browser builds use cmd/go-gpui. A phone build binds this
// package with ebitenmobile. The generated view calls the game registered
// here. Do not call ebiten.RunGame from this package.
package mobile

import (
	"context"
	"fmt"

	ebitenmobile "github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/chinmay-sawant/go-gpui/internal/login"
	"github.com/chinmay-sawant/go-gpui/internal/window"
)

func init() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	app, err := login.New()
	if err != nil {
		return fmt.Errorf("mobile: login: %w", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		return fmt.Errorf("mobile: redraw: %w", err)
	}

	ebitenmobile.SetGame(window.NewGame(app))

	return nil
}

// Dummy exists so ebitenmobile bind will compile this package.
// The bind tool skips a package that exports nothing.
func Dummy() {}
