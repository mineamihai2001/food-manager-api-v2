package services

import "github.com/mineamihai2001/fm/internal/domain/model"

type IIngredientsServices interface {
	Create(name string) (*model.Ingredient, error)
	CreateMany(names []string) ([]model.Ingredient, error)
	GetById(id string) (*model.Ingredient, error)
	GetManyById(ids []string) ([]model.Ingredient, error)
	GetAll() ([]model.Ingredient, error)
	Delete(id string) (bool, error)
	DeleteMany(ids []string) (int, error)
	GetPage(page int, pageSize int, sort int) ([]model.Ingredient, error)
	GetByName(name string) ([]model.Ingredient, error)
}
