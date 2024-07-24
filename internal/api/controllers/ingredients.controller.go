package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	api_error "github.com/mineamihai2001/fm/internal/api/api-error"
	"github.com/mineamihai2001/fm/internal/api/dtos"
	"github.com/mineamihai2001/fm/internal/api/middleware"
	domain "github.com/mineamihai2001/fm/internal/domain/services"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
)

type IngredientsController struct {
	ingredientsService domain.IIngredientsServices
}

func NewIngredientsController(ingredientsService domain.IIngredientsServices) *IngredientsController {
	c := &IngredientsController{
		ingredientsService,
	}

	return c
}

func (c *IngredientsController) Create(ctx *gin.Context) {
	body, err := middleware.Body[dtos.CreateIngredientDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	res, err := c.ingredientsService.Create(body.Name)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IngredientsController) CreateMany(ctx *gin.Context) {
	body, err := middleware.Body[dtos.CreateManyIngredientsDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	res, err := c.ingredientsService.CreateMany(body.Names)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IngredientsController) GetAll(ctx *gin.Context) {
	res, err := c.ingredientsService.GetAll()

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IngredientsController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.ingredientsService.GetById(id)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IngredientsController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.ingredientsService.Delete(id)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"deleted": res,
	})
}

func (c *IngredientsController) DeleteMany(ctx *gin.Context) {
	body, err := middleware.Body[dtos.DeleteIngredientsDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	res, err := c.ingredientsService.DeleteMany(body.Ids)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"deleted": res,
	})
}

func (c *IngredientsController) GetPage(ctx *gin.Context) {
	query, err := middleware.Query[dtos.GetIngredientsPageDto](ctx)
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

	res, err := c.ingredientsService.GetPage(page, pageSize, sort)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *IngredientsController) GetByName(ctx *gin.Context) {
	query, err := middleware.Body[dtos.GetIngredientsNameDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	res, err := c.ingredientsService.GetByName(query.Name)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
