package dtos

import "github.com/mineamihai2001/fm/internal/domain/entity"

type GetIngredientsPageResponseDto struct {
	Page     int                 `json:"page"`
	Sort     int                 `json:"sort"`
	PageSize int                 `json:"pageSize"`
	Items    []entity.Ingredient `json:"items"`
}

func NewGetIngredientsPageResponseDto(
	page int,
	sort int,
	pageSize int,
	items []entity.Ingredient,
) GetIngredientsPageResponseDto {
	return GetIngredientsPageResponseDto{
		Page:     page,
		Sort:     sort,
		PageSize: pageSize,
		Items:    items,
	}
}
