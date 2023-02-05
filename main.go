package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tarkanaciksoz/api-todo-app/config"
	"github.com/tarkanaciksoz/api-todo-app/pkg/server"
)

func main() {
	logger := log.New(os.Stdout, "api-todo-app: ", log.LstdFlags)
	if err := run(logger); err != nil {
		logger.Printf("error - server failed to start. err: %v", err)
	}
}

func run(logger *log.Logger) error {
	config := config.Init(logger)

	s := http.Server{
		Addr:         config.BindAddress,
		Handler:      server.Init(logger),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Printf("Starting server on port %s\n", ":9090")

		err := s.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(ctx)
	if err != nil {
		logger.Printf("Shutdown problem: %s\n", err.Error())
		return err
	}

	return nil
}
