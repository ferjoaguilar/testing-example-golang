package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ferjoaguilar/testing-example-golang/database"
	"github.com/ferjoaguilar/testing-example-golang/handler"
	"github.com/ferjoaguilar/testing-example-golang/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	bindRoutes(router)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")
	DATABASEURL := os.Getenv("DATABASE_URL")
	DBNAME := os.Getenv("DATABASE_NAME")

	// Start database
	repo, err := database.NewMongoRepository(DATABASEURL, DBNAME)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetHeroesRepository(repo)

	log.Printf("Server running on port %s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

func bindRoutes(r *mux.Router) {
	r.NotFoundHandler = notFound()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/heroes", handler.CreateHeroe()).Methods(http.MethodPost)
}

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"Message": "Endpoint not found",
			"Url":     r.URL.Path,
		})
	}
}
