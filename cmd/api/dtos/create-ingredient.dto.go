package dtos

type CreateIngredientDto struct {
	Name string `json:"name" binding:"required"`
}
