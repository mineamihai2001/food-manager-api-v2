package dishes

import "github.com/mineamihai2001/fm/cmd/domain"

type DishDetails struct {
	Id          string              `bson:"_id,omitempty" json:"id"`
	Kitchen     domain.Kitchen      `bson:"kitchen" json:"kitchen"`
	Name        string              `bson:"name" json:"name"`
	Ingredients []domain.Ingredient `bson:"ingredients" json:"ingredients"`
	Duration    int                 `bson:"duration" json:"duration"`
	Rating      int                 `bson:"rating" json:"rating"`
	Images      []string            `bson:"images" json:"images"`
}

func NewDishDetails(kitchen domain.Kitchen, name string, ingredients []domain.Ingredient, duration int, rating int, images []string) *DishDetails {
	return &DishDetails{
		Kitchen:     kitchen,
		Name:        name,
		Ingredients: ingredients,
		Duration:    duration,
		Rating:      rating,
		Images:      images,
	}
}
