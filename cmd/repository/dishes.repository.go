package repository

import (
	"context"

	"github.com/mineamihai2001/fm/cmd/domain"
	"github.com/mineamihai2001/fm/internal/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DishesRepository struct {
	dataSource *mongo.Collection[domain.Dish]
}

func NewDishesRepository() *DishesRepository {
	client := mongo.GetInstance("fm", context.Background())

	return &DishesRepository{
		dataSource: mongo.GetCollection[domain.Dish](client, "dishes"),
	}
}

func (r *DishesRepository) GetById(id string) (domain.Dish, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Dish{}, err
	}
	return r.dataSource.FindOne(bson.D{{Key: "_id", Value: objectId}})
}

func (r *DishesRepository) GetAll() ([]domain.Dish, error) {
	return r.dataSource.Find(bson.D{})
}

func (r *DishesRepository) Create(i domain.Dish) (domain.Dish, error) {
	res, err := r.dataSource.InsertOne(i)
	if err != nil {
		return domain.Dish{}, err
	}

	created := i
	created.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return created, nil
}

func (r *DishesRepository) Delete(id string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	res, err := r.dataSource.DeleteOne(bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		return false, err
	}

	return res.DeletedCount > 0, nil
}
