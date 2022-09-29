package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JJoddyZZ/home-inventory/internal/server"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	r := mux.NewRouter()
	addRoutes(r)

	// initialize server and wait for stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	srv := server.NewServer("8080", r) //TODO: change for a configuration
	go srv.Serve()
	log.Info().Msg("server started")
	<-stop

	// clean up
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //TODO: change for a configuration
	defer cancel()
	srv.Shutdown(ctx)
}

func addRoutes(r *mux.Router) {
	r.HandleFunc("/ping", pong)
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}
