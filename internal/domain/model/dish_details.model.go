package model

import ()

type DishDetails struct {
	Id          string              `json:"id"`
	KitchenId   string              `json:"kitchenId"`
	Name        string              `json:"name"`
	Duration    int                 `json:"duration"`
	Rating      int                 `json:"rating"`
	Images      []string            `json:"images"`
	Steps       []string            `json:"steps"`
	Ingredients []IngredientDetails `json:"ingredients"`
}

func NewDishDetails(
	id string,
	kitchenId string,
	name string,
	duration int,
	rating int,
	images []string,
	steps []string,
	ingredients []IngredientDetails,
) DishDetails {
	return DishDetails{
		Id:          id,
		KitchenId:   kitchenId,
		Name:        name,
		Duration:    duration,
		Rating:      rating,
		Images:      images,
		Steps:       steps,
		Ingredients: ingredients,
	}
}
