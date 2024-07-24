package services

import "github.com/mineamihai2001/fm/internal/domain/model"

type IKitchensService interface {
	Create(name string) (*model.Kitchen, error)
	GetById(id string) (*model.Kitchen, error)
	GetAll() (*[]model.Kitchen, error)
	Delete(id string) (bool, error)
}
