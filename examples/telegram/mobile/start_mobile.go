//go:build android || ios

package mobile

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/telegram/telegram"
)

func start() error {
	a, err := telegram.New()
	if err != nil {
		return fmt.Errorf("mobile: telegram: %w", err)
	}

	app = a

	if err := gpui.BindMobile(context.Background(), a.Page()); err != nil {
		return fmt.Errorf("mobile: bind: %w", err)
	}

	return nil
}
