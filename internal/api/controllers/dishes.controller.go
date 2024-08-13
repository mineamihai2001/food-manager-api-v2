package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	api_error "github.com/mineamihai2001/fm/internal/api/api_error"
	"github.com/mineamihai2001/fm/internal/api/dtos"
	"github.com/mineamihai2001/fm/internal/api/middleware"
	"github.com/mineamihai2001/fm/internal/domain/model"
	domain "github.com/mineamihai2001/fm/internal/domain/services"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
)

type DishesController struct {
	dishesService domain.IDishesService
}

func NewDishesController(dishesService domain.IDishesService) *DishesController {
	c := &DishesController{
		dishesService,
	}

	return c
}

func (c *DishesController) Create(ctx *gin.Context) {
	body, err := middleware.Body[dtos.CreateDishDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	kitchenId := ctx.GetString("kitchenId")

	var ingredientParts = make([]model.IngredientPart, 0)
	for _, i := range body.Ingredients {
		ingredientParts = append(ingredientParts, model.NewIngredientPart(i.Id, i.Unit, *i.Size))
	}

	res, err := c.dishesService.Create(
		kitchenId,
		body.Name,
		*body.Duration,
		*body.Rating,
		body.Images,
		body.Steps,
		ingredientParts,
	)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishesController) GetAll(ctx *gin.Context) {
	res, err := c.dishesService.GetAll(ctx.GetString("kitchenId"))

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishesController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.dishesService.GetById(id)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishesController) GetRandom(ctx *gin.Context) {
	res, err := c.dishesService.GetRandom(ctx.GetString("kitchenId"))

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishesController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.dishesService.Delete(id)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	var status int
	if res {
		status = http.StatusOK
	} else {
		status = http.StatusNotModified
	}

	ctx.JSON(status, gin.H{
		"deleted": res,
	})
}

func (c *DishesController) GetPage(ctx *gin.Context) {
	query, err := middleware.Query[dtos.GetDishesPageDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	page, pageErr := strconv.Atoi(query.Page)
	sort, sortErr := strconv.Atoi(query.Sort)
	pageSize, pageSizeErr := strconv.Atoi(query.PageSize)

	if pageErr != nil || sortErr != nil || pageSizeErr != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("type error, query parameters invalid type. Page, sort and pageSize must be of type int")),
		)
		return
	}

	res, err := c.dishesService.GetPage(page, pageSize, sort, ctx.GetString("kitchenId"))

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewGetDishesPageResponseDto(page, sort, pageSize, res))
}
