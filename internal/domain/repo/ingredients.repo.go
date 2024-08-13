package repo

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IIngredientsRepository interface {
	GetById(id string) (entity.Ingredient, error)
	GetAll() ([]entity.Ingredient, error)
	GetManyById(ids []string) ([]entity.Ingredient, error)
	Create(i entity.Ingredient) (entity.Ingredient, error)
	CreateMany(i []entity.Ingredient) ([]entity.Ingredient, error)
	Delete(id string) (bool, error)
	DeleteMany(ids []string) (int, error)
	GetInterval(limit int, start int, sort int) ([]entity.Ingredient, error)
	GetByName(name string) ([]entity.Ingredient, error)
}
