package repository

import (
	"context"

	"github.com/mineamihai2001/fm/internal/domain"
	"github.com/mineamihai2001/fm/internal/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo_driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *DishesRepository) GetAll(kitchenId string) ([]domain.Dish, error) {
	return r.dataSource.Find(bson.D{{Key: "kitchenId", Value: kitchenId}})
}

func (r *DishesRepository) Create(d domain.Dish) (domain.Dish, error) {
	res, err := r.dataSource.InsertOne(d)
	if err != nil {
		return domain.Dish{}, err
	}

	created := d
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

func (r *DishesRepository) GetRandom(kitchenId string) (domain.Dish, error) {
	matchStage := bson.D{{
		Key: "$match", Value: bson.D{
			{
				Key:   "kitchenId",
				Value: kitchenId,
			},
		},
	}}

	sampleStage := bson.D{
		{
			Key: "$sample", Value: bson.D{
				{
					Key:   "size",
					Value: 1,
				},
			},
		},
	}

	res, err := r.dataSource.Aggregate(mongo_driver.Pipeline{matchStage, sampleStage})

	if len(res) == 0 {
		return domain.Dish{Name: "not found"}, nil
	}

	return res[0], err
}

func (r *DishesRepository) GetInterval(limit int, start int, sort int, kitchenId string) ([]domain.Dish, error) {
	var opts *options.FindOptions
	if sort == 1 || sort == -1 {
		opts = options.Find().SetLimit(int64(limit)).SetSkip(int64(start)).SetSort(bson.D{{Key: "name", Value: sort}})
	} else {
		opts = options.Find().SetLimit(int64(limit)).SetSkip(int64(start))
	}

	return r.dataSource.Find(bson.D{{Key: "kitchenId", Value: kitchenId}}, opts)
}

func (r *DishesRepository) GetByIngredientIds(ingredientIds []string) ([]domain.Dish, error) {
	return r.dataSource.Find(bson.M{"ingredientIds": bson.M{"$in": ingredientIds}})
}
