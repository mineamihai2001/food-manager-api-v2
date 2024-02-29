package kitchens

import (
	"github.com/mineamihai2001/fm/cmd/domain"
	"github.com/mineamihai2001/fm/cmd/repository"
	"github.com/mineamihai2001/fm/cmd/services"
)

type KitchensService struct {
	repository *repository.KitchensRepository
}

func New() *KitchensService {
	return &KitchensService{
		repository: repository.NewKitchensRepository(),
	}
}

func (s *KitchensService) Create(name string) (*domain.Kitchen, error) {
	res, err := s.repository.Create(domain.NewKitchen(name))

	if err != nil {
		return &domain.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) GetById(id string) (*domain.Kitchen, error) {
	res, err := s.repository.GetById(id)

	if err != nil {
		return &domain.Kitchen{},
			services.NewServiceError(services.InternalServerError, err.Error())
	}

	return &res, nil
}

func (s *KitchensService) GetAll() (*[]domain.Kitchen, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return &[]domain.Kitchen{},
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
