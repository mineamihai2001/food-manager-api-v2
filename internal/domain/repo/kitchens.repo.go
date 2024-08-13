package repo

import "github.com/mineamihai2001/fm/internal/domain/entity"

type IKitchensRepository interface {
	GetById(id string) (entity.Kitchen, error)
	GetAll() ([]entity.Kitchen, error)
	Create(i entity.Kitchen) (entity.Kitchen, error)
	Delete(id string) (bool, error)
}
