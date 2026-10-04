// Package web shows a screen's PNG in a browser.
// It is one way to deliver mouse and keyboard events to a host.Screen.
// The native window is internal/window.
package web

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"sync"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

type shellArea struct {
	Coords string
	Href   string
	Alt    string
}

type shellData struct {
	Width  int
	Height int
	Areas  []shellArea
	Reload bool
}

type server struct {
	mu       sync.Mutex
	app      host.Screen
	shell    *template.Template
	lastNote string
}

// Serve listens on addr and blocks. The page at / shows the latest PNG.
func Serve(app host.Screen, addr string) error {
	shell, err := template.New("shell").Parse(shellHTML)
	if err != nil {
		return err
	}

	srv := &server{
		app:   app,
		shell: shell,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", srv.page)
	mux.HandleFunc("GET /frame.png", srv.frame)
	mux.HandleFunc("GET /click", srv.click)
	mux.HandleFunc("POST /type", srv.typeText)
	mux.HandleFunc("POST /backspace", srv.backspace)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	fmt.Println("http://" + addr + "/")

	return http.Serve(ln, mux)
}
