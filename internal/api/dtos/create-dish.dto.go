package dtos

type CreateDishDto struct {
	Name        string               `json:"name" binding:"required"`
	Duration    *int                 `json:"duration" binding:"required"`
	Rating      *int                 `json:"rating" binding:"required"`
	Images      []string             `json:"images" binding:"required"`
	Steps       []string             `json:"steps" binding:"required"`
	Ingredients []IngredientPartDto `json:"ingredients" binding:"dive"`
}
