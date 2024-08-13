package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	api_error "github.com/mineamihai2001/fm/internal/api/api_error"
	"github.com/mineamihai2001/fm/internal/api/dtos"
	"github.com/mineamihai2001/fm/internal/api/middleware"
	domain "github.com/mineamihai2001/fm/internal/domain/services"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
)

type DishIngredientController struct {
	dishIngredientService domain.IDishIngredientService
}

func NewDishIngredientController(dishIngredientService domain.IDishIngredientService) *DishIngredientController {
	return &DishIngredientController{
		dishIngredientService,
	}
}

func (c *DishIngredientController) Create(ctx *gin.Context) {
	body, err := middleware.Body[dtos.CreateDishIngredientDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	res, err := c.dishIngredientService.Create(body.DishId, body.IngredientId, body.Unit, body.Size)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishIngredientController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.dishIngredientService.GetById(id)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishIngredientController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.dishIngredientService.Delete(id)
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

func (c *DishIngredientController) GetByDishId(ctx *gin.Context) {
	dishId := ctx.Param("dishId")
	if dishId == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing dishId param")),
		)
		return
	}

	res, err := c.dishIngredientService.GetByDishId(dishId)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DishIngredientController) GetByIngredientId(ctx *gin.Context) {
	dishId := ctx.Param("ingredientId")
	if dishId == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing ingredientId param")),
		)
		return
	}

	res, err := c.dishIngredientService.GetByIngredientId(dishId)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
