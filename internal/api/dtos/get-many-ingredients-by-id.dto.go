package dtos

type GetManyIngredientsByIdDto struct {
	Ids []string `json:"ids" binding:"required"`
}
