package dtos

type GetDishesPageDto struct {
	Page     string `form:"page" binding:"required"`
	Sort     string `form:"sort" binding:"required"`
	PageSize string `form:"pageSize" binding:"required"`
}
