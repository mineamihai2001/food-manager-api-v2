package services

import "github.com/mineamihai2001/fm/internal/domain/model"

type IDishesService interface {
	Create(
		kitchenId string,
		name string,
		ingredientIds []string,
		duration int,
		rating int,
		images []string,
	) (*model.Dish, error)
	GetById(id string) (*model.Dish, error)
	GetAll(kitchenId string) (*[]model.Dish, error)
	Delete(id string) (bool, error)
	GetRandom(kitchenId string) (*model.Dish, error)
	GetPage(page int, pageSize int, sort int, kitchenId string) ([]model.Dish, error)
	GetByIngredientIds(ids []string) ([]model.Dish, error)
}
