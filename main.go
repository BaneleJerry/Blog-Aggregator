package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BaneleJerry/Blog-Aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading ENV")
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	if port == "" || dbURL == "" {
		panic("Could not get ENV")
	}
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	dbQuries := database.New(db)

	cfg := apiConfig{
		DB: dbQuries,
	}

	filepathRoot := "/"

	mux := http.NewServeMux()

	mux.HandleFunc(filepathRoot, greet)
	mux.HandleFunc("GET /v1/readiness", readinessHandler)
	mux.HandleFunc("GET /v1/err", errorHandler)
	mux.HandleFunc("POST /v1/users", cfg.createUserHandler)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.getUserHandler))
	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.createFeedHandler))
	mux.HandleFunc("GET /v1/feeds", cfg.getFeedsHandler)
	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.handlerCreateFeedFollower))
	mux.HandleFunc("GET /v1/feed_follows", cfg.middlewareAuth(cfg.handlerGetFeedFollowers))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.handlerDeleteFeedFollower))

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
