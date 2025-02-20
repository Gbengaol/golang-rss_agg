package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gbengaol/rss/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
		// Load environmenr variables from .env file
		godotenv.Load(".env")

		portString := os.Getenv("PORT")
		if portString == "" {
			log.Fatal("PORT is not found in environment")
		}

		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			log.Fatal("DB_URL is not found in environment")
		}

		conn, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatal("Cannot connect to database", err)
		}

		apiCfg := apiConfig{
			DB: database.New(conn),
		}

		router := chi.NewRouter()
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"Link"},
			AllowCredentials: false,
			MaxAge: 300,
		}))

		v1Router := chi.NewRouter()
		v1Router.Get("/healthz", handlerReadiness)
		v1Router.Get("/err", handlerError)
		v1Router.Post("/users", apiCfg.handlerCreateUser)

		router.Mount("/v1", v1Router)

		log.Println("Server is running on port", portString)
		http.ListenAndServe(":"+portString, router)
}