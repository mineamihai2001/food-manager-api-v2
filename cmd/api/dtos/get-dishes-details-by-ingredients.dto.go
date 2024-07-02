package dtos

type GetDishesDetailsByIngredientsDto struct {
	IngredientIds []string `form:"ingredientIds" binding:"required"`
}
