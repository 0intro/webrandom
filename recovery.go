package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func Recovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				b := make([]byte, 32768)
				n := runtime.Stack(b, false)
				fmt.Println(string(b[:n]))
				http.Error(w, "", http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
