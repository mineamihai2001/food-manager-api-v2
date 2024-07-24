package repository

import (
	"context"

	"github.com/mineamihai2001/fm/internal/domain/model"
	"github.com/mineamihai2001/fm/internal/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KitchensRepository struct {
	dataSource *mongo.Collection[model.Kitchen]
}

func NewKitchensRepository() *KitchensRepository {
	client := mongo.GetInstance("fm", context.Background())

	return &KitchensRepository{
		dataSource: mongo.GetCollection[model.Kitchen](client, "kitchens"),
	}
}

func (r *KitchensRepository) GetById(id string) (model.Kitchen, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Kitchen{}, err
	}
	return r.dataSource.FindOne(bson.D{{Key: "_id", Value: objectId}})
}

func (r *KitchensRepository) GetAll() ([]model.Kitchen, error) {
	return r.dataSource.Find(bson.D{})
}

func (r *KitchensRepository) Create(i model.Kitchen) (model.Kitchen, error) {
	res, err := r.dataSource.InsertOne(i)
	if err != nil {
		return model.Kitchen{}, err
	}

	created := i
	created.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return created, nil
}

func (r *KitchensRepository) Delete(id string) (bool, error) {
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
