package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type exampleHandler struct {
	logger *logrus.Logger
}

func (h exampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start Example Handler")
	defer h.logger.Info("End Example Handler")

	w.WriteHeader(200)
	w.Write([]byte("Example"))
}
