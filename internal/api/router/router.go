package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mineamihai2001/fm/internal/api/controllers"
	"github.com/mineamihai2001/fm/internal/api/middleware"
	"github.com/mineamihai2001/fm/pkg/tracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Create() *gin.Engine {
	app := gin.Default()
	// auto instrument requests
	app.Use(otelgin.Middleware("fm-v2-api", otelgin.WithFilter(tracing.IgnorePaths)))
	app.Use(middleware.RequestIdMiddleware)

	v1 := app.Group("/v1")

	gen := v1.Group("/")
	{
		pingController := controllers.NewPingController()
		gen.GET("/ping", pingController.Ping)
	}

	ingredientsRouter := v1.Group("/ingredients")
	{
		ingredientsRouter.Use(middleware.Kitchen)

		ingredientsController := controllers.NewIngredientsController()

		ingredientsRouter.GET("/:kitchenId", ingredientsController.GetAll)
		ingredientsRouter.POST("/:kitchenId", ingredientsController.Create)
		ingredientsRouter.DELETE("/:kitchenId", ingredientsController.DeleteMany)
		ingredientsRouter.GET("/:kitchenId/:id", ingredientsController.GetById)
		ingredientsRouter.DELETE("/:kitchenId/:id", ingredientsController.Delete)
		ingredientsRouter.GET("/:kitchenId/query", ingredientsController.GetPage)
		ingredientsRouter.POST("/:kitchenId/query", ingredientsController.GetByName)
		ingredientsRouter.POST("/:kitchenId/batch", ingredientsController.CreateMany)
	}

	kitchensRouter := v1.Group("/kitchens")
	{
		kitchensController := controllers.NewKitchensController()

		kitchensRouter.POST("/", kitchensController.Create)
		kitchensRouter.GET("/", kitchensController.GetAll)
		kitchensRouter.GET("/:id", kitchensController.GetById)
		kitchensRouter.DELETE("/:id", kitchensController.Delete)
	}

	dishesRouter := v1.Group("/dishes")
	{
		dishesRouter.Use(middleware.Kitchen)
		// d.Use(middleware.DishDetails(:kitchenId))

		dishesController := controllers.NewDishesController()

		dishesRouter.POST("/:kitchenId", dishesController.Create)
		dishesRouter.GET("/:kitchenId", dishesController.GetAll)
		dishesRouter.GET("/:kitchenId/random", dishesController.GetRandom)
		dishesRouter.GET("/:kitchenId/:id", dishesController.GetById)
		dishesRouter.DELETE("/:kitchenId/:id", dishesController.Delete)
		dishesRouter.GET("/:kitchenId/query", dishesController.GetPage)
		dishesRouter.GET("/:kitchenId/details/:id", dishesController.GetDetailsById)
		dishesRouter.POST("/:kitchenId/ingredients", dishesController.GetByIngredientIds)
		dishesRouter.POST("/:kitchenId/ingredients/details", dishesController.GetDetailsByIngredientIds)
	}

	return app
}
