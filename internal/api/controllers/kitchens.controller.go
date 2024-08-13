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

type KitchensController struct {
	kitchensService domain.IKitchensService
}

func NewKitchensController(kitchensService domain.IKitchensService) *KitchensController {
	c := &KitchensController{
		kitchensService,
	}

	return c
}

func (c *KitchensController) Create(ctx *gin.Context) {
	body, err := middleware.Body[dtos.CreateKitchenDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, err),
		)
		return
	}

	res, err := c.kitchensService.Create(body.Name)

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(200, res)
}

func (c *KitchensController) GetAll(ctx *gin.Context) {
	res, err := c.kitchensService.GetAll()

	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *KitchensController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.kitchensService.GetById(id)
	if err != nil {
		ctx.JSON(
			err.(*services.ServiceError).HttpStatus(),
			api_error.New(err.(*services.ServiceError).HttpStatus(), err),
		)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *KitchensController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(
			http.StatusBadRequest,
			api_error.New(http.StatusBadRequest, errors.New("missing id param")),
		)
		return
	}

	res, err := c.kitchensService.Delete(id)
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
