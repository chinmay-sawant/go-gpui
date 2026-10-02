//go:build android || ios

package mobile

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/platform/platform"
)

func start() error {
	app, err := platform.New()
	if err != nil {
		return fmt.Errorf("mobile: platform: %w", err)
	}

	if err := gpui.BindMobile(context.Background(), app.Page()); err != nil {
		return fmt.Errorf("mobile: bind: %w", err)
	}

	return nil
}
