package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ferjoaguilar/testing-example-golang/handler"
	"github.com/ferjoaguilar/testing-example-golang/models"
	"github.com/ferjoaguilar/testing-example-golang/repository"
	"github.com/ferjoaguilar/testing-example-golang/repository/mocks"
	"go.mongodb.org/mongo-driver/mongo"
)

var heroesRepo *mocks.HeroesRepository
var ctx = context.Background()
var inserted *mongo.InsertOneResult

//var input = models.Heroes{}

func TestMain(m *testing.M) {
	heroesRepo = &mocks.HeroesRepository{}
	repository.SetHeroesRepository(heroesRepo)
	code := m.Run()
	os.Exit(code)
}

func TestCreateHeroeF(t *testing.T) {
	h := models.Heroes{
		Name:        "Mei",
		Role:        "Damage",
		Description: "Testing heroe",
	}
	heroesRepo.On("CreateHero", ctx, &h).Return(inserted, nil)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(h)
	if err != nil {
		t.Errorf("Error to parse heroe: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/heroes", &buf)
	res := httptest.NewRecorder()
	defer req.Body.Close()

	handler.CreateHeroeF(res, req)
	log.Println(res.Code)
	log.Println(res.Body)
}
