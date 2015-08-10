package main

import "net/http"

type healthzHandler struct{}

func (h healthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
