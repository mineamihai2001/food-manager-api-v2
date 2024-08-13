package repository

import (
	"context"

	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/helpers"
	"github.com/mineamihai2001/fm/internal/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KitchensRepository struct {
	dataSource *mongo.Collection[entity.Kitchen]
}

func NewKitchensRepository() *KitchensRepository {
	env := helpers.Env()
	client := mongo.GetInstance(env.Db.Database, context.Background())

	return &KitchensRepository{
		dataSource: mongo.GetCollection[entity.Kitchen](client, "kitchens"),
	}
}

func (r *KitchensRepository) GetById(id string) (entity.Kitchen, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Kitchen{}, err
	}
	return r.dataSource.FindOne(bson.D{{Key: "_id", Value: objectId}})
}

func (r *KitchensRepository) GetAll() ([]entity.Kitchen, error) {
	return r.dataSource.Find(bson.D{})
}

func (r *KitchensRepository) Create(i entity.Kitchen) (entity.Kitchen, error) {
	res, err := r.dataSource.InsertOne(i)
	if err != nil {
		return entity.Kitchen{}, err
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
