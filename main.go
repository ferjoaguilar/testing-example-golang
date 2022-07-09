package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ferjoaguilar/testing-example-golang/database"
	"github.com/ferjoaguilar/testing-example-golang/handler"
	"github.com/ferjoaguilar/testing-example-golang/repository"
	"github.com/ferjoaguilar/testing-example-golang/utils"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	bindRoutes(router)

	port, err := utils.Enviroment("PORT")
	if err != nil {
		log.Fatal(err)
	}
	url, _ := utils.Enviroment("DATABASE_URL")
	dbname, _ := utils.Enviroment("DATABASE_NAME")

	// Start database
	repo, err := database.NewMongoRepository(url, dbname)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetHeroesRepository(repo)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
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
