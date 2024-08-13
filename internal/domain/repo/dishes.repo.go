package repo

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IDishesRepository interface {
	GetById(id string) (entity.Dish, error)
	GetAll(kitchenId string) ([]entity.Dish, error)
	Create(d entity.Dish) (entity.Dish, error)
	Delete(id string) (bool, error)
	GetRandom(kitchenId string) (entity.Dish, error)
	GetInterval(limit int, start int, sort int, kitchenId string) ([]entity.Dish, error)
}
