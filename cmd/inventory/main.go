package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JJoddyZZ/home-inventory/internal/repository"
	"github.com/JJoddyZZ/home-inventory/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	// TODO: load config

	// db client
	db, err := sql.Open("mysql", "root:password@/mysql") // TODO: get from config
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// TODO: move this
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	st := repository.NewStorage(db)
	err = st.Ping()
	if err != nil {
		log.Panic().Err(err).Msg("error pinging db")
	}

	// initialize server and wait for stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()
	addRoutes(r)
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
