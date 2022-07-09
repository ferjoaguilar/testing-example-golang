// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/ferjoaguilar/testing-example-golang/models"
import mongo "go.mongodb.org/mongo-driver/mongo"

// HeroesRepository is an autogenerated mock type for the HeroesRepository type
type HeroesRepository struct {
	mock.Mock
}

// CreateHero provides a mock function with given fields: ctx, hero
func (_m *HeroesRepository) CreateHero(ctx context.Context, hero *models.Heroes) (*mongo.InsertOneResult, error) {
	ret := _m.Called(ctx, hero)

	var r0 *mongo.InsertOneResult
	if rf, ok := ret.Get(0).(func(context.Context, *models.Heroes) *mongo.InsertOneResult); ok {
		r0 = rf(ctx, hero)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertOneResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Heroes) error); ok {
		r1 = rf(ctx, hero)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}