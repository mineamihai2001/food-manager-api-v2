package services

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IKitchensService interface {
	Create(name string) (*entity.Kitchen, error)
	GetById(id string) (*entity.Kitchen, error)
	GetAll() (*[]entity.Kitchen, error)
	Delete(id string) (bool, error)
}
