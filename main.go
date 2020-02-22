package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bungysheep/news-api/pkg/protocols/rest"
	_ "github.com/lib/pq"
)

func main() {
	if err := startUp(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start http server, error: %v", err)
	}
}

func startUp() error {
	restServer := rest.NewRestServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			ctx := context.TODO()

			log.Printf("Stoping http server...\n")
			restServer.Shutdown(ctx)
		}
	}()

	log.Printf("Starting http server...\n")
	return restServer.RunServer()
}
