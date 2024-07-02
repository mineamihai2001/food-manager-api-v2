package dtos

type DeleteIngredientsDto struct {
	Ids []string `json:"ids" binding:"required"`
}
