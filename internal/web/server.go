// Package web shows the login PNG in a browser.
// It is one way to deliver mouse and keyboard events to login.App.
// The native window is internal/window. Pass -web to use this package.
package web

import (
	"bytes"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"math"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/chinmay-sawant/go-gpui/internal/login"
)

const shellHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>go-gpui</title>
</head>
<body>
<p>This page only displays the picture. The login screen is the image, not this HTML.</p>
<img src="/frame.png" usemap="#screen" alt="login screen"{{if .Width}} width="{{.Width}}" height="{{.Height}}"{{end}}>
<map name="screen">
{{range .Areas}}<area shape="rect" coords="{{.Coords}}" href="{{.Href}}" alt="{{.Alt}}">
{{end}}</map>
<form method="post" action="/type">
<label>Text <input type="text" name="text"></label>
<button type="submit">Type</button>
</form>
<form method="post" action="/backspace">
<button type="submit">Backspace</button>
</form>
</body>
</html>
`

type shellArea struct {
	Coords string
	Href   string
	Alt    string
}

type shellData struct {
	Width  int
	Height int
	Areas  []shellArea
}

type server struct {
	mu    sync.Mutex
	app   *login.App
	shell *template.Template
}

// Serve listens on addr and blocks. The page at / shows the latest PNG.
func Serve(app *login.App, addr string) error {
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

func (s *server) page(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data := shellData{
		Areas: areas(s.app),
	}

	if cfg, err := png.DecodeConfig(bytes.NewReader(s.app.PNG())); err == nil {
		data.Width = cfg.Width
		data.Height = cfg.Height
	}

	var buf bytes.Buffer
	if err := s.shell.Execute(&buf, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("write page: %v", err)
	}
}

func (s *server) frame(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pngBytes := s.app.PNG()
	if len(pngBytes) == 0 {
		http.Error(w, "no frame", http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := w.Write(pngBytes); err != nil {
		log.Printf("write png: %v", err)
	}
}

func (s *server) click(w http.ResponseWriter, r *http.Request) {
	x, errX := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	y, errY := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

	if errX != nil || errY != nil {
		http.Error(w, "bad coordinates", http.StatusBadRequest)

		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.app.Click(r.Context(), x, y); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) typeText(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)

		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.app.Type(r.Context(), r.PostFormValue("text")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) backspace(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.app.Backspace(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// areas lists actionable boxes, inner ones first.
// An HTML image map uses the first matching area, which is the opposite
// of the engine's document order.
func areas(app *login.App) []shellArea {
	boxes := app.Boxes()
	out := make([]shellArea, 0)

	for i := len(boxes) - 1; i >= 0; i-- {
		b := boxes[i]
		if b.Action == "" {
			continue
		}

		q := url.Values{}
		q.Set("x", strconv.FormatFloat(b.X+b.W/2, 'f', -1, 64))
		q.Set("y", strconv.FormatFloat(b.Y+b.H/2, 'f', -1, 64))

		alt := b.ID
		if alt == "" {
			alt = b.Action
		}

		out = append(out, shellArea{
			Coords: rectCoords(b.X, b.Y, b.W, b.H),
			Href:   "/click?" + q.Encode(),
			Alt:    alt,
		})
	}

	return out
}

func rectCoords(x, y, w, h float64) string {
	x1 := int(math.Floor(x))
	y1 := int(math.Floor(y))
	x2 := int(math.Ceil(x + w))
	y2 := int(math.Ceil(y + h))

	return strconv.Itoa(x1) + "," + strconv.Itoa(y1) + "," + strconv.Itoa(x2) + "," + strconv.Itoa(y2)
}
