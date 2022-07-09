package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/ferjoaguilar/testing-example-golang/models"
	"github.com/ferjoaguilar/testing-example-golang/repository"
	"github.com/ferjoaguilar/testing-example-golang/repository/mocks"
	"go.mongodb.org/mongo-driver/mongo"
)

var heroesRepo *mocks.HeroesRepository

func TestMain(m *testing.M) {
	heroesRepo = &mocks.HeroesRepository{}

	repository.SetHeroesRepository(heroesRepo)
	code := m.Run()
	os.Exit(code)
}

func TestCreateHero(t *testing.T) {
	// Arrage
	input := models.Heroes{}
	ctx := context.Background()
	var inserted *mongo.InsertOneResult

	// Act
	t.Parallel()
	heroesRepo.On("CreateHero", ctx, &input).Return(inserted, nil)
	_, err := repository.CreateHero(ctx, &input)

	// Assert
	if err != nil {
		t.Errorf("Created hero inscorrect got %v want %v", err, nil)
	}
}
