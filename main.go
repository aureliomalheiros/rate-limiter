package main

import (
	"fmt"
	"net/http"
	_"os"
	_"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/aureliomalheiros/rate-limiter/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.RateLimiterMiddleware)

	// Handlers
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", r)
}
