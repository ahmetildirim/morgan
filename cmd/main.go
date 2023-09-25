package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"morgan.io/config"
	"morgan.io/internal/auth"
	"morgan.io/internal/user"
)

const (
	wait time.Duration = time.Second * 5
)

func main() {
	cfg := config.New()

	conn, err := pgx.Connect(context.Background(), cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepository(conn)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	authService := auth.NewService(userService, cfg.SecretKey)
	authHandler := auth.NewHandler(authService)
	r := mux.NewRouter()

	// Add your routes as needed
	r.HandleFunc("/v1/users/register", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/v1/auth/login", authHandler.Login).Methods(http.MethodPost)

	postSubrouter := r.PathPrefix("/v1/posts").Subrouter()
	postSubrouter.Use(auth.AuthMiddleware(cfg.SecretKey))
	postSubrouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from posts"))
	}).Methods(http.MethodGet)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
