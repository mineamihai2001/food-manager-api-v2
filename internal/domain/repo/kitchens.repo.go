package repo

import "github.com/mineamihai2001/fm/internal/domain/model"

type IKitchensRepository interface {
	GetById(id string) (model.Kitchen, error)
	GetAll() ([]model.Kitchen, error)
	Create(i model.Kitchen) (model.Kitchen, error)
	Delete(id string) (bool, error)
}
