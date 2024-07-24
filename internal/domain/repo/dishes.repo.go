package repo

import "github.com/mineamihai2001/fm/internal/domain/model"

type IDishesRepository interface {
	GetById(id string) (model.Dish, error)
	GetAll(kitchenId string) ([]model.Dish, error)
	Create(d model.Dish) (model.Dish, error)
	Delete(id string) (bool, error)
	GetRandom(kitchenId string) (model.Dish, error)
	GetInterval(limit int, start int, sort int, kitchenId string) ([]model.Dish, error)
	GetByIngredientIds(ingredientIds []string) ([]model.Dish, error)
}
