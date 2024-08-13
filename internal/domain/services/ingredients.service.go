package services

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IIngredientsServices interface {
	Create(name string) (*entity.Ingredient, error)
	CreateMany(names []string) ([]entity.Ingredient, error)
	GetById(id string) (*entity.Ingredient, error)
	GetManyById(ids []string) ([]entity.Ingredient, error)
	GetAll() ([]entity.Ingredient, error)
	Delete(id string) (bool, error)
	DeleteMany(ids []string) (int, error)
	GetPage(page int, pageSize int, sort int) ([]entity.Ingredient, error)
	GetByName(name string) ([]entity.Ingredient, error)
}
