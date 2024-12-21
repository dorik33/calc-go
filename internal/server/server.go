package server

import (
	"log"
	"net/http"
	"time"

	"github.com/dorik33/calc-go/internal/handler"
	"github.com/dorik33/calc-go/internal/middleware"
)

func RunServer(addr string) {
	mux := http.NewServeMux()
	calculateHandler := middleware.LoggingMiddleware(http.HandlerFunc(handler.CalculateHandler))
	mux.Handle("/api/v1/calculate", calculateHandler)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Server start with addr: %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
