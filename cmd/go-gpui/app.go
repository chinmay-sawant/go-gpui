package main

import (
	"context"
	"log"

	"github.com/chinmay-sawant/go-gpui/internal/login"
)

func mustApp() *login.App {
	app, err := login.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		log.Fatal(err)
	}

	return app
}
