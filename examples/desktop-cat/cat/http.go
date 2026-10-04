package cat

import (
	"encoding/json"
	"io"
	"net/http"
)

// Handler exposes the local agent notification API.
func (i *Inbox) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /expressions", func(w http.ResponseWriter, r *http.Request) {
		reply(w, http.StatusOK, Expressions())
	})
	mux.HandleFunc("GET /state", func(w http.ResponseWriter, r *http.Request) {
		n, v := i.Latest()
		reply(w, http.StatusOK, struct {
			Notification
			Sequence uint64 `json:"sequence"`
		}{n, v})
	})
	mux.HandleFunc("POST /notify", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		var n Notification
		if err := decoder.Decode(&n); err != nil {
			reply(w, 400, map[string]string{"error": err.Error()})
			return
		}
		var extra any
		if err := decoder.Decode(&extra); err != io.EOF {
			reply(w, 400, map[string]string{"error": "send one JSON object"})
			return
		}
		if err := i.Submit(n); err != nil {
			reply(w, 400, map[string]string{"error": err.Error()})
			return
		}
		reply(w, http.StatusAccepted, n)
	})
	return mux
}

func reply(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
