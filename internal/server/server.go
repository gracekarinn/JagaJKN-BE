package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	router *http.Handler
	port   string
}

func NewServer(handler http.Handler, port string) *Server {
	return &Server{
		router: &handler,
		port:   port,
	}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:         ":" + s.port,
		Handler:      *s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverErrors := make(chan error, 1)


	go func() {
		log.Printf("API listening on port %s", s.port)
		serverErrors <- server.ListenAndServe()
	}()


	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)


	select {
	case err := <-serverErrors:
		return fmt.Errorf("error starting server: %v", err)

	case <-shutdown:
		log.Println("Starting shutdown...")
		

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("could not stop server gracefully: %v", err)
		}
	}

	return nil
}