// cmd/serve/main.go

package main

import (
	"net/http"

	"github.com/SishaarRao/goyagi/pkg/server"
	"github.com/lob/logger-go"
)

func main() {
	log := logger.New()

	srv := server.New()

	log.Info("Server Started")

	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Err(err).Fatal("Server Stopped")
	}

	log.Info("Server Stopped")
}
