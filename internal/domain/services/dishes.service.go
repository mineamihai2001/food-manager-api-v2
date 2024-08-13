package services

import (
	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/domain/model"
)

type IDishesService interface {
	Create(
		kitchenId string,
		name string,
		duration int,
		rating int,
		images []string,
		steps []string,
		ingredientParts []model.IngredientPart,
	) (*model.DishDetails, error)
	GetById(id string) (*model.DishDetails, error)
	GetAll(kitchenId string) (*[]entity.Dish, error)
	Delete(id string) (bool, error)
	GetRandom(kitchenId string) (*entity.Dish, error)
	GetPage(page int, pageSize int, sort int, kitchenId string) ([]model.DishDetails, error)
}
