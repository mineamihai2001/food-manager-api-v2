package kitchens

import (
	"github.com/mineamihai2001/fm/internal/domain/entity"
	"github.com/mineamihai2001/fm/internal/domain/repo"
	"github.com/mineamihai2001/fm/internal/infrastructure/services"
)

type KitchensService struct {
	repository repo.IKitchensRepository
}

func NewKitchensService(repository repo.IKitchensRepository) *KitchensService {
	return &KitchensService{
		repository,
	}
}

func (s *KitchensService) Create(name string) (*entity.Kitchen, error) {
	res, err := s.repository.Create(entity.NewKitchen(name))

	if err != nil {
		return &entity.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) GetById(id string) (*entity.Kitchen, error) {
	res, err := s.repository.GetById(id)

	if err != nil {
		return &entity.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) GetAll() (*[]entity.Kitchen, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return &[]entity.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) Delete(id string) (bool, error) {
	res, err := s.repository.Delete(id)

	if err != nil {
		return false,
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return res, nil
}
