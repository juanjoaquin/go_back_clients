package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/juanjoaquin/go-back-clients/internal"
)

func main() {
	_ = godotenv.Load()

	logger := internal.InitLogger()
	logger.Println("🚀 Starting go-backend-clients...")

	_, err := internal.DBConnection()
	if err != nil {
		logger.Fatalf("❌ Database connection failed: %v", err)
	}

	logger.Println("✅ Database connected successfully")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// healthcheck simple
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("🟢 Server listening on port", port)
	log.Fatal(srv.ListenAndServe())
}
