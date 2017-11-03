package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Serve serves the handler at specified port
// and shuts down server gracefully in a shutdown timeout.
func Serve(handler http.Handler, port int, shutdownTimeout time.Duration) {
	srv := startServer(handler, port)
	shutdownServerGracefully(srv, shutdownTimeout)
}

func startServer(handler http.Handler, port int) *http.Server {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Server started on %s", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// Cannot panic, because this probably is an intentional close.
			log.Println(err)
		}
	}()

	return srv
}

func shutdownServerGracefully(
	srv *http.Server, shutdownTimeout time.Duration,
) {
	// Wait for interrupt signal to gracefully shutdown the server
	// with a specified timeout.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	log.Println("Server gracefully stopped.")
}
