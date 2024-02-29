package dtos

type DeleteIngredients struct {
	Ids []string `json:"ids" binding:"required"`
}
