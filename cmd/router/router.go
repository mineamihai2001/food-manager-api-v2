package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mineamihai2001/fm/cmd/api/controllers"
	"github.com/mineamihai2001/fm/cmd/api/middleware"
)

func Create() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	gen := v1.Group("/")
	{
		pingController := controllers.NewPingController()
		gen.GET("/ping", pingController.Ping)
	}

	i := v1.Group("/ingredients")
	{
		i.Use(middleware.Kitchen)

		ingredientsController := controllers.NewIngredientsController()

		i.GET("/:kitchenId", ingredientsController.GetAll)
		i.POST("/:kitchenId", ingredientsController.Create)
		i.DELETE("/:kitchenId", ingredientsController.DeleteMany)
		i.GET("/:kitchenId/:id", ingredientsController.GetById)
		i.DELETE("/:kitchenId/:id", ingredientsController.Delete)
		i.GET("/:kitchenId/query", ingredientsController.GetPage)
		i.POST("/:kitchenId/query", ingredientsController.GetByName)
		i.POST("/:kitchenId/batch", ingredientsController.CreateMany)
	}

	k := v1.Group("/kitchens")
	{
		kitchensController := controllers.NewKitchensController()
		
		k.POST("/", kitchensController.Create)
		k.GET("/", kitchensController.GetAll)
		k.GET("/:id", kitchensController.GetById)
		k.DELETE("/:id", kitchensController.Delete)
	}

	d := v1.Group("/dishes")
	{
		d.Use(middleware.Kitchen)
		// d.Use(middleware.DishDetails(:kitchenId))

		dishesController := controllers.NewDishesController()

		d.POST("/:kitchenId", dishesController.Create)
		d.GET("/:kitchenId", dishesController.GetAll)
		d.GET("/:kitchenId/random", dishesController.GetRandom)
		d.GET("/:kitchenId/:id", dishesController.GetById)
		d.DELETE("/:kitchenId/:id", dishesController.Delete)
		d.GET("/:kitchenId/query", dishesController.GetPage)
		d.GET("/:kitchenId/details/:id", dishesController.GetDetailsById)
		d.POST("/:kitchenId/ingredients", dishesController.GetByIngredientIds)
		d.POST("/:kitchenId/ingredients/details", dishesController.GetDetailsByIngredientIds)
	}

	return router
}
