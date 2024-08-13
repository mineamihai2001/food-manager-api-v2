package dtos

import "github.com/mineamihai2001/fm/internal/domain/entity"

type CreateDishIngredientDto struct {
	DishId       string      `json:"dishId" binding:"required"`
	IngredientId string      `json:"ingredientId" binding:"required"`
	Unit         entity.Unit `json:"unit" binding:"required"`
	Size         int         `json:"size" binding:"required"`
}
