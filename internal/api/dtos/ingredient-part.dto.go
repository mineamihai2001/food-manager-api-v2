package dtos

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IngredientPartDto struct {
	Id string      `json:"id" binding:"required"`
	Unit         entity.Unit `json:"unit" binding:"required"`
	Size         *int         `json:"size" binding:"required"`
}
