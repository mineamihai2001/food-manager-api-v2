package services

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IDishIngredientService interface {
	GetById(id string) (entity.DishIngredient, error)
	Delete(id string) (bool, error)
	Create(dishId string, ingredientId string, unit entity.Unit, size int) (entity.DishIngredient, error)
	GetByDishId(dishId string) ([]entity.DishIngredient, error)
	GetByIngredientId(ingredientId string) ([]entity.DishIngredient, error)
}
