package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

func main() {
	port := flag.Int("p", 8080, "Server port")
	dir := flag.String("d", ".", "Static directory")
	flag.Parse()

	path, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(path))
	handler := noCache(fileServer)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: handler,
	}

	go start(server)
	waitForSignal(server)
}

func start(server *http.Server) {
	log.Printf("Server at http://localhost%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Error starting server:", err)
	}
}

func waitForSignal(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error stopping server:", err)
	}
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}
