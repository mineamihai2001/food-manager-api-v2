package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	api_error "github.com/mineamihai2001/fm/internal/api/api-error"
	"github.com/mineamihai2001/fm/internal/infrastructure/repository"
	"github.com/mineamihai2001/fm/internal/infrastructure/services/kitchens"
)

func Kitchen(ctx *gin.Context) {
	// validate if kitchen exists
	service := kitchens.NewKitchensService(repository.NewKitchensRepository())
	kitchen, err := service.GetById(ctx.Param("kitchenId"))

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			api_error.New(http.StatusNotFound, errors.New("kitchen not found")),
		)
		ctx.Abort()
		return
	}

	ctx.Set("kitchenId", kitchen.Id)
	ctx.Next()
}
