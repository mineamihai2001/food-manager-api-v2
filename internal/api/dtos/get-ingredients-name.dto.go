package dtos

type GetIngredientsNameDto struct {
	Name string `form:"name" binding:"required"`
}
