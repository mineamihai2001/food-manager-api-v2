package repository

import (
	"context"

	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/helpers"
	"github.com/mineamihai2001/fm/internal/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IngredientsRepository struct {
	dataSource *mongo.Collection[entity.Ingredient]
}

func NewIngredientsRepository() *IngredientsRepository {
	env := helpers.Env()
	client := mongo.GetInstance(env.Db.Database, context.Background())

	return &IngredientsRepository{
		dataSource: mongo.GetCollection[entity.Ingredient](client, "ingredients"),
	}
}

func (r *IngredientsRepository) GetById(id string) (entity.Ingredient, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Ingredient{}, err
	}
	return r.dataSource.FindOne(bson.D{{Key: "_id", Value: objectId}})
}

func (r *IngredientsRepository) GetAll() ([]entity.Ingredient, error) {
	return r.dataSource.Find(bson.D{})
}

func (r *IngredientsRepository) GetManyById(ids []string) ([]entity.Ingredient, error) {
	objectIds := make([]primitive.ObjectID, len(ids))

	for _, id := range ids {
		current, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return []entity.Ingredient{}, err
		}
		objectIds = append(objectIds, current)
	}

	return r.dataSource.Find(bson.M{"_id": bson.M{"$in": objectIds}})
}

func (r *IngredientsRepository) Create(i entity.Ingredient) (entity.Ingredient, error) {
	res, err := r.dataSource.InsertOne(i)
	if err != nil {
		return entity.Ingredient{}, err
	}

	created := i
	created.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return created, nil
}

func (r *IngredientsRepository) CreateMany(i []entity.Ingredient) ([]entity.Ingredient, error) {
	res, err := r.dataSource.InsertMany(i)
	if err != nil {
		return []entity.Ingredient{}, err
	}

	created := i
	for i := range created {
		created[i].Id = res.InsertedIDs[i].(primitive.ObjectID).Hex()
	}
	return created, nil
}

func (r *IngredientsRepository) Delete(id string) (bool, error) {
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

func (r *IngredientsRepository) DeleteMany(ids []string) (int, error) {
	objectIds := make([]primitive.ObjectID, len(ids))

	for _, id := range ids {
		current, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return 0, err
		}
		objectIds = append(objectIds, current)
	}

	res, err := r.dataSource.DeleteMany(bson.M{"_id": bson.M{"$in": objectIds}})
	if err != nil {
		return 0, err
	}

	return int(res.DeletedCount), nil
}

func (r *IngredientsRepository) GetInterval(limit int, start int, sort int) ([]entity.Ingredient, error) {
	var opts *options.FindOptions
	if sort == 1 || sort == -1 {
		opts = options.Find().SetLimit(int64(limit)).SetSkip(int64(start)).SetSort(bson.D{{Key: "name", Value: sort}})
	} else {
		opts = options.Find().SetLimit(int64(limit)).SetSkip(int64(start))
	}

	return r.dataSource.Find(bson.D{}, opts)
}

func (r *IngredientsRepository) GetByName(name string) ([]entity.Ingredient, error) {
	return r.dataSource.Find(bson.D{{
		Key: "name",
		Value: bson.D{{
			Key:   "$regex",
			Value: name,
		},
		},
	},
	},
	)
}
