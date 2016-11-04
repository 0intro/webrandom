package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
)

func (h *handlers) Null(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	_, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
