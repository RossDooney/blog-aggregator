package main

import (
	"RossDooney/blog-aggregator/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in env")
	}
	fmt.Printf("Port: %v \n", portString)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL not found in env")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Could not connec to db")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	fmt.Printf("Server starting on port: %v \n", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
