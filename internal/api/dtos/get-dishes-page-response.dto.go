package dtos

import (
	"github.com/mineamihai2001/fm/internal/domain/model"
)

type GetDishesPageResponseDto struct {
	Page     int                 `json:"page"`
	Sort     int                 `json:"sort"`
	PageSize int                 `json:"pageSize"`
	Items    []model.DishDetails `json:"items"`
}

func NewGetDishesPageResponseDto(
	page int,
	sort int,
	pageSize int,
	items []model.DishDetails,
) GetDishesPageResponseDto {
	return GetDishesPageResponseDto{
		Page:     page,
		Sort:     sort,
		PageSize: pageSize,
		Items:    items,
	}
}
