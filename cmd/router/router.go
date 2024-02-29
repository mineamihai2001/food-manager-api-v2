package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mineamihai2001/fm/cmd/api/controllers"
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
		ingredientsController := controllers.NewIngredientsController()
		i.GET("/", ingredientsController.GetAll)
		i.POST("/", ingredientsController.Create)
		i.DELETE("/", ingredientsController.DeleteMany)
		i.GET("/:id", ingredientsController.GetById)
		i.DELETE("/:id", ingredientsController.Delete)
		i.GET("/query", ingredientsController.GetPage)
		i.POST("/query", ingredientsController.GetByName)
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
		dishesController := controllers.NewDishesController()
		d.POST("/", dishesController.Create)
		d.GET("/", dishesController.GetAll)
		d.GET("/random", dishesController.GetRandom)
		d.GET("/:id", dishesController.GetById)
		d.DELETE("/:id", dishesController.Delete)
		d.GET("/query", dishesController.GetPage)
	}

	return router
}
