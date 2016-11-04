package handlers

import "net/http"

func (h *handlers) Root(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
}
