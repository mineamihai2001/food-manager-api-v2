package model

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IngredientDetails struct {
	Id   string      `json:"id"`
	Name string      `json:"name"`
	Unit entity.Unit `json:"unit"`
	Size int         `json:"size"`
}

func NewIngredientDetails(id string, name string, unit entity.Unit, size int) IngredientDetails {
	return IngredientDetails{
		Id:   id,
		Name: name,
		Unit: unit,
		Size: size,
	}
}
