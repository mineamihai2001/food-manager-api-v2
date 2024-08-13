package model

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IngredientPart struct {
	Id   string      `json:"id"`
	Unit entity.Unit `json:"unit"`
	Size int         `json:"size"`
}

func NewIngredientPart(id string, unit entity.Unit, size int) IngredientPart {
	return IngredientPart{
		Id:   id,
		Unit: unit,
		Size: size,
	}
}
