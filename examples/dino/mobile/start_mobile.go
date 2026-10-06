//go:build android || ios

package mobile

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/dino/dino"
)

func start() error {
	app, err := dino.New()
	if err != nil {
		return fmt.Errorf("mobile: dino: %w", err)
	}

	app.BindTouch()

	if err := gpui.BindMobile(context.Background(), app.Page()); err != nil {
		return fmt.Errorf("mobile: bind: %w", err)
	}

	return nil
}
