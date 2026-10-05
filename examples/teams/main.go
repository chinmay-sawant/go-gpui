//go:build !js

// Command teams opens the Microsoft Teams-like example in a native window.
// Pass -web to serve the picture in a browser instead.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/teams/app"
)

func main() {
	webMode := flag.Bool("web", false, "serve the picture in a browser on -addr")
	addr := flag.String("addr", "127.0.0.1:8118", "listen address for -web")
	dbPath := flag.String("db", defaultDBPath(), "SQLite state file; :memory: keeps it in memory")
	flag.Parse()

	screen, err := app.New(app.WithDB(*dbPath))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if *webMode {
		if err := gpui.Serve(ctx, screen.Page(), *addr); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := gpui.Run(ctx, screen.Page()); err != nil {
		log.Fatal(err)
	}
}

// defaultDBPath returns the per-user state file. An empty string means the
// in-memory database.
func defaultDBPath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, "go-gpui-teams", "teams.db")
}
