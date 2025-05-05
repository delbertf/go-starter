package main

import (
	"fmt"
	"go2start/internal/assets"
	"go2start/internal/features/home"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	router := chi.NewRouter()

	// mount features to router
	home.Mount(router, home.NewHandler("hello"))

	// assets
	assets.Mount(router)

	server := &http.Server{
		Addr:    os.Getenv("HTTP_LISTEN_ADDR"),
		Handler: http.TimeoutHandler(router, 30*time.Second, "request timed out"),
	}

	// Display the localhost address and port
	fmt.Printf("Listening on http://localhost%s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
