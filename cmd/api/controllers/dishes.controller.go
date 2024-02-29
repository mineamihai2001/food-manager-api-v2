package controllers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	api_error "github.com/mineamihai2001/fm/cmd/api/api-error"
	"github.com/mineamihai2001/fm/cmd/api/dtos"
	"github.com/mineamihai2001/fm/cmd/api/middleware"
	"github.com/mineamihai2001/fm/cmd/services"
	"github.com/mineamihai2001/fm/cmd/services/dishes"
)

type DishesController struct {
	context       context.Context
	cancelContext context.CancelFunc
	dishesService *dishes.DishesService
}

func NewDishesController() *DishesController {
	c := &DishesController{
		dishesService: dishes.New(),
	}

	c.context, c.cancelContext = c.createContext()
	return c
}

func (c *DishesController) createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
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

	res, err := c.dishesService.Create(body.KitchenId, body.Name, body.IngredientIds, *body.Duration, *body.Rating, body.Images)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(200, res)
}

func (c *DishesController) GetAll(ctx *gin.Context) {
	res, err := c.dishesService.GetAll()

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

func (c *DishesController) GetDetailsById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.dishesService.GetDetailsById(id)
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
	res, err := c.dishesService.GetRandom()

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

	ctx.JSON(http.StatusOK, gin.H{
		"deleted": res,
	})
}

func (c *DishesController) GetPage(ctx *gin.Context) {
	query, err := middleware.Query[dtos.GetDishesPage](ctx)
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

	res, err := c.dishesService.GetPage(page, pageSize, sort)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
