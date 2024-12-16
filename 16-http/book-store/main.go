package main

import (
	"book-store/handlers"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// initialize http service
	//chi, http.DefaultServeMux, gin
	api := http.Server{
		Addr:              ":8082",
		ReadHeaderTimeout: time.Second * 200,
		WriteTimeout:      time.Second * 200,
		IdleTimeout:       time.Second * 200,
		Handler:           handlers.SetupGINRoutes(),
	}
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGTERM)
	serverError := make(chan error)

	go func() {
		serverError <- api.ListenAndServe()
	}()

	select {
	// this error would happen if the service is not able to start
	case err := <-serverError:
		panic(err)
	case <-shutdown:
		fmt.Println("Graceful Shutdown Server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		err := api.Shutdown(ctx)
		if err != nil {
			// force close
			err := api.Close()
			panic(err)
		}

	}

}
