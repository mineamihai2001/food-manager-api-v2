package dtos

type CreateKitchenDto struct {
	Name string `json:"name" binding:"required"`
}
