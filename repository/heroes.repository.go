package repository

import (
	"context"

	"github.com/ferjoaguilar/testing-example-golang/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type HeroRepository interface {
	CreateHero(ctx context.Context, hero *models.Heroes) (*mongo.InsertOneResult, error)
}

var HeroImplementation HeroRepository

func SetHeroRepository(repository HeroRepository) {
	HeroImplementation = repository
}

func CreateHero(ctx context.Context, hero *models.Heroes) (*mongo.InsertOneResult, error) {
	return HeroImplementation.CreateHero(ctx, hero)
}
