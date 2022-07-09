package database

import (
	"context"

	"github.com/ferjoaguilar/testing-example-golang/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongodbRepository) CreateHero(ctx context.Context, hero *models.Heroes) (*mongo.InsertOneResult, error) {
	inserted, err := repo.DB.Collection("heroes").InsertOne(ctx, hero)
	if err != nil {
		return nil, err
	}

	return inserted, nil
}
