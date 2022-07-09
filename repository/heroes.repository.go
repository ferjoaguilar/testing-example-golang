package repository

import (
	"context"

	"github.com/ferjoaguilar/testing-example-golang/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type HeroesRepository interface {
	CreateHero(ctx context.Context, hero *models.Heroes) (*mongo.InsertOneResult, error)
}

var HeroesImplementation HeroesRepository

func SetHeroesRepository(repository HeroesRepository) {
	HeroesImplementation = repository
}

func CreateHero(ctx context.Context, hero *models.Heroes) (*mongo.InsertOneResult, error) {
	return HeroesImplementation.CreateHero(ctx, hero)
}
