package repo

import "github.com/mineamihai2001/fm/internal/domain/model"

type IIngredientsRepository interface {
	GetById(id string) (model.Ingredient, error)
	GetAll() ([]model.Ingredient, error)
	GetManyById(ids []string) ([]model.Ingredient, error)
	Create(i model.Ingredient) (model.Ingredient, error)
	CreateMany(i []model.Ingredient) ([]model.Ingredient, error)
	Delete(id string) (bool, error)
	DeleteMany(ids []string) (int, error)
	GetInterval(limit int, start int, sort int) ([]model.Ingredient, error)
	GetByName(name string) ([]model.Ingredient, error)
}
