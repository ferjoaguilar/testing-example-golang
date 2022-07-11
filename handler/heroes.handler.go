package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ferjoaguilar/testing-example-golang/models"
	"github.com/ferjoaguilar/testing-example-golang/repository"
	"github.com/ferjoaguilar/testing-example-golang/utils"
)

func CreateHeroe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = models.Heroes{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.ResponseError(w, http.StatusInternalServerError, "Error to parse json information", err.Error())
		}

		newHero := models.Heroes{
			Name:        request.Name,
			Role:        request.Role,
			Description: request.Description,
		}

		reponse, err := repository.CreateHero(r.Context(), &newHero)

		if err != nil {
			utils.ResponseError(w, http.StatusBadRequest, "Create subscription failed", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  true,
			"message": "New hero created",
			"data":    reponse,
		})
	}
}

func CreateHeroeF(w http.ResponseWriter, r *http.Request) {
	var request = models.Heroes{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Error to parse json information", err.Error())
	}

	newHero := models.Heroes{
		Name:        request.Name,
		Role:        request.Role,
		Description: request.Description,
	}

	reponse, err := repository.CreateHero(r.Context(), &newHero)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Create subscription failed", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "New hero created",
		"data":    reponse,
	})
}
