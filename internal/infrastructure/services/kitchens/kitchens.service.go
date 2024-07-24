package kitchens

import (
	"github.com/mineamihai2001/fm/internal/domain/model"
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

func (s *KitchensService) Create(name string) (*model.Kitchen, error) {
	res, err := s.repository.Create(model.NewKitchen(name))

	if err != nil {
		return &model.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) GetById(id string) (*model.Kitchen, error) {
	res, err := s.repository.GetById(id)

	if err != nil {
		return &model.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) GetAll() (*[]model.Kitchen, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return &[]model.Kitchen{},
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