package main

import (
	"net/http"
)

type exampleHandler struct {
	logger Logger
}

func (h exampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start Example Handler")
	defer h.logger.Info("End Example Handler")

	w.WriteHeader(200)
	w.Write([]byte("Example"))
}
