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
	server := startServer(handler, port)
	shutdownServerGracefully(server, shutdownTimeout)
}

func startServer(handler http.Handler, port int) *http.Server {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Server started on %s", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			// Cannot panic, because this probably is an intentional close.
			log.Println(err)
		}
	}()

	return server
}

func shutdownServerGracefully(
	server *http.Server, shutdownTimeout time.Duration,
) {
	// Wait for interrupt signal to gracefully shutdown the server
	// with a specified timeout.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	log.Println("Server gracefully stopped.")
}
