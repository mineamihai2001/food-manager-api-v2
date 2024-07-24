package dtos

type CreateManyIngredientsDto struct {
	Names []string `json:"names" binding:"required"`
}
