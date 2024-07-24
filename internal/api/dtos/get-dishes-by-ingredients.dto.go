package dtos

type GetDishesByIngredientsDto struct {
	IngredientIds []string `form:"ingredientIds" binding:"required"`
}
