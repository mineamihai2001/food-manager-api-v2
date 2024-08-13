package repo

import (
	"github.com/mineamihai2001/fm/internal/domain/entity"
)

type IDishIngredientRepository interface {
	GetById(id string) (entity.DishIngredient, error)
	GetAll(kitchenId string) ([]entity.DishIngredient, error)
	Create(d entity.DishIngredient) (entity.DishIngredient, error)
	Delete(id string) (bool, error)
	GetByDishId(dishId string) ([]entity.DishIngredient, error)
	GetByIngredientId(ingredientId string) ([]entity.DishIngredient, error)
	CountDishIngredient(dishId string, ingredientId string) (int64, error)
}
