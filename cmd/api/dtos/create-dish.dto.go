package dtos

type CreateDishDto struct {
	KitchenId   string   `json:"kitchenId" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
	Duration    *int     `json:"duration" binding:"required"`
	Rating      *int     `json:"rating" binding:"required"`
	Images      []string `json:"images" binding:"required"`
}
