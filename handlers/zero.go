package handlers

import (
	"net/http"
	"strconv"
)

func (h *handlers) Zero(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path[len("/zero/"):]
	if path == "" {
		http.NotFound(w, r)
		return
	}
	u, err := strconv.ParseUint(path, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nbytes := int(u)
	var b [BufSize]byte
	n := len(b)
	m := 0
	for tot := 0; tot < nbytes; tot += m {
		if nbytes-tot < n {
			n = nbytes - tot
		}
		m, err = w.Write(b[:n])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
