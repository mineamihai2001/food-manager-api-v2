package repository

import (
	"context"

	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/helpers"
	"github.com/mineamihai2001/fm/internal/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DishIngredientRepository struct {
	dataSource *mongo.Collection[entity.DishIngredient]
}

func NewDishIngredientRepository() *DishIngredientRepository {
	env := helpers.Env()
	client := mongo.GetInstance(env.Db.Database, context.Background())

	return &DishIngredientRepository{
		dataSource: mongo.GetCollection[entity.DishIngredient](client, "dish-ingredients"),
	}
}

func (r *DishIngredientRepository) GetById(id string) (entity.DishIngredient, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.DishIngredient{}, err
	}

	return r.dataSource.FindOne(bson.D{{Key: "_id", Value: objectId}})
}

func (r *DishIngredientRepository) GetAll(kitchenId string) ([]entity.DishIngredient, error) {
	return r.dataSource.Find(bson.D{{Key: "kitchenId", Value: kitchenId}})
}

func (r *DishIngredientRepository) Create(d entity.DishIngredient) (entity.DishIngredient, error) {
	res, err := r.dataSource.InsertOne(d)
	if err != nil {
		return entity.DishIngredient{}, err
	}

	created := d
	created.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return created, nil
}

func (r *DishIngredientRepository) Delete(id string) (bool, error) {
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

func (r *DishIngredientRepository) GetByDishId(dishId string) ([]entity.DishIngredient, error) {
	return r.dataSource.Find(bson.D{{Key: "dishId", Value: dishId}})
}

func (r *DishIngredientRepository) GetByIngredientId(ingredientId string) ([]entity.DishIngredient, error) {
	return r.dataSource.Find(bson.D{{Key: "ingredientId", Value: ingredientId}})
}

func (r *DishIngredientRepository) CountDishIngredient(dishId string, ingredientId string) (int64, error) {
	return r.dataSource.CountDocuments(bson.D{
		{Key: "ingredientId", Value: ingredientId},
		{Key: "dishId", Value: dishId},
	})
}
