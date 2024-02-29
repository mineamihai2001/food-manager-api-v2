package dtos

type GetIngredientsName struct {
	Name string `form:"name" binding:"required"`
}
