package dtos

type GetIngredientsPage struct {
	Page     string `form:"page" binding:"required"`
	Sort     string `form:"sort" binding:"required"`
	PageSize string `form:"pageSize" binding:"required"`
}
