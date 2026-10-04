//go:build linux && !android

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/chinmay-sawant/go-gpui/examples/desktop-cat/cat"
)

// WSLg does not carry this window's alpha and input regions to Windows.
// Run the same example as a Windows window instead.
func launchNative(ctx context.Context, inbox *cat.Inbox, addr string) (bool, error) {
	version, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil || !strings.Contains(strings.ToLower(string(version)), "microsoft") {
		return false, nil
	}

	if _, err := os.Stat("/proc/sys/fs/binfmt_misc/WSLInterop"); err != nil {
		return true, fmt.Errorf("desktop cat: Windows interop is required for the WSL overlay: %w", err)
	}

	dir, err := os.MkdirTemp("", "gpui-cat-")
	if err != nil {
		return true, err
	}
	defer os.RemoveAll(dir)

	_, source, _, _ := runtime.Caller(0)
	exe := filepath.Join(dir, "desktop-cat.exe")
	build := exec.CommandContext(ctx, "go", "build", "-p", "1", "-o", exe, "./examples/desktop-cat")
	build.Dir = filepath.Dir(filepath.Dir(filepath.Dir(source)))
	build.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64", "CGO_ENABLED=0")
	build.Stdout, build.Stderr = os.Stdout, os.Stderr
	log.Print("desktop cat: WSL detected; building and opening a native Windows overlay")
	if err := build.Run(); err != nil {
		return true, fmt.Errorf("desktop cat: Windows build: %w", err)
	}

	args := append(append([]string{}, os.Args[1:]...), "-notify-addr=off", "-stdin-notifications", "-media-endpoint="+nativeEndpoint(addr))
	child := exec.CommandContext(ctx, exe, args...)
	child.Stdout, child.Stderr = os.Stdout, os.Stderr
	pipe, err := child.StdinPipe()
	if err != nil {
		return true, err
	}
	defer pipe.Close()
	go forwardNotifications(ctx, inbox, pipe)
	err = child.Run()
	if ctx.Err() != nil {
		return true, ctx.Err()
	}

	return true, err
}
