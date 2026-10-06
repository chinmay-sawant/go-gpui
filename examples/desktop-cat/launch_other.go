//go:build !linux && !js && !android && !ios

package main

import (
	"context"
	"github.com/chinmay-sawant/ownframe/examples/desktop-cat/cat"
)

func launchNative(context.Context, *cat.Inbox, string) (bool, error) {
	return false, nil
}
